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
type BACnetPropertyStates struct {
	OpeningTag      *BACnetOpeningTag
	PeekedTagHeader *BACnetTagHeader
	ClosingTag      *BACnetClosingTag

	// Arguments.
	TagNumber uint8
	Child     IBACnetPropertyStatesChild
}

// The corresponding interface
type IBACnetPropertyStates interface {
	// PeekedTagNumber returns PeekedTagNumber
	PeekedTagNumber() uint8
	// GetOpeningTag returns OpeningTag
	GetOpeningTag() *BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader
	GetPeekedTagHeader() *BACnetTagHeader
	// GetClosingTag returns ClosingTag
	GetClosingTag() *BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetPropertyStatesParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetPropertyStates, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetPropertyStatesChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetPropertyStates, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag)
	GetTypeName() string
	IBACnetPropertyStates
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetPropertyStates) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetPropertyStates) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetPropertyStates) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetPropertyStates) GetPeekedTagNumber() uint8 {
	return m.GetPeekedTagHeader().GetActualTagNumber()
}

// NewBACnetPropertyStates factory function for BACnetPropertyStates
func NewBACnetPropertyStates(openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetPropertyStates {
	return &BACnetPropertyStates{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetPropertyStates(structType interface{}) *BACnetPropertyStates {
	castFunc := func(typ interface{}) *BACnetPropertyStates {
		if casted, ok := typ.(BACnetPropertyStates); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetPropertyStates); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetPropertyStates) GetTypeName() string {
	return "BACnetPropertyStates"
}

func (m *BACnetPropertyStates) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStates) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetPropertyStates) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStates) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetPropertyStates, error) {
	if pullErr := readBuffer.PullContext("BACnetPropertyStates"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType_OPENING_TAG)
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Peek Field (peekedTagHeader)
	currentPos := readBuffer.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetPropertyStates
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetPropertyStatesBoolean
		_parent, typeSwitchError = BACnetPropertyStatesBooleanParse(readBuffer, tagNumber, peekedTagNumber)
	case peekedTagNumber == uint8(1): // BACnetPropertyStatesBinaryValue
		_parent, typeSwitchError = BACnetPropertyStatesBinaryValueParse(readBuffer, tagNumber, peekedTagNumber)
	case peekedTagNumber == uint8(16): // BACnetPropertyStatesAction
		_parent, typeSwitchError = BACnetPropertyStatesActionParse(readBuffer, tagNumber, peekedTagNumber)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType_CLOSING_TAG)
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStates"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, openingTag, peekedTagHeader, closingTag)
	return _parent, nil
}

func (m *BACnetPropertyStates) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetPropertyStates) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetPropertyStates, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetPropertyStates"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyStates"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetPropertyStates) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
