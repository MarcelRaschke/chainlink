// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	client "github.com/smartcontractkit/chainlink/v2/evm/client"

	config "github.com/smartcontractkit/chainlink/v2/evm/config"

	context "context"

	evmtypes "github.com/smartcontractkit/chainlink/v2/evm/types"

	gas "github.com/smartcontractkit/chainlink/v2/evm/gas"

	headtracker "github.com/smartcontractkit/chainlink-framework/chains/headtracker"

	log "github.com/smartcontractkit/chainlink/v2/core/chains/evm/log"

	logger "github.com/smartcontractkit/chainlink/v2/core/logger"

	logpoller "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"

	mock "github.com/stretchr/testify/mock"

	monitor "github.com/smartcontractkit/chainlink/v2/evm/monitor"

	txmgr "github.com/smartcontractkit/chainlink-framework/chains/txmgr"

	types "github.com/smartcontractkit/chainlink-common/pkg/types"
)

// Chain is an autogenerated mock type for the Chain type
type Chain struct {
	mock.Mock
}

type Chain_Expecter struct {
	mock *mock.Mock
}

func (_m *Chain) EXPECT() *Chain_Expecter {
	return &Chain_Expecter{mock: &_m.Mock}
}

// BalanceMonitor provides a mock function with no fields
func (_m *Chain) BalanceMonitor() monitor.BalanceMonitor {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BalanceMonitor")
	}

	var r0 monitor.BalanceMonitor
	if rf, ok := ret.Get(0).(func() monitor.BalanceMonitor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(monitor.BalanceMonitor)
		}
	}

	return r0
}

// Chain_BalanceMonitor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BalanceMonitor'
type Chain_BalanceMonitor_Call struct {
	*mock.Call
}

// BalanceMonitor is a helper method to define mock.On call
func (_e *Chain_Expecter) BalanceMonitor() *Chain_BalanceMonitor_Call {
	return &Chain_BalanceMonitor_Call{Call: _e.mock.On("BalanceMonitor")}
}

func (_c *Chain_BalanceMonitor_Call) Run(run func()) *Chain_BalanceMonitor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_BalanceMonitor_Call) Return(_a0 monitor.BalanceMonitor) *Chain_BalanceMonitor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_BalanceMonitor_Call) RunAndReturn(run func() monitor.BalanceMonitor) *Chain_BalanceMonitor_Call {
	_c.Call.Return(run)
	return _c
}

// Client provides a mock function with no fields
func (_m *Chain) Client() client.Client {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Client")
	}

	var r0 client.Client
	if rf, ok := ret.Get(0).(func() client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	return r0
}

// Chain_Client_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Client'
type Chain_Client_Call struct {
	*mock.Call
}

// Client is a helper method to define mock.On call
func (_e *Chain_Expecter) Client() *Chain_Client_Call {
	return &Chain_Client_Call{Call: _e.mock.On("Client")}
}

func (_c *Chain_Client_Call) Run(run func()) *Chain_Client_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_Client_Call) Return(_a0 client.Client) *Chain_Client_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_Client_Call) RunAndReturn(run func() client.Client) *Chain_Client_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with no fields
func (_m *Chain) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chain_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Chain_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Chain_Expecter) Close() *Chain_Close_Call {
	return &Chain_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Chain_Close_Call) Run(run func()) *Chain_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_Close_Call) Return(_a0 error) *Chain_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_Close_Call) RunAndReturn(run func() error) *Chain_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Config provides a mock function with no fields
func (_m *Chain) Config() config.ChainScopedConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Config")
	}

	var r0 config.ChainScopedConfig
	if rf, ok := ret.Get(0).(func() config.ChainScopedConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.ChainScopedConfig)
		}
	}

	return r0
}

// Chain_Config_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Config'
type Chain_Config_Call struct {
	*mock.Call
}

// Config is a helper method to define mock.On call
func (_e *Chain_Expecter) Config() *Chain_Config_Call {
	return &Chain_Config_Call{Call: _e.mock.On("Config")}
}

