package mercury

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/logger"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

var _ ocrtypes.ContractTransmitter = &MercuryTransmitter{}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MercuryTransmitter struct {
	lggr       logger.Logger
	httpClient HTTPClient

	fromAccount common.Address
	contract    common.Address

	reportURL string
	username  string
	password  string
}

type MercuryReport struct {
	ReportCtx  ocrtypes.ReportContext
	Report     ocrtypes.Report
	Signatures []ocrtypes.AttributedOnchainSignature

	FromAccount     common.Address
	ContractAddress common.Address
}

func NewTransmitter(lggr logger.Logger, contractAddress, fromAccount common.Address, httpClient HTTPClient, reportURL, username, password string) *MercuryTransmitter {
	return &MercuryTransmitter{lggr.Named("Mercury"), httpClient, fromAccount, contractAddress, reportURL, username, password}
}

// Transmit sends the report to the on-chain smart contract's Transmit method.
func (mt *MercuryTransmitter) Transmit(ctx context.Context, reportCtx ocrtypes.ReportContext, report ocrtypes.Report, signatures []ocrtypes.AttributedOnchainSignature) error {
	fmt.Println("BALLS Transmit")
	mt.lggr.Infow("Transmitting report", "report", report, "reportCtx", reportCtx, "signatures", signatures)
	mr := MercuryReport{reportCtx, report, signatures, mt.fromAccount, mt.contract}

	b, err := json.Marshal(mr)
	if err != nil {
		return errors.Wrap(err, "failed to marshal MercuryReport json")
	}

	req, err := http.NewRequest(http.MethodPost, mt.reportURL, bytes.NewReader(b))
	if err != nil {
		return errors.Wrap(err, "failed to instantiate mercury server http request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(mt.username, mt.password)
	req.WithContext(ctx)

	res, err := mt.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to POST to mercury server")
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		mt.lggr.Errorw("Failed to read response body", "err", err)
	}

	fmt.Println("BALLS Transmit 2")
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		mt.lggr.Infow("Transmit report success", "responseStatus", res.Status, "reponseBody", string(respBody), "reportCtx", reportCtx)
	} else {
		mt.lggr.Errorw("Transmit report failed", "responseStatus", res.Status, "reponseBody", string(respBody), "reportCtx", reportCtx)

	}

	return nil
}

func (mt *MercuryTransmitter) FromAccount() ocrtypes.Account {
	return ocrtypes.Account(mt.fromAccount.Hex())
}

// LatestConfigDigestAndEpoch retrieves the latest config digest and epoch from the OCR2 contract.
// It is plugin independent, in particular avoids use of the plugin specific generated evm wrappers
// by using the evm client Call directly for functions/events that are part of OCR2Abstract.
func (mt *MercuryTransmitter) LatestConfigDigestAndEpoch(ctx context.Context) (cd ocrtypes.ConfigDigest, epoch uint32, err error) {
	// TODO: implement this
	err = errors.New("TODO: implement this")
	return
}
