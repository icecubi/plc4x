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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLogDataLogDataEntry is the corresponding interface of BACnetLogDataLogDataEntry
type BACnetLogDataLogDataEntry interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetLogDataLogDataEntry is the data-structure of this message
type _BACnetLogDataLogDataEntry struct {
	_BACnetLogDataLogDataEntryChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetLogDataLogDataEntryChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type BACnetLogDataLogDataEntryParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetLogDataLogDataEntry, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetLogDataLogDataEntryChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent BACnetLogDataLogDataEntry, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetLogDataLogDataEntry

	GetTypeName() string
	BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntry) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntry) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntry factory function for _BACnetLogDataLogDataEntry
func NewBACnetLogDataLogDataEntry(peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntry {
	return &_BACnetLogDataLogDataEntry{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntry(structType interface{}) BACnetLogDataLogDataEntry {
	if casted, ok := structType.(BACnetLogDataLogDataEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntry) GetTypeName() string {
	return "BACnetLogDataLogDataEntry"
}

func (m *_BACnetLogDataLogDataEntry) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataEntryParse(readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetLogDataLogDataEntryChildSerializeRequirement interface {
		BACnetLogDataLogDataEntry
		InitializeParent(BACnetLogDataLogDataEntry, BACnetTagHeader)
		GetParent() BACnetLogDataLogDataEntry
	}
	var _childTemp interface{}
	var _child BACnetLogDataLogDataEntryChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetLogDataLogDataEntryBooleanValue
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryBooleanValueParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	case peekedTagNumber == uint8(1): // BACnetLogDataLogDataEntryRealValue
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryRealValueParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	case peekedTagNumber == uint8(2): // BACnetLogDataLogDataEntryEnumeratedValue
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryEnumeratedValueParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	case peekedTagNumber == uint8(3): // BACnetLogDataLogDataEntryUnsignedValue
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryUnsignedValueParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	case peekedTagNumber == uint8(4): // BACnetLogDataLogDataEntryIntegerValue
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryIntegerValueParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	case peekedTagNumber == uint8(5): // BACnetLogDataLogDataEntryBitStringValue
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryBitStringValueParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	case peekedTagNumber == uint8(6): // BACnetLogDataLogDataEntryNullValue
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryNullValueParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	case peekedTagNumber == uint8(7): // BACnetLogDataLogDataEntryFailure
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryFailureParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	case peekedTagNumber == uint8(8): // BACnetLogDataLogDataEntryAnyValue
		_childTemp, typeSwitchError = BACnetLogDataLogDataEntryAnyValueParse(readBuffer)
		_child = _childTemp.(BACnetLogDataLogDataEntryChildSerializeRequirement)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntry")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetLogDataLogDataEntry) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetLogDataLogDataEntry, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntry")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntry")
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