func (_c *Chain_Config_Call) Run(run func()) *Chain_Config_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_Config_Call) Return(_a0 config.ChainScopedConfig) *Chain_Config_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_Config_Call) RunAndReturn(run func() config.ChainScopedConfig) *Chain_Config_Call {
	_c.Call.Return(run)
	return _c
}

// GasEstimator provides a mock function with no fields
func (_m *Chain) GasEstimator() gas.EvmFeeEstimator {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GasEstimator")
	}

	var r0 gas.EvmFeeEstimator
	if rf, ok := ret.Get(0).(func() gas.EvmFeeEstimator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gas.EvmFeeEstimator)
		}
	}

	return r0
}

// Chain_GasEstimator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GasEstimator'
type Chain_GasEstimator_Call struct {
	*mock.Call
}

// GasEstimator is a helper method to define mock.On call
func (_e *Chain_Expecter) GasEstimator() *Chain_GasEstimator_Call {
	return &Chain_GasEstimator_Call{Call: _e.mock.On("GasEstimator")}
}

func (_c *Chain_GasEstimator_Call) Run(run func()) *Chain_GasEstimator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_GasEstimator_Call) Return(_a0 gas.EvmFeeEstimator) *Chain_GasEstimator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_GasEstimator_Call) RunAndReturn(run func() gas.EvmFeeEstimator) *Chain_GasEstimator_Call {
	_c.Call.Return(run)
	return _c
}

// GetChainStatus provides a mock function with given fields: ctx
func (_m *Chain) GetChainStatus(ctx context.Context) (types.ChainStatus, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetChainStatus")
	}

	var r0 types.ChainStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.ChainStatus, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.ChainStatus); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.ChainStatus)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Chain_GetChainStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChainStatus'
type Chain_GetChainStatus_Call struct {
	*mock.Call
}

// GetChainStatus is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Chain_Expecter) GetChainStatus(ctx interface{}) *Chain_GetChainStatus_Call {
	return &Chain_GetChainStatus_Call{Call: _e.mock.On("GetChainStatus", ctx)}
}

func (_c *Chain_GetChainStatus_Call) Run(run func(ctx context.Context)) *Chain_GetChainStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Chain_GetChainStatus_Call) Return(_a0 types.ChainStatus, _a1 error) *Chain_GetChainStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Chain_GetChainStatus_Call) RunAndReturn(run func(context.Context) (types.ChainStatus, error)) *Chain_GetChainStatus_Call {
	_c.Call.Return(run)
	return _c
}

// HeadBroadcaster provides a mock function with no fields
func (_m *Chain) HeadBroadcaster() headtracker.HeadBroadcaster[*evmtypes.Head, common.Hash] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HeadBroadcaster")
	}

	var r0 headtracker.HeadBroadcaster[*evmtypes.Head, common.Hash]
	if rf, ok := ret.Get(0).(func() headtracker.HeadBroadcaster[*evmtypes.Head, common.Hash]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(headtracker.HeadBroadcaster[*evmtypes.Head, common.Hash])
		}
	}

	return r0
}

// Chain_HeadBroadcaster_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeadBroadcaster'
type Chain_HeadBroadcaster_Call struct {
	*mock.Call
}

// HeadBroadcaster is a helper method to define mock.On call
func (_e *Chain_Expecter) HeadBroadcaster() *Chain_HeadBroadcaster_Call {
	return &Chain_HeadBroadcaster_Call{Call: _e.mock.On("HeadBroadcaster")}
}

func (_c *Chain_HeadBroadcaster_Call) Run(run func()) *Chain_HeadBroadcaster_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_HeadBroadcaster_Call) Return(_a0 headtracker.HeadBroadcaster[*evmtypes.Head, common.Hash]) *Chain_HeadBroadcaster_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_HeadBroadcaster_Call) RunAndReturn(run func() headtracker.HeadBroadcaster[*evmtypes.Head, common.Hash]) *Chain_HeadBroadcaster_Call {
	_c.Call.Return(run)
	return _c
}

