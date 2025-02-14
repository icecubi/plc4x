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

package options

import mock "github.com/stretchr/testify/mock"

// MockDiscoveryOptionLocalAddress is an autogenerated mock type for the DiscoveryOptionLocalAddress type
type MockDiscoveryOptionLocalAddress struct {
	mock.Mock
}

type MockDiscoveryOptionLocalAddress_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDiscoveryOptionLocalAddress) EXPECT() *MockDiscoveryOptionLocalAddress_Expecter {
	return &MockDiscoveryOptionLocalAddress_Expecter{mock: &_m.Mock}
}

// GetLocalAddress provides a mock function with given fields:
func (_m *MockDiscoveryOptionLocalAddress) GetLocalAddress() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDiscoveryOptionLocalAddress_GetLocalAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLocalAddress'
type MockDiscoveryOptionLocalAddress_GetLocalAddress_Call struct {
	*mock.Call
}

// GetLocalAddress is a helper method to define mock.On call
func (_e *MockDiscoveryOptionLocalAddress_Expecter) GetLocalAddress() *MockDiscoveryOptionLocalAddress_GetLocalAddress_Call {
	return &MockDiscoveryOptionLocalAddress_GetLocalAddress_Call{Call: _e.mock.On("GetLocalAddress")}
}

func (_c *MockDiscoveryOptionLocalAddress_GetLocalAddress_Call) Run(run func()) *MockDiscoveryOptionLocalAddress_GetLocalAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDiscoveryOptionLocalAddress_GetLocalAddress_Call) Return(_a0 string) *MockDiscoveryOptionLocalAddress_GetLocalAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscoveryOptionLocalAddress_GetLocalAddress_Call) RunAndReturn(run func() string) *MockDiscoveryOptionLocalAddress_GetLocalAddress_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDiscoveryOptionLocalAddress creates a new instance of MockDiscoveryOptionLocalAddress. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDiscoveryOptionLocalAddress(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDiscoveryOptionLocalAddress {
	mock := &MockDiscoveryOptionLocalAddress{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
