package soltxm

import (
	"context"
	"fmt"
	"sync"
	"time"

	solanaGo "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-solana/pkg/solana"
	solanaClient "github.com/smartcontractkit/chainlink-solana/pkg/solana/client"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/config"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/utils"
)

const (
	MaxQueueLen      = 1000
	MaxRetryTimeMs   = 250 // max tx retry time (exponential retry will taper to retry every 0.25s)
	MaxSigsToConfirm = 256 // max number of signatures in GetSignatureStatus call
)

var (
	_ services.ServiceCtx = (*Txm)(nil)
	_ solana.TxManager    = (*Txm)(nil)
)

// Txm manages transactions for the solana blockchain.
// simple implementation with no persistently stored txs
type Txm struct {
	starter utils.StartStopOnce
	lggr    logger.Logger
	chSend  chan PendingTx
	chStop  chan struct{}
	done    sync.WaitGroup
	cfg     config.Config
	txs     PendingTxs // interface so DB can be plugged in
	ks      keystore.Solana
	client  *utils.LazyLoad[solanaClient.ReaderWriter]

	// compute budget unit price parameters
	fee     uint64
	feeLock sync.RWMutex
}

// NewTxm creates a txm. Uses simulation so should only be used to send txes to trusted contracts i.e. OCR.
func NewTxm(chainID string, tc func() (solanaClient.ReaderWriter, error), cfg config.Config, ks keystore.Solana, lggr logger.Logger) *Txm {
	lggr = lggr.Named("Txm")
	return &Txm{
		starter: utils.StartStopOnce{},
		lggr:    lggr,
		chSend:  make(chan PendingTx, MaxQueueLen), // queue can support 1000 pending txs
		chStop:  make(chan struct{}),
		cfg:     cfg,
		txs:     newPendingTxMemoryWithProm(chainID),
		ks:      ks,
		client:  utils.NewLazyLoad(tc),
	}
}

// Start subscribes to queuing channel and processes them.
func (txm *Txm) Start(context.Context) error {
	return txm.starter.StartOnce("solana_txm", func() error {
		txm.done.Add(4) // waitgroup: tx retry, simulator
		go txm.run()
		return nil
	})
}

func (txm *Txm) run() {
	defer txm.done.Done()
	ctx, cancel := utils.ContextFromChan(txm.chStop)
	defer cancel()

	// start confirmer
	go txm.confirm(ctx)

	for {
		select {
		case msg := <-txm.chSend:
			var id uuid.UUID
			// check if tx has been broadcast before and fetch or create UUID
			if msg.broadcast {
				// this should never happen
				if len(msg.signatures) == 0 {
					txm.lggr.Criticalw("tx was already broadcast but has no signatures - dropping from queue", "tx", msg)
					continue
				}
				// this should never happen
				var exists bool
				if id, _, exists = txm.txs.Get(msg.signatures[0]); !exists {
					txm.lggr.Criticalw("tx has signatures but could not find UUID", "signatures", msg.signatures)
					// TODO: requeue or drop?
					continue
				}
			} else {
				id = txm.txs.New(msg)
			}

			// Set compute unit price for transaction, returns a copy of the base tx
			tx, err := msg.SetComputeUnitPrice(ComputeUnitPrice(txm.GetFee()))
			if err != nil {
				txm.lggr.Errorw("failed to set compute unit price in tx", "error", err)
				// TODO: requeue tx?
				continue
			}

			// marshal + sign transaction
			txMsg, err := tx.Message.MarshalBinary()
			if err != nil {
				txm.lggr.Errorw("failed to marshal transaction message", "error", err)
				continue
			}
			sigBytes, err := msg.key.Sign(txMsg) // sign with stored key
			if err != nil {
				txm.lggr.Errorw("failed to sign transaction", "error", err)
				continue
			}
			var finalSig [64]byte
			copy(finalSig[:], sigBytes)
			tx.Signatures = append(tx.Signatures, finalSig)

			// process tx
			sig, err := txm.send(ctx, id, tx)
			if err != nil {
				txm.lggr.Errorw("failed to send transaction", "error", err)
				txm.client.Reset() // clear client if tx fails immediately (potentially bad RPC)
				// TODO: don't give up on tx
				continue // skip remainining
			}

			txm.lggr.Debugw("transaction sent", "signature", sig.String())
		case <-txm.chStop:
			return
		}
	}
}

// sendWithExpBackoff broadcasts a transaction at an exponential backoff rate to increase chances of inclusion by the next validator by rebroadcasting more tx packets
func (txm *Txm) send(chanCtx context.Context, id uuid.UUID, tx *solanaGo.Transaction) (solanaGo.Signature, error) {
	// fetch client
	client, err := txm.client.Get()
	if err != nil {
		return solanaGo.Signature{}, errors.Wrap(err, "failed to get client in soltxm.sendWithExpBackoff")
	}

	// create timeout context
	ctx, cancel := context.WithTimeout(chanCtx, txm.cfg.TxTimeout())
	defer cancel()

	// send initial tx (do not retry and exit early if fails)
	sig, err := client.SendTx(ctx, tx)
	if err != nil {
		txm.txs.OnError(sig, TxFailReject) // increment RPC reject metric
		return solanaGo.Signature{}, errors.Wrap(err, "tx failed initial transmit")
	}

	// store tx signature + cancel function
	if err := txm.txs.Add(id, sig); err != nil {
		return solanaGo.Signature{}, errors.Wrapf(err, "failed to save tx signature (%s) to inflight txs", sig)
	}

	// return signature for use in simulation
	return sig, nil
}

