// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	assets "github.com/smartcontractkit/chainlink/v2/evm/assets"

	fees "github.com/smartcontractkit/chainlink-framework/chains/fees"

	gas "github.com/smartcontractkit/chainlink/v2/evm/gas"

	mock "github.com/stretchr/testify/mock"

	rollups "github.com/smartcontractkit/chainlink/v2/evm/gas/rollups"

	types "github.com/smartcontractkit/chainlink/v2/evm/types"
)

// EvmEstimator is an autogenerated mock type for the EvmEstimator type
type EvmEstimator struct {
	mock.Mock
}

type EvmEstimator_Expecter struct {
	mock *mock.Mock
}

func (_m *EvmEstimator) EXPECT() *EvmEstimator_Expecter {
	return &EvmEstimator_Expecter{mock: &_m.Mock}
}

// BumpDynamicFee provides a mock function with given fields: ctx, original, maxGasPriceWei, attempts
func (_m *EvmEstimator) BumpDynamicFee(ctx context.Context, original gas.DynamicFee, maxGasPriceWei *assets.Wei, attempts []gas.EvmPriorAttempt) (gas.DynamicFee, error) {
	ret := _m.Called(ctx, original, maxGasPriceWei, attempts)

	if len(ret) == 0 {
		panic("no return value specified for BumpDynamicFee")
	}

	var r0 gas.DynamicFee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, gas.DynamicFee, *assets.Wei, []gas.EvmPriorAttempt) (gas.DynamicFee, error)); ok {
		return rf(ctx, original, maxGasPriceWei, attempts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, gas.DynamicFee, *assets.Wei, []gas.EvmPriorAttempt) gas.DynamicFee); ok {
		r0 = rf(ctx, original, maxGasPriceWei, attempts)
	} else {
		r0 = ret.Get(0).(gas.DynamicFee)
	}

	if rf, ok := ret.Get(1).(func(context.Context, gas.DynamicFee, *assets.Wei, []gas.EvmPriorAttempt) error); ok {
		r1 = rf(ctx, original, maxGasPriceWei, attempts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EvmEstimator_BumpDynamicFee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BumpDynamicFee'
type EvmEstimator_BumpDynamicFee_Call struct {
	*mock.Call
}

// BumpDynamicFee is a helper method to define mock.On call
//   - ctx context.Context
//   - original gas.DynamicFee
//   - maxGasPriceWei *assets.Wei
//   - attempts []gas.EvmPriorAttempt
func (_e *EvmEstimator_Expecter) BumpDynamicFee(ctx interface{}, original interface{}, maxGasPriceWei interface{}, attempts interface{}) *EvmEstimator_BumpDynamicFee_Call {
	return &EvmEstimator_BumpDynamicFee_Call{Call: _e.mock.On("BumpDynamicFee", ctx, original, maxGasPriceWei, attempts)}
}

func (_c *EvmEstimator_BumpDynamicFee_Call) Run(run func(ctx context.Context, original gas.DynamicFee, maxGasPriceWei *assets.Wei, attempts []gas.EvmPriorAttempt)) *EvmEstimator_BumpDynamicFee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(gas.DynamicFee), args[2].(*assets.Wei), args[3].([]gas.EvmPriorAttempt))
	})
	return _c
}

func (_c *EvmEstimator_BumpDynamicFee_Call) Return(bumped gas.DynamicFee, err error) *EvmEstimator_BumpDynamicFee_Call {
	_c.Call.Return(bumped, err)
	return _c
}

func (_c *EvmEstimator_BumpDynamicFee_Call) RunAndReturn(run func(context.Context, gas.DynamicFee, *assets.Wei, []gas.EvmPriorAttempt) (gas.DynamicFee, error)) *EvmEstimator_BumpDynamicFee_Call {
	_c.Call.Return(run)
	return _c
}

