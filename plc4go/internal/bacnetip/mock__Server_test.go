/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.30.1. DO NOT EDIT.

package bacnetip

import mock "github.com/stretchr/testify/mock"

// mock_Server is an autogenerated mock type for the _Server type
type mock_Server struct {
	mock.Mock
}

type mock_Server_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_Server) EXPECT() *mock_Server_Expecter {
	return &mock_Server_Expecter{mock: &_m.Mock}
}

// Indication provides a mock function with given fields: pdu
func (_m *mock_Server) Indication(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_Server_Indication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Indication'
type mock_Server_Indication_Call struct {
	*mock.Call
}

// Indication is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_Server_Expecter) Indication(pdu interface{}) *mock_Server_Indication_Call {
	return &mock_Server_Indication_Call{Call: _e.mock.On("Indication", pdu)}
}

func (_c *mock_Server_Indication_Call) Run(run func(pdu _PDU)) *mock_Server_Indication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_Server_Indication_Call) Return(_a0 error) *mock_Server_Indication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_Server_Indication_Call) RunAndReturn(run func(_PDU) error) *mock_Server_Indication_Call {
	_c.Call.Return(run)
	return _c
}

// Response provides a mock function with given fields: pdu
func (_m *mock_Server) Response(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_Server_Response_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Response'
type mock_Server_Response_Call struct {
	*mock.Call
}

// Response is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_Server_Expecter) Response(pdu interface{}) *mock_Server_Response_Call {
	return &mock_Server_Response_Call{Call: _e.mock.On("Response", pdu)}
}

func (_c *mock_Server_Response_Call) Run(run func(pdu _PDU)) *mock_Server_Response_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_Server_Response_Call) Return(_a0 error) *mock_Server_Response_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_Server_Response_Call) RunAndReturn(run func(_PDU) error) *mock_Server_Response_Call {
	_c.Call.Return(run)
	return _c
}

// _setServerPeer provides a mock function with given fields: serverPeer
func (_m *mock_Server) _setServerPeer(serverPeer _Client) {
	_m.Called(serverPeer)
}

// mock_Server__setServerPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setServerPeer'
type mock_Server__setServerPeer_Call struct {
	*mock.Call
}

// _setServerPeer is a helper method to define mock.On call
//   - serverPeer _Client
func (_e *mock_Server_Expecter) _setServerPeer(serverPeer interface{}) *mock_Server__setServerPeer_Call {
	return &mock_Server__setServerPeer_Call{Call: _e.mock.On("_setServerPeer", serverPeer)}
}

func (_c *mock_Server__setServerPeer_Call) Run(run func(serverPeer _Client)) *mock_Server__setServerPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_Client))
	})
	return _c
}

func (_c *mock_Server__setServerPeer_Call) Return() *mock_Server__setServerPeer_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_Server__setServerPeer_Call) RunAndReturn(run func(_Client)) *mock_Server__setServerPeer_Call {
	_c.Call.Return(run)
	return _c
}

// getServerId provides a mock function with given fields:
func (_m *mock_Server) getServerId() *int {
	ret := _m.Called()

	var r0 *int
	if rf, ok := ret.Get(0).(func() *int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	return r0
}

// mock_Server_getServerId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getServerId'
type mock_Server_getServerId_Call struct {
	*mock.Call
}

// getServerId is a helper method to define mock.On call
func (_e *mock_Server_Expecter) getServerId() *mock_Server_getServerId_Call {
	return &mock_Server_getServerId_Call{Call: _e.mock.On("getServerId")}
}

func (_c *mock_Server_getServerId_Call) Run(run func()) *mock_Server_getServerId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_Server_getServerId_Call) Return(_a0 *int) *mock_Server_getServerId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_Server_getServerId_Call) RunAndReturn(run func() *int) *mock_Server_getServerId_Call {
	_c.Call.Return(run)
	return _c
}

// newMock_Server creates a new instance of mock_Server. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMock_Server(t interface {
	mock.TestingT
	Cleanup(func())
}) *mock_Server {
	mock := &mock_Server{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
