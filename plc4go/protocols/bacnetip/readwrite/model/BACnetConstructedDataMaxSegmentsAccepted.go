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

// BACnetConstructedDataMaxSegmentsAccepted is the data-structure of this message
type BACnetConstructedDataMaxSegmentsAccepted struct {
	*BACnetConstructedData
	MaxSegmentsAccepted *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMaxSegmentsAccepted is the corresponding interface of BACnetConstructedDataMaxSegmentsAccepted
type IBACnetConstructedDataMaxSegmentsAccepted interface {
	IBACnetConstructedData
	// GetMaxSegmentsAccepted returns MaxSegmentsAccepted (property field)
	GetMaxSegmentsAccepted() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataMaxSegmentsAccepted) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMaxSegmentsAccepted) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_SEGMENTS_ACCEPTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMaxSegmentsAccepted) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMaxSegmentsAccepted) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMaxSegmentsAccepted) GetMaxSegmentsAccepted() *BACnetApplicationTagUnsignedInteger {
	return m.MaxSegmentsAccepted
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaxSegmentsAccepted factory function for BACnetConstructedDataMaxSegmentsAccepted
func NewBACnetConstructedDataMaxSegmentsAccepted(maxSegmentsAccepted *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMaxSegmentsAccepted {
	_result := &BACnetConstructedDataMaxSegmentsAccepted{
		MaxSegmentsAccepted:   maxSegmentsAccepted,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMaxSegmentsAccepted(structType interface{}) *BACnetConstructedDataMaxSegmentsAccepted {
	if casted, ok := structType.(BACnetConstructedDataMaxSegmentsAccepted); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxSegmentsAccepted); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaxSegmentsAccepted(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaxSegmentsAccepted(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMaxSegmentsAccepted) GetTypeName() string {
	return "BACnetConstructedDataMaxSegmentsAccepted"
}

func (m *BACnetConstructedDataMaxSegmentsAccepted) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMaxSegmentsAccepted) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxSegmentsAccepted)
	lengthInBits += m.MaxSegmentsAccepted.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMaxSegmentsAccepted) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaxSegmentsAcceptedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMaxSegmentsAccepted, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxSegmentsAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxSegmentsAccepted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxSegmentsAccepted)
	if pullErr := readBuffer.PullContext("maxSegmentsAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxSegmentsAccepted")
	}
	_maxSegmentsAccepted, _maxSegmentsAcceptedErr := BACnetApplicationTagParse(readBuffer)
	if _maxSegmentsAcceptedErr != nil {
		return nil, errors.Wrap(_maxSegmentsAcceptedErr, "Error parsing 'maxSegmentsAccepted' field")
	}
	maxSegmentsAccepted := CastBACnetApplicationTagUnsignedInteger(_maxSegmentsAccepted)
	if closeErr := readBuffer.CloseContext("maxSegmentsAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxSegmentsAccepted")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxSegmentsAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxSegmentsAccepted")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMaxSegmentsAccepted{
		MaxSegmentsAccepted:   CastBACnetApplicationTagUnsignedInteger(maxSegmentsAccepted),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMaxSegmentsAccepted) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxSegmentsAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxSegmentsAccepted")
		}

		// Simple Field (maxSegmentsAccepted)
		if pushErr := writeBuffer.PushContext("maxSegmentsAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxSegmentsAccepted")
		}
		_maxSegmentsAcceptedErr := m.MaxSegmentsAccepted.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("maxSegmentsAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxSegmentsAccepted")
		}
		if _maxSegmentsAcceptedErr != nil {
			return errors.Wrap(_maxSegmentsAcceptedErr, "Error serializing 'maxSegmentsAccepted' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxSegmentsAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxSegmentsAccepted")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMaxSegmentsAccepted) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}