// BumpLegacyGas provides a mock function with given fields: ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts
func (_m *EvmEstimator) BumpLegacyGas(ctx context.Context, originalGasPrice *assets.Wei, gasLimit uint64, maxGasPriceWei *assets.Wei, attempts []gas.EvmPriorAttempt) (*assets.Wei, uint64, error) {
	ret := _m.Called(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)

	if len(ret) == 0 {
		panic("no return value specified for BumpLegacyGas")
	}

	var r0 *assets.Wei
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *assets.Wei, uint64, *assets.Wei, []gas.EvmPriorAttempt) (*assets.Wei, uint64, error)); ok {
		return rf(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *assets.Wei, uint64, *assets.Wei, []gas.EvmPriorAttempt) *assets.Wei); ok {
		r0 = rf(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *assets.Wei, uint64, *assets.Wei, []gas.EvmPriorAttempt) uint64); ok {
		r1 = rf(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *assets.Wei, uint64, *assets.Wei, []gas.EvmPriorAttempt) error); ok {
		r2 = rf(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EvmEstimator_BumpLegacyGas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BumpLegacyGas'
type EvmEstimator_BumpLegacyGas_Call struct {
	*mock.Call
}

// BumpLegacyGas is a helper method to define mock.On call
//   - ctx context.Context
//   - originalGasPrice *assets.Wei
//   - gasLimit uint64
//   - maxGasPriceWei *assets.Wei
//   - attempts []gas.EvmPriorAttempt
func (_e *EvmEstimator_Expecter) BumpLegacyGas(ctx interface{}, originalGasPrice interface{}, gasLimit interface{}, maxGasPriceWei interface{}, attempts interface{}) *EvmEstimator_BumpLegacyGas_Call {
	return &EvmEstimator_BumpLegacyGas_Call{Call: _e.mock.On("BumpLegacyGas", ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)}
}

func (_c *EvmEstimator_BumpLegacyGas_Call) Run(run func(ctx context.Context, originalGasPrice *assets.Wei, gasLimit uint64, maxGasPriceWei *assets.Wei, attempts []gas.EvmPriorAttempt)) *EvmEstimator_BumpLegacyGas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*assets.Wei), args[2].(uint64), args[3].(*assets.Wei), args[4].([]gas.EvmPriorAttempt))
	})
	return _c
}

func (_c *EvmEstimator_BumpLegacyGas_Call) Return(bumpedGasPrice *assets.Wei, chainSpecificGasLimit uint64, err error) *EvmEstimator_BumpLegacyGas_Call {
	_c.Call.Return(bumpedGasPrice, chainSpecificGasLimit, err)
	return _c
}

func (_c *EvmEstimator_BumpLegacyGas_Call) RunAndReturn(run func(context.Context, *assets.Wei, uint64, *assets.Wei, []gas.EvmPriorAttempt) (*assets.Wei, uint64, error)) *EvmEstimator_BumpLegacyGas_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with no fields
func (_m *EvmEstimator) Close() error {
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

// EvmEstimator_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type EvmEstimator_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *EvmEstimator_Expecter) Close() *EvmEstimator_Close_Call {
	return &EvmEstimator_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *EvmEstimator_Close_Call) Run(run func()) *EvmEstimator_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmEstimator_Close_Call) Return(_a0 error) *EvmEstimator_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmEstimator_Close_Call) RunAndReturn(run func() error) *EvmEstimator_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetDynamicFee provides a mock function with given fields: ctx, maxGasPriceWei
func (_m *EvmEstimator) GetDynamicFee(ctx context.Context, maxGasPriceWei *assets.Wei) (gas.DynamicFee, error) {
	ret := _m.Called(ctx, maxGasPriceWei)

	if len(ret) == 0 {
		panic("no return value specified for GetDynamicFee")
	}

	var r0 gas.DynamicFee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *assets.Wei) (gas.DynamicFee, error)); ok {
		return rf(ctx, maxGasPriceWei)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *assets.Wei) gas.DynamicFee); ok {
		r0 = rf(ctx, maxGasPriceWei)
	} else {
		r0 = ret.Get(0).(gas.DynamicFee)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *assets.Wei) error); ok {
		r1 = rf(ctx, maxGasPriceWei)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EvmEstimator_GetDynamicFee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDynamicFee'
type EvmEstimator_GetDynamicFee_Call struct {
	*mock.Call
}

// GetDynamicFee is a helper method to define mock.On call
//   - ctx context.Context
//   - maxGasPriceWei *assets.Wei
func (_e *EvmEstimator_Expecter) GetDynamicFee(ctx interface{}, maxGasPriceWei interface{}) *EvmEstimator_GetDynamicFee_Call {
	return &EvmEstimator_GetDynamicFee_Call{Call: _e.mock.On("GetDynamicFee", ctx, maxGasPriceWei)}
}

func (_c *EvmEstimator_GetDynamicFee_Call) Run(run func(ctx context.Context, maxGasPriceWei *assets.Wei)) *EvmEstimator_GetDynamicFee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*assets.Wei))
	})
	return _c
}

