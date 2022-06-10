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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAPDULength is the data-structure of this message
type BACnetConstructedDataAPDULength struct {
	*BACnetConstructedData
	ApduLength *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAPDULength is the corresponding interface of BACnetConstructedDataAPDULength
type IBACnetConstructedDataAPDULength interface {
	IBACnetConstructedData
	// GetApduLength returns ApduLength (property field)
	GetApduLength() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataAPDULength) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAPDULength) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APDU_LENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAPDULength) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAPDULength) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAPDULength) GetApduLength() *BACnetApplicationTagUnsignedInteger {
	return m.ApduLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAPDULength factory function for BACnetConstructedDataAPDULength
func NewBACnetConstructedDataAPDULength(apduLength *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAPDULength {
	_result := &BACnetConstructedDataAPDULength{
		ApduLength:            apduLength,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAPDULength(structType interface{}) *BACnetConstructedDataAPDULength {
	if casted, ok := structType.(BACnetConstructedDataAPDULength); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAPDULength); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAPDULength(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAPDULength(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAPDULength) GetTypeName() string {
	return "BACnetConstructedDataAPDULength"
}

func (m *BACnetConstructedDataAPDULength) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAPDULength) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (apduLength)
	lengthInBits += m.ApduLength.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAPDULength) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAPDULengthParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAPDULength, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAPDULength"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAPDULength")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (apduLength)
	if pullErr := readBuffer.PullContext("apduLength"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for apduLength")
	}
	_apduLength, _apduLengthErr := BACnetApplicationTagParse(readBuffer)
	if _apduLengthErr != nil {
		return nil, errors.Wrap(_apduLengthErr, "Error parsing 'apduLength' field")
	}
	apduLength := CastBACnetApplicationTagUnsignedInteger(_apduLength)
	if closeErr := readBuffer.CloseContext("apduLength"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for apduLength")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAPDULength"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAPDULength")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAPDULength{
		ApduLength:            CastBACnetApplicationTagUnsignedInteger(apduLength),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAPDULength) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAPDULength"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAPDULength")
		}

		// Simple Field (apduLength)
		if pushErr := writeBuffer.PushContext("apduLength"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for apduLength")
		}
		_apduLengthErr := m.ApduLength.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("apduLength"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for apduLength")
		}
		if _apduLengthErr != nil {
			return errors.Wrap(_apduLengthErr, "Error serializing 'apduLength' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAPDULength"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAPDULength")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAPDULength) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}