// goroutine that polls to confirm implementation
// cancels the exponential retry once confirmed
func (txm *Txm) confirm(ctx context.Context) {
	defer txm.done.Done()

	tick := time.After(0)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick:
			// get list of tx signatures to confirm
			sigs := txm.txs.ListSignatures()

			// exit switch if not txs to confirm
			if len(sigs) == 0 {
				break
			}

			// get client
			client, err := txm.client.Get()
			if err != nil {
				txm.lggr.Errorw("failed to get client in soltxm.confirm", "error", err)
				break // exit switch
			}

			// batch sigs no more than MaxSigsToConfirm each
			sigsBatch, err := utils.BatchSplit(sigs, MaxSigsToConfirm)
			if err != nil { // this should never happen
				txm.lggr.Criticalw("failed to batch signatures", "error", err)
				break // exit switch
			}

			// process signatures
			processSigs := func(s []solanaGo.Signature, res []*rpc.SignatureStatusesResult) {
				for i := 0; i < len(res); i++ {
					// if status is nil (sig not found), continue polling
					// sig not found could mean invalid tx or not picked up yet
					if res[i] == nil {
						txm.lggr.Debugw("tx state: not found",
							"signature", s[i],
						)

						// check confirm timeout exceeded -> rebroadcast with new bumped fee
						if _, tx, exists := txm.txs.Get(s[i]); exists && time.Since(tx.timestamp) > txm.cfg.TxConfirmTimeout() {
							select {
							case txm.chSend <- tx:
							default:
								txm.lggr.Errorw("failed to enqeue tx", "queueFull", len(txm.chSend) == MaxQueueLen, "tx", tx)
							}
						}
						continue
					}

					// if signature has an error, end polling
					if res[i].Err != nil {
						txm.lggr.Errorw("tx state: failed",
							"signature", s[i],
							"error", res[i].Err,
							"status", res[i].ConfirmationStatus,
						)
						txm.txs.OnError(s[i], TxFailRevert)
						continue
					}

					// if signature is processed, keep polling, don't retry yet (either will become confirmed or dropped)
					if res[i].ConfirmationStatus == rpc.ConfirmationStatusProcessed {
						txm.lggr.Debugw("tx state: processed",
							"signature", s[i],
						)
						continue
					}

					// if signature is confirmed/finalized, end polling
					if res[i].ConfirmationStatus == rpc.ConfirmationStatusConfirmed || res[i].ConfirmationStatus == rpc.ConfirmationStatusFinalized {
						txm.lggr.Debugw(fmt.Sprintf("tx state: %s", res[i].ConfirmationStatus),
							"signature", s[i],
						)
						txm.txs.OnSuccess(s[i])
						continue
					}
				}
			}

			// waitgroup for processing
			var wg sync.WaitGroup
			wg.Add(len(sigsBatch))

			// loop through batch
			for i := 0; i < len(sigsBatch); i++ {
				// fetch signature statuses
				statuses, err := client.SignatureStatuses(ctx, sigsBatch[i])
				if err != nil {
					txm.lggr.Errorw("failed to get signature statuses in soltxm.confirm", "error", err)
					wg.Done() // don't block if exit early
					break     // exit for loop
				}

				// nonblocking: process batches as soon as they come in
				go func(index int) {
					defer wg.Done()
					processSigs(sigsBatch[index], statuses)
				}(i)
			}
			wg.Wait() // wait for processing to finish
		}
		tick = time.After(utils.WithJitter(txm.cfg.ConfirmPollPeriod()))
	}
}

// Enqueue enqueue a msg destined for the solana chain.
func (txm *Txm) Enqueue(accountID string, tx *solanaGo.Transaction) error {
	// validate nil pointer
	if tx == nil {
		return errors.New("error in soltxm.Enqueue: tx is nil pointer")
	}
	// validate account keys slice
	if len(tx.Message.AccountKeys) == 0 {
		return errors.New("error in soltxm.Enqueue: not enough account keys in tx")
	}

	// get signer key
	// fee payer account is index 0 account
	// https://github.com/gagliardetto/solana-go/blob/main/transaction.go#L252
	key, err := txm.ks.Get(tx.Message.AccountKeys[0].String())
	if err != nil {
		return errors.Wrap(err, "error in soltxm.Enqueue.GetKey")
	}

	msg := PendingTx{
		baseTx: tx,
		key:    key,
	}

	select {
	case txm.chSend <- msg:
	default:
		txm.lggr.Errorw("failed to enqeue tx", "queueFull", len(txm.chSend) == MaxQueueLen, "tx", msg)
		return errors.Errorf("failed to enqueue transaction for %s", accountID)
	}
	return nil
}

func (txm *Txm) InflightTxs() int {
	return len(txm.txs.ListSignatures())
}

func (txm *Txm) SetFee(v uint64) {
	txm.feeLock.Lock()
	defer txm.feeLock.Unlock()
	txm.fee = v
}

func (txm *Txm) GetFee() uint64 {
	txm.feeLock.RLock()
	defer txm.feeLock.RUnlock()
	return txm.fee
}

// Close close service
func (txm *Txm) Close() error {
	return txm.starter.StopOnce("solanatxm", func() error {
		close(txm.chStop)
		txm.done.Wait()
		return nil
	})
}

// Healthy service is healthy
func (txm *Txm) Healthy() error {
	return nil
}

// Ready service is ready
func (txm *Txm) Ready() error {
	return nil
}