func (_c *EvmEstimator_GetDynamicFee_Call) Return(fee gas.DynamicFee, err error) *EvmEstimator_GetDynamicFee_Call {
	_c.Call.Return(fee, err)
	return _c
}

func (_c *EvmEstimator_GetDynamicFee_Call) RunAndReturn(run func(context.Context, *assets.Wei) (gas.DynamicFee, error)) *EvmEstimator_GetDynamicFee_Call {
	_c.Call.Return(run)
	return _c
}

// GetLegacyGas provides a mock function with given fields: ctx, calldata, gasLimit, maxGasPriceWei, opts
func (_m *EvmEstimator) GetLegacyGas(ctx context.Context, calldata []byte, gasLimit uint64, maxGasPriceWei *assets.Wei, opts ...fees.Opt) (*assets.Wei, uint64, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, calldata, gasLimit, maxGasPriceWei)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLegacyGas")
	}

	var r0 *assets.Wei
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint64, *assets.Wei, ...fees.Opt) (*assets.Wei, uint64, error)); ok {
		return rf(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint64, *assets.Wei, ...fees.Opt) *assets.Wei); ok {
		r0 = rf(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, uint64, *assets.Wei, ...fees.Opt) uint64); ok {
		r1 = rf(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, []byte, uint64, *assets.Wei, ...fees.Opt) error); ok {
		r2 = rf(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EvmEstimator_GetLegacyGas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLegacyGas'
type EvmEstimator_GetLegacyGas_Call struct {
	*mock.Call
}

// GetLegacyGas is a helper method to define mock.On call
//   - ctx context.Context
//   - calldata []byte
//   - gasLimit uint64
//   - maxGasPriceWei *assets.Wei
//   - opts ...fees.Opt
func (_e *EvmEstimator_Expecter) GetLegacyGas(ctx interface{}, calldata interface{}, gasLimit interface{}, maxGasPriceWei interface{}, opts ...interface{}) *EvmEstimator_GetLegacyGas_Call {
	return &EvmEstimator_GetLegacyGas_Call{Call: _e.mock.On("GetLegacyGas",
		append([]interface{}{ctx, calldata, gasLimit, maxGasPriceWei}, opts...)...)}
}

func (_c *EvmEstimator_GetLegacyGas_Call) Run(run func(ctx context.Context, calldata []byte, gasLimit uint64, maxGasPriceWei *assets.Wei, opts ...fees.Opt)) *EvmEstimator_GetLegacyGas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]fees.Opt, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(fees.Opt)
			}
		}
		run(args[0].(context.Context), args[1].([]byte), args[2].(uint64), args[3].(*assets.Wei), variadicArgs...)
	})
	return _c
}

func (_c *EvmEstimator_GetLegacyGas_Call) Return(gasPrice *assets.Wei, chainSpecificGasLimit uint64, err error) *EvmEstimator_GetLegacyGas_Call {
	_c.Call.Return(gasPrice, chainSpecificGasLimit, err)
	return _c
}

func (_c *EvmEstimator_GetLegacyGas_Call) RunAndReturn(run func(context.Context, []byte, uint64, *assets.Wei, ...fees.Opt) (*assets.Wei, uint64, error)) *EvmEstimator_GetLegacyGas_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with no fields
func (_m *EvmEstimator) HealthReport() map[string]error {
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

// EvmEstimator_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type EvmEstimator_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *EvmEstimator_Expecter) HealthReport() *EvmEstimator_HealthReport_Call {
	return &EvmEstimator_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *EvmEstimator_HealthReport_Call) Run(run func()) *EvmEstimator_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmEstimator_HealthReport_Call) Return(_a0 map[string]error) *EvmEstimator_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmEstimator_HealthReport_Call) RunAndReturn(run func() map[string]error) *EvmEstimator_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// L1Oracle provides a mock function with no fields