// HeadTracker provides a mock function with no fields
func (_m *Chain) HeadTracker() headtracker.HeadTracker[*evmtypes.Head, common.Hash] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HeadTracker")
	}

	var r0 headtracker.HeadTracker[*evmtypes.Head, common.Hash]
	if rf, ok := ret.Get(0).(func() headtracker.HeadTracker[*evmtypes.Head, common.Hash]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(headtracker.HeadTracker[*evmtypes.Head, common.Hash])
		}
	}

	return r0
}

// Chain_HeadTracker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeadTracker'
type Chain_HeadTracker_Call struct {
	*mock.Call
}

// HeadTracker is a helper method to define mock.On call
func (_e *Chain_Expecter) HeadTracker() *Chain_HeadTracker_Call {
	return &Chain_HeadTracker_Call{Call: _e.mock.On("HeadTracker")}
}

func (_c *Chain_HeadTracker_Call) Run(run func()) *Chain_HeadTracker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_HeadTracker_Call) Return(_a0 headtracker.HeadTracker[*evmtypes.Head, common.Hash]) *Chain_HeadTracker_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_HeadTracker_Call) RunAndReturn(run func() headtracker.HeadTracker[*evmtypes.Head, common.Hash]) *Chain_HeadTracker_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with no fields
func (_m *Chain) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Chain_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type Chain_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *Chain_Expecter) HealthReport() *Chain_HealthReport_Call {
	return &Chain_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *Chain_HealthReport_Call) Run(run func()) *Chain_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_HealthReport_Call) Return(_a0 map[string]error) *Chain_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_HealthReport_Call) RunAndReturn(run func() map[string]error) *Chain_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with no fields
func (_m *Chain) ID() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// Chain_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type Chain_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *Chain_Expecter) ID() *Chain_ID_Call {
	return &Chain_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *Chain_ID_Call) Run(run func()) *Chain_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_ID_Call) Return(_a0 *big.Int) *Chain_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_ID_Call) RunAndReturn(run func() *big.Int) *Chain_ID_Call {
	_c.Call.Return(run)
	return _c
}

// LatestHead provides a mock function with given fields: ctx
func (_m *Chain) LatestHead(ctx context.Context) (types.Head, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LatestHead")
	}

	var r0 types.Head
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.Head, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.Head); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.Head)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Chain_LatestHead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LatestHead'
type Chain_LatestHead_Call struct {
	*mock.Call
}

// LatestHead is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Chain_Expecter) LatestHead(ctx interface{}) *Chain_LatestHead_Call {
	return &Chain_LatestHead_Call{Call: _e.mock.On("LatestHead", ctx)}
}

func (_c *Chain_LatestHead_Call) Run(run func(ctx context.Context)) *Chain_LatestHead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Chain_LatestHead_Call) Return(_a0 types.Head, _a1 error) *Chain_LatestHead_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Chain_LatestHead_Call) RunAndReturn(run func(context.Context) (types.Head, error)) *Chain_LatestHead_Call {
	_c.Call.Return(run)
	return _c
}

