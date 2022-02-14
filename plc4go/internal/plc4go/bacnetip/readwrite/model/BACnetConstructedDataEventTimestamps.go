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
type BACnetConstructedDataEventTimestamps struct {
	*BACnetConstructedData
	ToOffnormal *BACnetContextTagTime
	ToFault     *BACnetContextTagUnsignedInteger
	ToNormal    *BACnetDateTime

	// Arguments.
	TagNumber                  uint8
	PropertyIdentifierArgument BACnetContextTagPropertyIdentifier
}

// The corresponding interface
type IBACnetConstructedDataEventTimestamps interface {
	// GetToOffnormal returns ToOffnormal
	GetToOffnormal() *BACnetContextTagTime
	// GetToFault returns ToFault
	GetToFault() *BACnetContextTagUnsignedInteger
	// GetToNormal returns ToNormal
	GetToNormal() *BACnetDateTime
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConstructedDataEventTimestamps) ObjectType() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEventTimestamps) GetObjectType() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEventTimestamps) PropertyIdentifierEnum() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_TIME_STAMPS
}

func (m *BACnetConstructedDataEventTimestamps) GetPropertyIdentifierEnum() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_TIME_STAMPS
}

func (m *BACnetConstructedDataEventTimestamps) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetConstructedDataEventTimestamps) GetToOffnormal() *BACnetContextTagTime {
	return m.ToOffnormal
}

func (m *BACnetConstructedDataEventTimestamps) GetToFault() *BACnetContextTagUnsignedInteger {
	return m.ToFault
}

func (m *BACnetConstructedDataEventTimestamps) GetToNormal() *BACnetDateTime {
	return m.ToNormal
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventTimestamps factory function for BACnetConstructedDataEventTimestamps
func NewBACnetConstructedDataEventTimestamps(toOffnormal *BACnetContextTagTime, toFault *BACnetContextTagUnsignedInteger, toNormal *BACnetDateTime, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8, propertyIdentifierArgument BACnetContextTagPropertyIdentifier) *BACnetConstructedData {
	child := &BACnetConstructedDataEventTimestamps{
		ToOffnormal:           toOffnormal,
		ToFault:               toFault,
		ToNormal:              toNormal,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber, propertyIdentifierArgument),
	}
	child.Child = child
	return child.BACnetConstructedData
}

func CastBACnetConstructedDataEventTimestamps(structType interface{}) *BACnetConstructedDataEventTimestamps {
	castFunc := func(typ interface{}) *BACnetConstructedDataEventTimestamps {
		if casted, ok := typ.(BACnetConstructedDataEventTimestamps); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConstructedDataEventTimestamps); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConstructedData); ok {
			return CastBACnetConstructedDataEventTimestamps(casted.Child)
		}
		if casted, ok := typ.(*BACnetConstructedData); ok {
			return CastBACnetConstructedDataEventTimestamps(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConstructedDataEventTimestamps) GetTypeName() string {
	return "BACnetConstructedDataEventTimestamps"
}

func (m *BACnetConstructedDataEventTimestamps) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEventTimestamps) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (toOffnormal)
	lengthInBits += m.ToOffnormal.GetLengthInBits()

	// Simple field (toFault)
	lengthInBits += m.ToFault.GetLengthInBits()

	// Simple field (toNormal)
	lengthInBits += m.ToNormal.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataEventTimestamps) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventTimestampsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument *BACnetContextTagPropertyIdentifier) (*BACnetConstructedData, error) {
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventTimestamps"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (toOffnormal)
	if pullErr := readBuffer.PullContext("toOffnormal"); pullErr != nil {
		return nil, pullErr
	}
	_toOffnormal, _toOffnormalErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType_TIME)
	if _toOffnormalErr != nil {
		return nil, errors.Wrap(_toOffnormalErr, "Error parsing 'toOffnormal' field")
	}
	toOffnormal := CastBACnetContextTagTime(_toOffnormal)
	if closeErr := readBuffer.CloseContext("toOffnormal"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toFault)
	if pullErr := readBuffer.PullContext("toFault"); pullErr != nil {
		return nil, pullErr
	}
	_toFault, _toFaultErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType_UNSIGNED_INTEGER)
	if _toFaultErr != nil {
		return nil, errors.Wrap(_toFaultErr, "Error parsing 'toFault' field")
	}
	toFault := CastBACnetContextTagUnsignedInteger(_toFault)
	if closeErr := readBuffer.CloseContext("toFault"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toNormal)
	if pullErr := readBuffer.PullContext("toNormal"); pullErr != nil {
		return nil, pullErr
	}
	_toNormal, _toNormalErr := BACnetDateTimeParse(readBuffer, uint8(uint8(2)))
	if _toNormalErr != nil {
		return nil, errors.Wrap(_toNormalErr, "Error parsing 'toNormal' field")
	}
	toNormal := CastBACnetDateTime(_toNormal)
	if closeErr := readBuffer.CloseContext("toNormal"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventTimestamps"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEventTimestamps{
		ToOffnormal:           CastBACnetContextTagTime(toOffnormal),
		ToFault:               CastBACnetContextTagUnsignedInteger(toFault),
		ToNormal:              CastBACnetDateTime(toNormal),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child.BACnetConstructedData, nil
}

func (m *BACnetConstructedDataEventTimestamps) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventTimestamps"); pushErr != nil {
			return pushErr
		}

		// Simple Field (toOffnormal)
		if pushErr := writeBuffer.PushContext("toOffnormal"); pushErr != nil {
			return pushErr
		}
		_toOffnormalErr := m.ToOffnormal.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("toOffnormal"); popErr != nil {
			return popErr
		}
		if _toOffnormalErr != nil {
			return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
		}

		// Simple Field (toFault)
		if pushErr := writeBuffer.PushContext("toFault"); pushErr != nil {
			return pushErr
		}
		_toFaultErr := m.ToFault.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("toFault"); popErr != nil {
			return popErr
		}
		if _toFaultErr != nil {
			return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
		}

		// Simple Field (toNormal)
		if pushErr := writeBuffer.PushContext("toNormal"); pushErr != nil {
			return pushErr
		}
		_toNormalErr := m.ToNormal.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("toNormal"); popErr != nil {
			return popErr
		}
		if _toNormalErr != nil {
			return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventTimestamps"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEventTimestamps) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