func (_m *EvmEstimator) L1Oracle() rollups.L1Oracle {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for L1Oracle")
	}

	var r0 rollups.L1Oracle
	if rf, ok := ret.Get(0).(func() rollups.L1Oracle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rollups.L1Oracle)
		}
	}

	return r0
}

// EvmEstimator_L1Oracle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'L1Oracle'
type EvmEstimator_L1Oracle_Call struct {
	*mock.Call
}

// L1Oracle is a helper method to define mock.On call
func (_e *EvmEstimator_Expecter) L1Oracle() *EvmEstimator_L1Oracle_Call {
	return &EvmEstimator_L1Oracle_Call{Call: _e.mock.On("L1Oracle")}
}

func (_c *EvmEstimator_L1Oracle_Call) Run(run func()) *EvmEstimator_L1Oracle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmEstimator_L1Oracle_Call) Return(_a0 rollups.L1Oracle) *EvmEstimator_L1Oracle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmEstimator_L1Oracle_Call) RunAndReturn(run func() rollups.L1Oracle) *EvmEstimator_L1Oracle_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with no fields
func (_m *EvmEstimator) Name() string {
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

// EvmEstimator_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type EvmEstimator_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *EvmEstimator_Expecter) Name() *EvmEstimator_Name_Call {
	return &EvmEstimator_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *EvmEstimator_Name_Call) Run(run func()) *EvmEstimator_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmEstimator_Name_Call) Return(_a0 string) *EvmEstimator_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmEstimator_Name_Call) RunAndReturn(run func() string) *EvmEstimator_Name_Call {
	_c.Call.Return(run)
	return _c
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *EvmEstimator) OnNewLongestChain(ctx context.Context, head *types.Head) {
	_m.Called(ctx, head)
}

// EvmEstimator_OnNewLongestChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnNewLongestChain'
type EvmEstimator_OnNewLongestChain_Call struct {
	*mock.Call
}

// OnNewLongestChain is a helper method to define mock.On call
//   - ctx context.Context
//   - head *types.Head
func (_e *EvmEstimator_Expecter) OnNewLongestChain(ctx interface{}, head interface{}) *EvmEstimator_OnNewLongestChain_Call {
	return &EvmEstimator_OnNewLongestChain_Call{Call: _e.mock.On("OnNewLongestChain", ctx, head)}
}

func (_c *EvmEstimator_OnNewLongestChain_Call) Run(run func(ctx context.Context, head *types.Head)) *EvmEstimator_OnNewLongestChain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Head))
	})
	return _c
}

func (_c *EvmEstimator_OnNewLongestChain_Call) Return() *EvmEstimator_OnNewLongestChain_Call {
	_c.Call.Return()
	return _c
}

func (_c *EvmEstimator_OnNewLongestChain_Call) RunAndReturn(run func(context.Context, *types.Head)) *EvmEstimator_OnNewLongestChain_Call {
	_c.Run(run)
	return _c
}

// Ready provides a mock function with no fields
func (_m *EvmEstimator) Ready() error {
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

// EvmEstimator_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type EvmEstimator_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *EvmEstimator_Expecter) Ready() *EvmEstimator_Ready_Call {
	return &EvmEstimator_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *EvmEstimator_Ready_Call) Run(run func()) *EvmEstimator_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmEstimator_Ready_Call) Return(_a0 error) *EvmEstimator_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmEstimator_Ready_Call) RunAndReturn(run func() error) *EvmEstimator_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *EvmEstimator) Start(_a0 context.Context) error {
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

// EvmEstimator_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type EvmEstimator_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *EvmEstimator_Expecter) Start(_a0 interface{}) *EvmEstimator_Start_Call {
	return &EvmEstimator_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *EvmEstimator_Start_Call) Run(run func(_a0 context.Context)) *EvmEstimator_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EvmEstimator_Start_Call) Return(_a0 error) *EvmEstimator_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmEstimator_Start_Call) RunAndReturn(run func(context.Context) error) *EvmEstimator_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewEvmEstimator creates a new instance of EvmEstimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvmEstimator(t interface {
	mock.TestingT
	Cleanup(func())
}) *EvmEstimator {
	mock := &EvmEstimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
