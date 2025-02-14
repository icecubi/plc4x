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

package model

import (
	values "github.com/apache/plc4x/plc4go/pkg/api/values"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcTag is an autogenerated mock type for the PlcTag type
type MockPlcTag struct {
	mock.Mock
}

type MockPlcTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcTag) EXPECT() *MockPlcTag_Expecter {
	return &MockPlcTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockPlcTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockPlcTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetAddressString() *MockPlcTag_GetAddressString_Call {
	return &MockPlcTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockPlcTag_GetAddressString_Call) Run(run func()) *MockPlcTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetAddressString_Call) Return(_a0 string) *MockPlcTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetAddressString_Call) RunAndReturn(run func() string) *MockPlcTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockPlcTag) GetArrayInfo() []ArrayInfo {
	ret := _m.Called()

	var r0 []ArrayInfo
	if rf, ok := ret.Get(0).(func() []ArrayInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ArrayInfo)
		}
	}

	return r0
}

// MockPlcTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockPlcTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetArrayInfo() *MockPlcTag_GetArrayInfo_Call {
	return &MockPlcTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockPlcTag_GetArrayInfo_Call) Run(run func()) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetArrayInfo_Call) Return(_a0 []ArrayInfo) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetArrayInfo_Call) RunAndReturn(run func() []ArrayInfo) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockPlcTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockPlcTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockPlcTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetValueType() *MockPlcTag_GetValueType_Call {
	return &MockPlcTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockPlcTag_GetValueType_Call) Run(run func()) *MockPlcTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockPlcTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockPlcTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcTag creates a new instance of MockPlcTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcTag(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcTag {
	mock := &MockPlcTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
