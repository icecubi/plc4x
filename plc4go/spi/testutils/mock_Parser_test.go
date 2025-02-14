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

package testutils

import (
	utils "github.com/apache/plc4x/plc4go/spi/utils"
	mock "github.com/stretchr/testify/mock"
)

// MockParser is an autogenerated mock type for the Parser type
type MockParser struct {
	mock.Mock
}

type MockParser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockParser) EXPECT() *MockParser_Expecter {
	return &MockParser_Expecter{mock: &_m.Mock}
}

// Parse provides a mock function with given fields: typeName, arguments, io
func (_m *MockParser) Parse(typeName string, arguments []string, io utils.ReadBuffer) (interface{}, error) {
	ret := _m.Called(typeName, arguments, io)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string, utils.ReadBuffer) (interface{}, error)); ok {
		return rf(typeName, arguments, io)
	}
	if rf, ok := ret.Get(0).(func(string, []string, utils.ReadBuffer) interface{}); ok {
		r0 = rf(typeName, arguments, io)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string, utils.ReadBuffer) error); ok {
		r1 = rf(typeName, arguments, io)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockParser_Parse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Parse'
type MockParser_Parse_Call struct {
	*mock.Call
}

// Parse is a helper method to define mock.On call
//   - typeName string
//   - arguments []string
//   - io utils.ReadBuffer
func (_e *MockParser_Expecter) Parse(typeName interface{}, arguments interface{}, io interface{}) *MockParser_Parse_Call {
	return &MockParser_Parse_Call{Call: _e.mock.On("Parse", typeName, arguments, io)}
}

func (_c *MockParser_Parse_Call) Run(run func(typeName string, arguments []string, io utils.ReadBuffer)) *MockParser_Parse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string), args[2].(utils.ReadBuffer))
	})
	return _c
}

func (_c *MockParser_Parse_Call) Return(_a0 interface{}, _a1 error) *MockParser_Parse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockParser_Parse_Call) RunAndReturn(run func(string, []string, utils.ReadBuffer) (interface{}, error)) *MockParser_Parse_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockParser creates a new instance of MockParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockParser(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockParser {
	mock := &MockParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
