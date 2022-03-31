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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CALReplyShort struct {
	*CALReply
}

// The corresponding interface
type ICALReplyShort interface {
	ICALReply
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

func (m *CALReplyShort) InitializeParent(parent *CALReply, calType byte, calData *CALData) {
	m.CALReply.CalType = calType
	m.CALReply.CalData = calData
}

func (m *CALReplyShort) GetParent() *CALReply {
	return m.CALReply
}

// NewCALReplyShort factory function for CALReplyShort
func NewCALReplyShort(calType byte, calData *CALData) *CALReplyShort {
	_result := &CALReplyShort{
		CALReply: NewCALReply(calType, calData),
	}
	_result.Child = _result
	return _result
}

func CastCALReplyShort(structType interface{}) *CALReplyShort {
	if casted, ok := structType.(CALReplyShort); ok {
		return &casted
	}
	if casted, ok := structType.(*CALReplyShort); ok {
		return casted
	}
	if casted, ok := structType.(CALReply); ok {
		return CastCALReplyShort(casted.Child)
	}
	if casted, ok := structType.(*CALReply); ok {
		return CastCALReplyShort(casted.Child)
	}
	return nil
}

func (m *CALReplyShort) GetTypeName() string {
	return "CALReplyShort"
}

func (m *CALReplyShort) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALReplyShort) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *CALReplyShort) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALReplyShortParse(readBuffer utils.ReadBuffer) (*CALReplyShort, error) {
	if pullErr := readBuffer.PullContext("CALReplyShort"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CALReplyShort"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CALReplyShort{
		CALReply: &CALReply{},
	}
	_child.CALReply.Child = _child
	return _child, nil
}

func (m *CALReplyShort) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyShort"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("CALReplyShort"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALReplyShort) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}