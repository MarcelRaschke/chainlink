package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/logger"
)

// NOTE: Currently only returns latest block, could be extended in future to
// return block by number or hash

// Return types:
//
// *evmtypes.Head
type ETHGetBlockTask struct {
	BaseTask   `mapstructure:",squash"`
	EVMChainID string `json:"evmChainID" mapstructure:"evmChainID"`

	chainSet evm.ChainSet
	config   Config
}

var _ Task = (*ETHGetBlockTask)(nil)

func (t *ETHGetBlockTask) Type() TaskType {
	return TaskTypeETHGetBlock
}

func (t *ETHGetBlockTask) Run(ctx context.Context, lggr logger.Logger, vars Vars, inputs []Result) (result Result, runInfo RunInfo) {
	_, err := CheckInputs(inputs, -1, -1, 0)
	if err != nil {
		return Result{Error: errors.Wrap(err, "task inputs")}, runInfo
	}

	var chainID StringParam
	err = errors.Wrap(ResolveParam(&chainID, From(VarExpr(t.EVMChainID, vars), NonemptyString(t.EVMChainID), "")), "evmChainID")
	if err != nil {
		return Result{Error: err}, runInfo
	}

	chain, err := getChainByString(t.chainSet, string(chainID))
	if err != nil {
		return Result{Error: err}, runInfo
	}

	head, err := chain.Client().HeadByNumber(ctx, nil)
	if err != nil {
		return Result{Error: err}, runInfo
	}

	return Result{Value: head}, runInfo
}
