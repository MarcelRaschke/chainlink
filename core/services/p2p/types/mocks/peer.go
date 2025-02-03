// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
)

// Peer is an autogenerated mock type for the Peer type
type Peer struct {
	mock.Mock
}

type Peer_Expecter struct {
	mock *mock.Mock
}

func (_m *Peer) EXPECT() *Peer_Expecter {
	return &Peer_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with no fields
func (_m *Peer) Close() error {
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

// Peer_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Peer_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Peer_Expecter) Close() *Peer_Close_Call {
	return &Peer_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Peer_Close_Call) Run(run func()) *Peer_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Peer_Close_Call) Return(_a0 error) *Peer_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_Close_Call) RunAndReturn(run func() error) *Peer_Close_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with no fields
func (_m *Peer) HealthReport() map[string]error {
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

// Peer_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type Peer_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *Peer_Expecter) HealthReport() *Peer_HealthReport_Call {
	return &Peer_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *Peer_HealthReport_Call) Run(run func()) *Peer_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Peer_HealthReport_Call) Return(_a0 map[string]error) *Peer_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_HealthReport_Call) RunAndReturn(run func() map[string]error) *Peer_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with no fields
func (_m *Peer) ID() ragep2ptypes.PeerID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 ragep2ptypes.PeerID
	if rf, ok := ret.Get(0).(func() ragep2ptypes.PeerID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ragep2ptypes.PeerID)
		}
	}

	return r0
}

// Peer_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type Peer_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *Peer_Expecter) ID() *Peer_ID_Call {
	return &Peer_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *Peer_ID_Call) Run(run func()) *Peer_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Peer_ID_Call) Return(_a0 ragep2ptypes.PeerID) *Peer_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_ID_Call) RunAndReturn(run func() ragep2ptypes.PeerID) *Peer_ID_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with no fields
func (_m *Peer) Name() string {
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

// Peer_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Peer_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Peer_Expecter) Name() *Peer_Name_Call {
	return &Peer_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Peer_Name_Call) Run(run func()) *Peer_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Peer_Name_Call) Return(_a0 string) *Peer_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_Name_Call) RunAndReturn(run func() string) *Peer_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with no fields
func (_m *Peer) Ready() error {
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

// Peer_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type Peer_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *Peer_Expecter) Ready() *Peer_Ready_Call {
	return &Peer_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *Peer_Ready_Call) Run(run func()) *Peer_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Peer_Ready_Call) Return(_a0 error) *Peer_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_Ready_Call) RunAndReturn(run func() error) *Peer_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Receive provides a mock function with no fields
func (_m *Peer) Receive() <-chan types.Message {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Receive")
	}

	var r0 <-chan types.Message
	if rf, ok := ret.Get(0).(func() <-chan types.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan types.Message)
		}
	}

	return r0
}

// Peer_Receive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Receive'
type Peer_Receive_Call struct {
	*mock.Call
}

// Receive is a helper method to define mock.On call
func (_e *Peer_Expecter) Receive() *Peer_Receive_Call {
	return &Peer_Receive_Call{Call: _e.mock.On("Receive")}
}

func (_c *Peer_Receive_Call) Run(run func()) *Peer_Receive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Peer_Receive_Call) Return(_a0 <-chan types.Message) *Peer_Receive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_Receive_Call) RunAndReturn(run func() <-chan types.Message) *Peer_Receive_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: peerID, msg
func (_m *Peer) Send(peerID ragep2ptypes.PeerID, msg []byte) error {
	ret := _m.Called(peerID, msg)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(ragep2ptypes.PeerID, []byte) error); ok {
		r0 = rf(peerID, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Peer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type Peer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - peerID ragep2ptypes.PeerID
//   - msg []byte
func (_e *Peer_Expecter) Send(peerID interface{}, msg interface{}) *Peer_Send_Call {
	return &Peer_Send_Call{Call: _e.mock.On("Send", peerID, msg)}
}

func (_c *Peer_Send_Call) Run(run func(peerID ragep2ptypes.PeerID, msg []byte)) *Peer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ragep2ptypes.PeerID), args[1].([]byte))
	})
	return _c
}

func (_c *Peer_Send_Call) Return(_a0 error) *Peer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_Send_Call) RunAndReturn(run func(ragep2ptypes.PeerID, []byte) error) *Peer_Send_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *Peer) Start(_a0 context.Context) error {
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

// Peer_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Peer_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Peer_Expecter) Start(_a0 interface{}) *Peer_Start_Call {
	return &Peer_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *Peer_Start_Call) Run(run func(_a0 context.Context)) *Peer_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Peer_Start_Call) Return(_a0 error) *Peer_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_Start_Call) RunAndReturn(run func(context.Context) error) *Peer_Start_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateConnections provides a mock function with given fields: peers
func (_m *Peer) UpdateConnections(peers map[ragep2ptypes.PeerID]types.StreamConfig) error {
	ret := _m.Called(peers)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConnections")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(map[ragep2ptypes.PeerID]types.StreamConfig) error); ok {
		r0 = rf(peers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Peer_UpdateConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateConnections'
type Peer_UpdateConnections_Call struct {
	*mock.Call
}

// UpdateConnections is a helper method to define mock.On call
//   - peers map[ragep2ptypes.PeerID]types.StreamConfig
func (_e *Peer_Expecter) UpdateConnections(peers interface{}) *Peer_UpdateConnections_Call {
	return &Peer_UpdateConnections_Call{Call: _e.mock.On("UpdateConnections", peers)}
}

func (_c *Peer_UpdateConnections_Call) Run(run func(peers map[ragep2ptypes.PeerID]types.StreamConfig)) *Peer_UpdateConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[ragep2ptypes.PeerID]types.StreamConfig))
	})
	return _c
}

func (_c *Peer_UpdateConnections_Call) Return(_a0 error) *Peer_UpdateConnections_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peer_UpdateConnections_Call) RunAndReturn(run func(map[ragep2ptypes.PeerID]types.StreamConfig) error) *Peer_UpdateConnections_Call {
	_c.Call.Return(run)
	return _c
}

// NewPeer creates a new instance of Peer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPeer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Peer {
	mock := &Peer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
