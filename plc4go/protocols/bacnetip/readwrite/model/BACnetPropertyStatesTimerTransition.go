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

// BACnetPropertyStatesTimerTransition is the data-structure of this message
type BACnetPropertyStatesTimerTransition struct {
	*BACnetPropertyStates
	TimerTransition *BACnetTimerTransitionTagged
}

// IBACnetPropertyStatesTimerTransition is the corresponding interface of BACnetPropertyStatesTimerTransition
type IBACnetPropertyStatesTimerTransition interface {
	IBACnetPropertyStates
	// GetTimerTransition returns TimerTransition (property field)
	GetTimerTransition() *BACnetTimerTransitionTagged
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

func (m *BACnetPropertyStatesTimerTransition) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesTimerTransition) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesTimerTransition) GetTimerTransition() *BACnetTimerTransitionTagged {
	return m.TimerTransition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesTimerTransition factory function for BACnetPropertyStatesTimerTransition
func NewBACnetPropertyStatesTimerTransition(timerTransition *BACnetTimerTransitionTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesTimerTransition {
	_result := &BACnetPropertyStatesTimerTransition{
		TimerTransition:      timerTransition,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesTimerTransition(structType interface{}) *BACnetPropertyStatesTimerTransition {
	if casted, ok := structType.(BACnetPropertyStatesTimerTransition); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesTimerTransition); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesTimerTransition(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesTimerTransition(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesTimerTransition) GetTypeName() string {
	return "BACnetPropertyStatesTimerTransition"
}

func (m *BACnetPropertyStatesTimerTransition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesTimerTransition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timerTransition)
	lengthInBits += m.TimerTransition.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesTimerTransition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesTimerTransitionParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesTimerTransition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesTimerTransition"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timerTransition)
	if pullErr := readBuffer.PullContext("timerTransition"); pullErr != nil {
		return nil, pullErr
	}
	_timerTransition, _timerTransitionErr := BACnetTimerTransitionTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _timerTransitionErr != nil {
		return nil, errors.Wrap(_timerTransitionErr, "Error parsing 'timerTransition' field")
	}
	timerTransition := CastBACnetTimerTransitionTagged(_timerTransition)
	if closeErr := readBuffer.CloseContext("timerTransition"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesTimerTransition"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesTimerTransition{
		TimerTransition:      CastBACnetTimerTransitionTagged(timerTransition),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesTimerTransition) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesTimerTransition"); pushErr != nil {
			return pushErr
		}

		// Simple Field (timerTransition)
		if pushErr := writeBuffer.PushContext("timerTransition"); pushErr != nil {
			return pushErr
		}
		_timerTransitionErr := m.TimerTransition.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timerTransition"); popErr != nil {
			return popErr
		}
		if _timerTransitionErr != nil {
			return errors.Wrap(_timerTransitionErr, "Error serializing 'timerTransition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesTimerTransition"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesTimerTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}