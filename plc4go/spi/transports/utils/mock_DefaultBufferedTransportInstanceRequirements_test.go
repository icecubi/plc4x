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

package utils

import (
	bufio "bufio"

	mock "github.com/stretchr/testify/mock"
)

// MockDefaultBufferedTransportInstanceRequirements is an autogenerated mock type for the DefaultBufferedTransportInstanceRequirements type
type MockDefaultBufferedTransportInstanceRequirements struct {
	mock.Mock
}

type MockDefaultBufferedTransportInstanceRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDefaultBufferedTransportInstanceRequirements) EXPECT() *MockDefaultBufferedTransportInstanceRequirements_Expecter {
	return &MockDefaultBufferedTransportInstanceRequirements_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields:
func (_m *MockDefaultBufferedTransportInstanceRequirements) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultBufferedTransportInstanceRequirements_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockDefaultBufferedTransportInstanceRequirements_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockDefaultBufferedTransportInstanceRequirements_Expecter) Connect() *MockDefaultBufferedTransportInstanceRequirements_Connect_Call {
	return &MockDefaultBufferedTransportInstanceRequirements_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockDefaultBufferedTransportInstanceRequirements_Connect_Call) Run(run func()) *MockDefaultBufferedTransportInstanceRequirements_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultBufferedTransportInstanceRequirements_Connect_Call) Return(_a0 error) *MockDefaultBufferedTransportInstanceRequirements_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultBufferedTransportInstanceRequirements_Connect_Call) RunAndReturn(run func() error) *MockDefaultBufferedTransportInstanceRequirements_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// GetReader provides a mock function with given fields:
func (_m *MockDefaultBufferedTransportInstanceRequirements) GetReader() *bufio.Reader {
	ret := _m.Called()

	var r0 *bufio.Reader
	if rf, ok := ret.Get(0).(func() *bufio.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bufio.Reader)
		}
	}

	return r0
}

// MockDefaultBufferedTransportInstanceRequirements_GetReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReader'
type MockDefaultBufferedTransportInstanceRequirements_GetReader_Call struct {
	*mock.Call
}

// GetReader is a helper method to define mock.On call
func (_e *MockDefaultBufferedTransportInstanceRequirements_Expecter) GetReader() *MockDefaultBufferedTransportInstanceRequirements_GetReader_Call {
	return &MockDefaultBufferedTransportInstanceRequirements_GetReader_Call{Call: _e.mock.On("GetReader")}
}

func (_c *MockDefaultBufferedTransportInstanceRequirements_GetReader_Call) Run(run func()) *MockDefaultBufferedTransportInstanceRequirements_GetReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultBufferedTransportInstanceRequirements_GetReader_Call) Return(_a0 *bufio.Reader) *MockDefaultBufferedTransportInstanceRequirements_GetReader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultBufferedTransportInstanceRequirements_GetReader_Call) RunAndReturn(run func() *bufio.Reader) *MockDefaultBufferedTransportInstanceRequirements_GetReader_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDefaultBufferedTransportInstanceRequirements creates a new instance of MockDefaultBufferedTransportInstanceRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDefaultBufferedTransportInstanceRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDefaultBufferedTransportInstanceRequirements {
	mock := &MockDefaultBufferedTransportInstanceRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
