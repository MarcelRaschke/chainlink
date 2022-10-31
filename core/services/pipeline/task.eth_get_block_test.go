package pipeline_test

import (
	"math/big"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	evmmocks "github.com/smartcontractkit/chainlink/core/chains/evm/mocks"
	evmtypes "github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	configtest "github.com/smartcontractkit/chainlink/core/internal/testutils/configtest/v2"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func Test_ETHGetBlockTask(t *testing.T) {
	cfg := configtest.NewGeneralConfig(t, func(c *chainlink.Config, s *chainlink.Secrets) {})

	lggr := logger.TestLogger(t)
	var vars pipeline.Vars
	var inputs []pipeline.Result

	t.Run("if HeadByNumber fails", func(t *testing.T) {
		ethClient := evmmocks.NewClient(t)
		chain := evmmocks.NewChain(t)
		chain.On("Client").Return(ethClient)

		cc := evmtest.NewMockChainSetWithChain(t, chain)

		task := pipeline.ETHGetBlockTask{}
		task.HelperSetDependencies(cc, cfg)
		err := errors.New("foo")
		ethClient.On("HeadByNumber", mock.AnythingOfType("*context.timerCtx"), (*big.Int)(nil)).Return(nil, err)
		res, ri := task.Run(testutils.Context(t), lggr, vars, inputs)

		assert.Equal(t, pipeline.Result(pipeline.Result{Value: interface{}(nil), Error: err}), res)
		assert.Equal(t, pipeline.RunInfo(pipeline.RunInfo{IsRetryable: false, IsPending: false}), ri)
	})

	t.Run("if HeadByNumber succeeds", func(t *testing.T) {
		ethClient := evmmocks.NewClient(t)
		chain := evmmocks.NewChain(t)
		chain.On("Client").Return(ethClient)

		cc := evmtest.NewMockChainSetWithChain(t, chain)

		task := pipeline.ETHGetBlockTask{}
		task.HelperSetDependencies(cc, cfg)
		h := evmtypes.Head{
			Number: testutils.NewRandomPositiveInt64(),
			Hash:   utils.NewHash(),
		}

		ethClient.On("HeadByNumber", mock.AnythingOfType("*context.timerCtx"), (*big.Int)(nil)).Return(&h, nil)
		res, ri := task.Run(testutils.Context(t), lggr, vars, inputs)

		assert.Nil(t, res.Error)
		hVal, is := res.Value.(*evmtypes.Head)
		require.True(t, is)
		assert.Equal(t, h.Number, hVal.Number)
		assert.Equal(t, h.Hash, hVal.Hash)
		assert.Equal(t, pipeline.RunInfo(pipeline.RunInfo{IsRetryable: false, IsPending: false}), ri)
	})
}
