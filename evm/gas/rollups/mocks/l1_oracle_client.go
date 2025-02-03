// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/ethereum/go-ethereum/rpc"
)

// L1OracleClient is an autogenerated mock type for the l1OracleClient type
type L1OracleClient struct {
	mock.Mock
}

type L1OracleClient_Expecter struct {
	mock *mock.Mock
}

func (_m *L1OracleClient) EXPECT() *L1OracleClient_Expecter {
	return &L1OracleClient_Expecter{mock: &_m.Mock}
}

// BatchCallContext provides a mock function with given fields: ctx, b
func (_m *L1OracleClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	ret := _m.Called(ctx, b)

	if len(ret) == 0 {
		panic("no return value specified for BatchCallContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []rpc.BatchElem) error); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// L1OracleClient_BatchCallContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchCallContext'
type L1OracleClient_BatchCallContext_Call struct {
	*mock.Call
}

// BatchCallContext is a helper method to define mock.On call
//   - ctx context.Context
//   - b []rpc.BatchElem
func (_e *L1OracleClient_Expecter) BatchCallContext(ctx interface{}, b interface{}) *L1OracleClient_BatchCallContext_Call {
	return &L1OracleClient_BatchCallContext_Call{Call: _e.mock.On("BatchCallContext", ctx, b)}
}

func (_c *L1OracleClient_BatchCallContext_Call) Run(run func(ctx context.Context, b []rpc.BatchElem)) *L1OracleClient_BatchCallContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]rpc.BatchElem))
	})
	return _c
}

func (_c *L1OracleClient_BatchCallContext_Call) Return(_a0 error) *L1OracleClient_BatchCallContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1OracleClient_BatchCallContext_Call) RunAndReturn(run func(context.Context, []rpc.BatchElem) error) *L1OracleClient_BatchCallContext_Call {
	_c.Call.Return(run)
	return _c
}

// CallContract provides a mock function with given fields: ctx, msg, blockNumber
func (_m *L1OracleClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, msg, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for CallContract")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)); ok {
		return rf(ctx, msg, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, msg, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, msg, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L1OracleClient_CallContract_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CallContract'
type L1OracleClient_CallContract_Call struct {
	*mock.Call
}

// CallContract is a helper method to define mock.On call
//   - ctx context.Context
//   - msg ethereum.CallMsg
//   - blockNumber *big.Int
func (_e *L1OracleClient_Expecter) CallContract(ctx interface{}, msg interface{}, blockNumber interface{}) *L1OracleClient_CallContract_Call {
	return &L1OracleClient_CallContract_Call{Call: _e.mock.On("CallContract", ctx, msg, blockNumber)}
}

func (_c *L1OracleClient_CallContract_Call) Run(run func(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int)) *L1OracleClient_CallContract_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.CallMsg), args[2].(*big.Int))
	})
	return _c
}

func (_c *L1OracleClient_CallContract_Call) Return(_a0 []byte, _a1 error) *L1OracleClient_CallContract_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L1OracleClient_CallContract_Call) RunAndReturn(run func(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)) *L1OracleClient_CallContract_Call {
	_c.Call.Return(run)
	return _c
}

// NewL1OracleClient creates a new instance of L1OracleClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewL1OracleClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *L1OracleClient {
	mock := &L1OracleClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
