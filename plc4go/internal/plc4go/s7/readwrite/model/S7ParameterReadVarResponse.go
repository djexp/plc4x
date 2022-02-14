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
type S7ParameterReadVarResponse struct {
	*S7Parameter
	NumItems uint8
}

// The corresponding interface
type IS7ParameterReadVarResponse interface {
	// GetNumItems returns NumItems
	GetNumItems() uint8
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
func (m *S7ParameterReadVarResponse) ParameterType() uint8 {
	return 0x04
}

func (m *S7ParameterReadVarResponse) GetParameterType() uint8 {
	return 0x04
}

func (m *S7ParameterReadVarResponse) MessageType() uint8 {
	return 0x03
}

func (m *S7ParameterReadVarResponse) GetMessageType() uint8 {
	return 0x03
}

func (m *S7ParameterReadVarResponse) InitializeParent(parent *S7Parameter) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *S7ParameterReadVarResponse) GetNumItems() uint8 {
	return m.NumItems
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewS7ParameterReadVarResponse factory function for S7ParameterReadVarResponse
func NewS7ParameterReadVarResponse(numItems uint8) *S7Parameter {
	child := &S7ParameterReadVarResponse{
		NumItems:    numItems,
		S7Parameter: NewS7Parameter(),
	}
	child.Child = child
	return child.S7Parameter
}

func CastS7ParameterReadVarResponse(structType interface{}) *S7ParameterReadVarResponse {
	castFunc := func(typ interface{}) *S7ParameterReadVarResponse {
		if casted, ok := typ.(S7ParameterReadVarResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*S7ParameterReadVarResponse); ok {
			return casted
		}
		if casted, ok := typ.(S7Parameter); ok {
			return CastS7ParameterReadVarResponse(casted.Child)
		}
		if casted, ok := typ.(*S7Parameter); ok {
			return CastS7ParameterReadVarResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7ParameterReadVarResponse) GetTypeName() string {
	return "S7ParameterReadVarResponse"
}

func (m *S7ParameterReadVarResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7ParameterReadVarResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numItems)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7ParameterReadVarResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterReadVarResponseParse(readBuffer utils.ReadBuffer, messageType uint8) (*S7Parameter, error) {
	if pullErr := readBuffer.PullContext("S7ParameterReadVarResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (numItems)
	_numItems, _numItemsErr := readBuffer.ReadUint8("numItems", 8)
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field")
	}
	numItems := _numItems

	if closeErr := readBuffer.CloseContext("S7ParameterReadVarResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7ParameterReadVarResponse{
		NumItems:    numItems,
		S7Parameter: &S7Parameter{},
	}
	_child.S7Parameter.Child = _child
	return _child.S7Parameter, nil
}

func (m *S7ParameterReadVarResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterReadVarResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (numItems)
		numItems := uint8(m.NumItems)
		_numItemsErr := writeBuffer.WriteUint8("numItems", 8, (numItems))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterReadVarResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7ParameterReadVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
