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

// BACnetEventTimestamps is the data-structure of this message
type BACnetEventTimestamps struct {
	OpeningTag  *BACnetOpeningTag
	ToOffnormal *BACnetTimeStamp
	ToFault     *BACnetTimeStamp
	ToNormal    *BACnetTimeStamp
	ClosingTag  *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetEventTimestamps is the corresponding interface of BACnetEventTimestamps
type IBACnetEventTimestamps interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetToOffnormal returns ToOffnormal (property field)
	GetToOffnormal() *BACnetTimeStamp
	// GetToFault returns ToFault (property field)
	GetToFault() *BACnetTimeStamp
	// GetToNormal returns ToNormal (property field)
	GetToNormal() *BACnetTimeStamp
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventTimestamps) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetEventTimestamps) GetToOffnormal() *BACnetTimeStamp {
	return m.ToOffnormal
}

func (m *BACnetEventTimestamps) GetToFault() *BACnetTimeStamp {
	return m.ToFault
}

func (m *BACnetEventTimestamps) GetToNormal() *BACnetTimeStamp {
	return m.ToNormal
}

func (m *BACnetEventTimestamps) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventTimestamps factory function for BACnetEventTimestamps
func NewBACnetEventTimestamps(openingTag *BACnetOpeningTag, toOffnormal *BACnetTimeStamp, toFault *BACnetTimeStamp, toNormal *BACnetTimeStamp, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetEventTimestamps {
	return &BACnetEventTimestamps{OpeningTag: openingTag, ToOffnormal: toOffnormal, ToFault: toFault, ToNormal: toNormal, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetEventTimestamps(structType interface{}) *BACnetEventTimestamps {
	if casted, ok := structType.(BACnetEventTimestamps); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventTimestamps); ok {
		return casted
	}
	return nil
}

func (m *BACnetEventTimestamps) GetTypeName() string {
	return "BACnetEventTimestamps"
}

func (m *BACnetEventTimestamps) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventTimestamps) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (toOffnormal)
	lengthInBits += m.ToOffnormal.GetLengthInBits()

	// Simple field (toFault)
	lengthInBits += m.ToFault.GetLengthInBits()

	// Simple field (toNormal)
	lengthInBits += m.ToNormal.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventTimestamps) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventTimestampsParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetEventTimestamps, error) {
	if pullErr := readBuffer.PullContext("BACnetEventTimestamps"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_OPENING_TAG))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toOffnormal)
	if pullErr := readBuffer.PullContext("toOffnormal"); pullErr != nil {
		return nil, pullErr
	}
	_toOffnormal, _toOffnormalErr := BACnetTimeStampParse(readBuffer)
	if _toOffnormalErr != nil {
		return nil, errors.Wrap(_toOffnormalErr, "Error parsing 'toOffnormal' field")
	}
	toOffnormal := CastBACnetTimeStamp(_toOffnormal)
	if closeErr := readBuffer.CloseContext("toOffnormal"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toFault)
	if pullErr := readBuffer.PullContext("toFault"); pullErr != nil {
		return nil, pullErr
	}
	_toFault, _toFaultErr := BACnetTimeStampParse(readBuffer)
	if _toFaultErr != nil {
		return nil, errors.Wrap(_toFaultErr, "Error parsing 'toFault' field")
	}
	toFault := CastBACnetTimeStamp(_toFault)
	if closeErr := readBuffer.CloseContext("toFault"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toNormal)
	if pullErr := readBuffer.PullContext("toNormal"); pullErr != nil {
		return nil, pullErr
	}
	_toNormal, _toNormalErr := BACnetTimeStampParse(readBuffer)
	if _toNormalErr != nil {
		return nil, errors.Wrap(_toNormalErr, "Error parsing 'toNormal' field")
	}
	toNormal := CastBACnetTimeStamp(_toNormal)
	if closeErr := readBuffer.CloseContext("toNormal"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_CLOSING_TAG))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventTimestamps"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetEventTimestamps(openingTag, toOffnormal, toFault, toNormal, closingTag, tagNumber), nil
}

func (m *BACnetEventTimestamps) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetEventTimestamps"); pushErr != nil {
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

	if popErr := writeBuffer.PopContext("BACnetEventTimestamps"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetEventTimestamps) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}