// ListNodeStatuses provides a mock function with given fields: ctx, pageSize, pageToken
func (_m *Chain) ListNodeStatuses(ctx context.Context, pageSize int32, pageToken string) ([]types.NodeStatus, string, int, error) {
	ret := _m.Called(ctx, pageSize, pageToken)

	if len(ret) == 0 {
		panic("no return value specified for ListNodeStatuses")
	}

	var r0 []types.NodeStatus
	var r1 string
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) ([]types.NodeStatus, string, int, error)); ok {
		return rf(ctx, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) []types.NodeStatus); ok {
		r0 = rf(ctx, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NodeStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, string) string); ok {
		r1 = rf(ctx, pageSize, pageToken)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int32, string) int); ok {
		r2 = rf(ctx, pageSize, pageToken)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(context.Context, int32, string) error); ok {
		r3 = rf(ctx, pageSize, pageToken)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Chain_ListNodeStatuses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListNodeStatuses'
type Chain_ListNodeStatuses_Call struct {
	*mock.Call
}

// ListNodeStatuses is a helper method to define mock.On call
//   - ctx context.Context
//   - pageSize int32
//   - pageToken string
func (_e *Chain_Expecter) ListNodeStatuses(ctx interface{}, pageSize interface{}, pageToken interface{}) *Chain_ListNodeStatuses_Call {
	return &Chain_ListNodeStatuses_Call{Call: _e.mock.On("ListNodeStatuses", ctx, pageSize, pageToken)}
}

func (_c *Chain_ListNodeStatuses_Call) Run(run func(ctx context.Context, pageSize int32, pageToken string)) *Chain_ListNodeStatuses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(string))
	})
	return _c
}

func (_c *Chain_ListNodeStatuses_Call) Return(stats []types.NodeStatus, nextPageToken string, total int, err error) *Chain_ListNodeStatuses_Call {
	_c.Call.Return(stats, nextPageToken, total, err)
	return _c
}

func (_c *Chain_ListNodeStatuses_Call) RunAndReturn(run func(context.Context, int32, string) ([]types.NodeStatus, string, int, error)) *Chain_ListNodeStatuses_Call {
	_c.Call.Return(run)
	return _c
}

// LogBroadcaster provides a mock function with no fields
func (_m *Chain) LogBroadcaster() log.Broadcaster {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LogBroadcaster")
	}

	var r0 log.Broadcaster
	if rf, ok := ret.Get(0).(func() log.Broadcaster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Broadcaster)
		}
	}

	return r0
}

// Chain_LogBroadcaster_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogBroadcaster'
type Chain_LogBroadcaster_Call struct {
	*mock.Call
}

// LogBroadcaster is a helper method to define mock.On call
func (_e *Chain_Expecter) LogBroadcaster() *Chain_LogBroadcaster_Call {
	return &Chain_LogBroadcaster_Call{Call: _e.mock.On("LogBroadcaster")}
}

func (_c *Chain_LogBroadcaster_Call) Run(run func()) *Chain_LogBroadcaster_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_LogBroadcaster_Call) Return(_a0 log.Broadcaster) *Chain_LogBroadcaster_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_LogBroadcaster_Call) RunAndReturn(run func() log.Broadcaster) *Chain_LogBroadcaster_Call {
	_c.Call.Return(run)
	return _c
}

// LogPoller provides a mock function with no fields
func (_m *Chain) LogPoller() logpoller.LogPoller {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LogPoller")
	}

	var r0 logpoller.LogPoller
	if rf, ok := ret.Get(0).(func() logpoller.LogPoller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logpoller.LogPoller)
		}
	}

	return r0
}

// Chain_LogPoller_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogPoller'
type Chain_LogPoller_Call struct {
	*mock.Call
}

// LogPoller is a helper method to define mock.On call
func (_e *Chain_Expecter) LogPoller() *Chain_LogPoller_Call {
	return &Chain_LogPoller_Call{Call: _e.mock.On("LogPoller")}
}

func (_c *Chain_LogPoller_Call) Run(run func()) *Chain_LogPoller_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_LogPoller_Call) Return(_a0 logpoller.LogPoller) *Chain_LogPoller_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_LogPoller_Call) RunAndReturn(run func() logpoller.LogPoller) *Chain_LogPoller_Call {
	_c.Call.Return(run)
	return _c
}

// Logger provides a mock function with no fields
func (_m *Chain) Logger() logger.Logger {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Logger")
	}

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func() logger.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// Chain_Logger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logger'
type Chain_Logger_Call struct {
	*mock.Call
}

