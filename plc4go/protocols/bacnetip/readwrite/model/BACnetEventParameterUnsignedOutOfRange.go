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

// BACnetEventParameterUnsignedOutOfRange is the data-structure of this message
type BACnetEventParameterUnsignedOutOfRange struct {
	*BACnetEventParameter
	OpeningTag *BACnetOpeningTag
	TimeDelay  *BACnetContextTagUnsignedInteger
	LowLimit   *BACnetContextTagUnsignedInteger
	HighLimit  *BACnetContextTagUnsignedInteger
	Deadband   *BACnetContextTagUnsignedInteger
	ClosingTag *BACnetClosingTag
}

// IBACnetEventParameterUnsignedOutOfRange is the corresponding interface of BACnetEventParameterUnsignedOutOfRange
type IBACnetEventParameterUnsignedOutOfRange interface {
	IBACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() *BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() *BACnetContextTagUnsignedInteger
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() *BACnetContextTagUnsignedInteger
	// GetDeadband returns Deadband (property field)
	GetDeadband() *BACnetContextTagUnsignedInteger
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
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetEventParameterUnsignedOutOfRange) InitializeParent(parent *BACnetEventParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetEventParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetParent() *BACnetEventParameter {
	return m.BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventParameterUnsignedOutOfRange) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetTimeDelay() *BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetLowLimit() *BACnetContextTagUnsignedInteger {
	return m.LowLimit
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetHighLimit() *BACnetContextTagUnsignedInteger {
	return m.HighLimit
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetDeadband() *BACnetContextTagUnsignedInteger {
	return m.Deadband
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterUnsignedOutOfRange factory function for BACnetEventParameterUnsignedOutOfRange
func NewBACnetEventParameterUnsignedOutOfRange(openingTag *BACnetOpeningTag, timeDelay *BACnetContextTagUnsignedInteger, lowLimit *BACnetContextTagUnsignedInteger, highLimit *BACnetContextTagUnsignedInteger, deadband *BACnetContextTagUnsignedInteger, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetEventParameterUnsignedOutOfRange {
	_result := &BACnetEventParameterUnsignedOutOfRange{
		OpeningTag:           openingTag,
		TimeDelay:            timeDelay,
		LowLimit:             lowLimit,
		HighLimit:            highLimit,
		Deadband:             deadband,
		ClosingTag:           closingTag,
		BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventParameterUnsignedOutOfRange(structType interface{}) *BACnetEventParameterUnsignedOutOfRange {
	if casted, ok := structType.(BACnetEventParameterUnsignedOutOfRange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventParameterUnsignedOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventParameter); ok {
		return CastBACnetEventParameterUnsignedOutOfRange(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return CastBACnetEventParameterUnsignedOutOfRange(casted.Child)
	}
	return nil
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetTypeName() string {
	return "BACnetEventParameterUnsignedOutOfRange"
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits()

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits()

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventParameterUnsignedOutOfRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterUnsignedOutOfRangeParse(readBuffer utils.ReadBuffer) (*BACnetEventParameterUnsignedOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterUnsignedOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterUnsignedOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(16)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field")
	}
	timeDelay := CastBACnetContextTagUnsignedInteger(_timeDelay)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (lowLimit)
	if pullErr := readBuffer.PullContext("lowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowLimit")
	}
	_lowLimit, _lowLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _lowLimitErr != nil {
		return nil, errors.Wrap(_lowLimitErr, "Error parsing 'lowLimit' field")
	}
	lowLimit := CastBACnetContextTagUnsignedInteger(_lowLimit)
	if closeErr := readBuffer.CloseContext("lowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowLimit")
	}

	// Simple Field (highLimit)
	if pullErr := readBuffer.PullContext("highLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for highLimit")
	}
	_highLimit, _highLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _highLimitErr != nil {
		return nil, errors.Wrap(_highLimitErr, "Error parsing 'highLimit' field")
	}
	highLimit := CastBACnetContextTagUnsignedInteger(_highLimit)
	if closeErr := readBuffer.CloseContext("highLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for highLimit")
	}

	// Simple Field (deadband)
	if pullErr := readBuffer.PullContext("deadband"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deadband")
	}
	_deadband, _deadbandErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _deadbandErr != nil {
		return nil, errors.Wrap(_deadbandErr, "Error parsing 'deadband' field")
	}
	deadband := CastBACnetContextTagUnsignedInteger(_deadband)
	if closeErr := readBuffer.CloseContext("deadband"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deadband")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(16)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterUnsignedOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterUnsignedOutOfRange")
	}

	// Create a partially initialized instance
	_child := &BACnetEventParameterUnsignedOutOfRange{
		OpeningTag:           CastBACnetOpeningTag(openingTag),
		TimeDelay:            CastBACnetContextTagUnsignedInteger(timeDelay),
		LowLimit:             CastBACnetContextTagUnsignedInteger(lowLimit),
		HighLimit:            CastBACnetContextTagUnsignedInteger(highLimit),
		Deadband:             CastBACnetContextTagUnsignedInteger(deadband),
		ClosingTag:           CastBACnetClosingTag(closingTag),
		BACnetEventParameter: &BACnetEventParameter{},
	}
	_child.BACnetEventParameter.Child = _child
	return _child, nil
}

func (m *BACnetEventParameterUnsignedOutOfRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterUnsignedOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterUnsignedOutOfRange")
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for openingTag")
		}
		_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for openingTag")
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeDelay")
		}
		_timeDelayErr := m.TimeDelay.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeDelay")
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (lowLimit)
		if pushErr := writeBuffer.PushContext("lowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lowLimit")
		}
		_lowLimitErr := m.LowLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lowLimit")
		}
		if _lowLimitErr != nil {
			return errors.Wrap(_lowLimitErr, "Error serializing 'lowLimit' field")
		}

		// Simple Field (highLimit)
		if pushErr := writeBuffer.PushContext("highLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for highLimit")
		}
		_highLimitErr := m.HighLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("highLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for highLimit")
		}
		if _highLimitErr != nil {
			return errors.Wrap(_highLimitErr, "Error serializing 'highLimit' field")
		}

		// Simple Field (deadband)
		if pushErr := writeBuffer.PushContext("deadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deadband")
		}
		_deadbandErr := m.Deadband.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deadband")
		}
		if _deadbandErr != nil {
			return errors.Wrap(_deadbandErr, "Error serializing 'deadband' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for closingTag")
		}
		_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for closingTag")
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterUnsignedOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterUnsignedOutOfRange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventParameterUnsignedOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}