/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetContextTagBoolean struct {
	*BACnetContextTag
	Value   uint8
	IsTrue  bool
	IsFalse bool
}

// The corresponding interface
type IBACnetContextTagBoolean interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagBoolean) DataType() BACnetDataType {
	return BACnetDataType_BOOLEAN
}

func (m *BACnetContextTagBoolean) LengthValueType() uint8 {
	return 0
}

func (m *BACnetContextTagBoolean) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader, tagNumber uint8, actualLength uint32, lengthValueType uint8) {
	m.Header = header
}

func NewBACnetContextTagBoolean(value uint8, isTrue bool, isFalse bool, header *BACnetTagHeader, tagNumber uint8, actualLength uint32, lengthValueType uint8) *BACnetContextTag {
	child := &BACnetContextTagBoolean{
		Value:            value,
		IsTrue:           isTrue,
		IsFalse:          isFalse,
		BACnetContextTag: NewBACnetContextTag(header, tagNumber, actualLength, lengthValueType),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagBoolean(structType interface{}) *BACnetContextTagBoolean {
	castFunc := func(typ interface{}) *BACnetContextTagBoolean {
		if casted, ok := typ.(BACnetContextTagBoolean); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagBoolean); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagBoolean(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagBoolean(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagBoolean) GetTypeName() string {
	return "BACnetContextTagBoolean"
}

func (m *BACnetContextTagBoolean) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTagBoolean) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagBoolean) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagBooleanParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagBoolean"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadUint8("value", 8)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	// Virtual field
	_isTrue := bool((value) == (1))
	isTrue := bool(_isTrue)

	// Virtual field
	_isFalse := bool((value) == (0))
	isFalse := bool(_isFalse)

	if closeErr := readBuffer.CloseContext("BACnetContextTagBoolean"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagBoolean{
		Value:            value,
		IsTrue:           isTrue,
		IsFalse:          isFalse,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagBoolean"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		value := uint8(m.Value)
		_valueErr := writeBuffer.WriteUint8("value", 8, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
		// Virtual field
		if _isTrueErr := writeBuffer.WriteVirtual("isTrue", m.IsTrue); _isTrueErr != nil {
			return errors.Wrap(_isTrueErr, "Error serializing 'isTrue' field")
		}
		// Virtual field
		if _isFalseErr := writeBuffer.WriteVirtual("isFalse", m.IsFalse); _isFalseErr != nil {
			return errors.Wrap(_isFalseErr, "Error serializing 'isFalse' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagBoolean"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