// Logger is a helper method to define mock.On call
func (_e *Chain_Expecter) Logger() *Chain_Logger_Call {
	return &Chain_Logger_Call{Call: _e.mock.On("Logger")}
}

func (_c *Chain_Logger_Call) Run(run func()) *Chain_Logger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_Logger_Call) Return(_a0 logger.Logger) *Chain_Logger_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_Logger_Call) RunAndReturn(run func() logger.Logger) *Chain_Logger_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with no fields
func (_m *Chain) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Chain_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Chain_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Chain_Expecter) Name() *Chain_Name_Call {
	return &Chain_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Chain_Name_Call) Run(run func()) *Chain_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_Name_Call) Return(_a0 string) *Chain_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_Name_Call) RunAndReturn(run func() string) *Chain_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with no fields
func (_m *Chain) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chain_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type Chain_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *Chain_Expecter) Ready() *Chain_Ready_Call {
	return &Chain_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *Chain_Ready_Call) Run(run func()) *Chain_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_Ready_Call) Return(_a0 error) *Chain_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_Ready_Call) RunAndReturn(run func() error) *Chain_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *Chain) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chain_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Chain_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Chain_Expecter) Start(_a0 interface{}) *Chain_Start_Call {
	return &Chain_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *Chain_Start_Call) Run(run func(_a0 context.Context)) *Chain_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Chain_Start_Call) Return(_a0 error) *Chain_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_Start_Call) RunAndReturn(run func(context.Context) error) *Chain_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Transact provides a mock function with given fields: ctx, from, to, amount, balanceCheck
func (_m *Chain) Transact(ctx context.Context, from string, to string, amount *big.Int, balanceCheck bool) error {
	ret := _m.Called(ctx, from, to, amount, balanceCheck)

	if len(ret) == 0 {
		panic("no return value specified for Transact")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *big.Int, bool) error); ok {
		r0 = rf(ctx, from, to, amount, balanceCheck)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chain_Transact_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transact'
type Chain_Transact_Call struct {
	*mock.Call
}

// Transact is a helper method to define mock.On call
//   - ctx context.Context
//   - from string
//   - to string
//   - amount *big.Int
//   - balanceCheck bool
func (_e *Chain_Expecter) Transact(ctx interface{}, from interface{}, to interface{}, amount interface{}, balanceCheck interface{}) *Chain_Transact_Call {
	return &Chain_Transact_Call{Call: _e.mock.On("Transact", ctx, from, to, amount, balanceCheck)}
}

func (_c *Chain_Transact_Call) Run(run func(ctx context.Context, from string, to string, amount *big.Int, balanceCheck bool)) *Chain_Transact_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*big.Int), args[4].(bool))
	})
	return _c
}

func (_c *Chain_Transact_Call) Return(_a0 error) *Chain_Transact_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_Transact_Call) RunAndReturn(run func(context.Context, string, string, *big.Int, bool) error) *Chain_Transact_Call {
	_c.Call.Return(run)
	return _c
}

// TxManager provides a mock function with no fields
func (_m *Chain) TxManager() txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TxManager")
	}

	var r0 txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee]
	if rf, ok := ret.Get(0).(func() txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee])
		}
	}

	return r0
}

// Chain_TxManager_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxManager'
type Chain_TxManager_Call struct {
	*mock.Call
}

// TxManager is a helper method to define mock.On call
func (_e *Chain_Expecter) TxManager() *Chain_TxManager_Call {
	return &Chain_TxManager_Call{Call: _e.mock.On("TxManager")}
}

func (_c *Chain_TxManager_Call) Run(run func()) *Chain_TxManager_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Chain_TxManager_Call) Return(_a0 txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee]) *Chain_TxManager_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Chain_TxManager_Call) RunAndReturn(run func() txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee]) *Chain_TxManager_Call {
	_c.Call.Return(run)
	return _c
}

// NewChain creates a new instance of Chain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChain(t interface {
	mock.TestingT
	Cleanup(func())
}) *Chain {
	mock := &Chain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
