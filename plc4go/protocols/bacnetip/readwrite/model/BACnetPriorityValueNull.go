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

// BACnetPriorityValueNull is the data-structure of this message
type BACnetPriorityValueNull struct {
	*BACnetPriorityValue
	NullValue *BACnetApplicationTagNull

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

// IBACnetPriorityValueNull is the corresponding interface of BACnetPriorityValueNull
type IBACnetPriorityValueNull interface {
	IBACnetPriorityValue
	// GetNullValue returns NullValue (property field)
	GetNullValue() *BACnetApplicationTagNull
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPriorityValueNull) InitializeParent(parent *BACnetPriorityValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPriorityValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPriorityValueNull) GetParent() *BACnetPriorityValue {
	return m.BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPriorityValueNull) GetNullValue() *BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueNull factory function for BACnetPriorityValueNull
func NewBACnetPriorityValueNull(nullValue *BACnetApplicationTagNull, peekedTagHeader *BACnetTagHeader, objectTypeArgument BACnetObjectType) *BACnetPriorityValueNull {
	_result := &BACnetPriorityValueNull{
		NullValue:           nullValue,
		BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPriorityValueNull(structType interface{}) *BACnetPriorityValueNull {
	if casted, ok := structType.(BACnetPriorityValueNull); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPriorityValueNull); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPriorityValue); ok {
		return CastBACnetPriorityValueNull(casted.Child)
	}
	if casted, ok := structType.(*BACnetPriorityValue); ok {
		return CastBACnetPriorityValueNull(casted.Child)
	}
	return nil
}

func (m *BACnetPriorityValueNull) GetTypeName() string {
	return "BACnetPriorityValueNull"
}

func (m *BACnetPriorityValueNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPriorityValueNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPriorityValueNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueNullParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (*BACnetPriorityValueNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
	_nullValue, _nullValueErr := BACnetApplicationTagParse(readBuffer)
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field")
	}
	nullValue := CastBACnetApplicationTagNull(_nullValue)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueNull")
	}

	// Create a partially initialized instance
	_child := &BACnetPriorityValueNull{
		NullValue:           CastBACnetApplicationTagNull(nullValue),
		BACnetPriorityValue: &BACnetPriorityValue{},
	}
	_child.BACnetPriorityValue.Child = _child
	return _child, nil
}

func (m *BACnetPriorityValueNull) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueNull")
		}

		// Simple Field (nullValue)
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nullValue")
		}
		_nullValueErr := m.NullValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nullValue")
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueNull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPriorityValueNull) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}