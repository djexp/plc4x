/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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

// CALDataNormalReplyReply is the corresponding interface of CALDataNormalReplyReply
type CALDataNormalReplyReply interface {
	utils.LengthAware
	utils.Serializable
	CALDataNormal
	// GetParamNumber returns ParamNumber (property field)
	GetParamNumber() uint8
	// GetData returns Data (property field)
	GetData() []byte
}

// CALDataNormalReplyReplyExactly can be used when we want exactly this type and not a type which fulfills CALDataNormalReplyReply.
// This is useful for switch cases.
type CALDataNormalReplyReplyExactly interface {
	CALDataNormalReplyReply
	isCALDataNormalReplyReply() bool
}

// _CALDataNormalReplyReply is the data-structure of this message
type _CALDataNormalReplyReply struct {
	*_CALDataNormal
	ParamNumber uint8
	Data        []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataNormalReplyReply) InitializeParent(parent CALDataNormal, commandTypeContainer CALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_CALDataNormalReplyReply) GetParent() CALDataNormal {
	return m._CALDataNormal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataNormalReplyReply) GetParamNumber() uint8 {
	return m.ParamNumber
}

func (m *_CALDataNormalReplyReply) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataNormalReplyReply factory function for _CALDataNormalReplyReply
func NewCALDataNormalReplyReply(paramNumber uint8, data []byte, commandTypeContainer CALCommandTypeContainer) *_CALDataNormalReplyReply {
	_result := &_CALDataNormalReplyReply{
		ParamNumber:    paramNumber,
		Data:           data,
		_CALDataNormal: NewCALDataNormal(commandTypeContainer),
	}
	_result._CALDataNormal._CALDataNormalChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataNormalReplyReply(structType interface{}) CALDataNormalReplyReply {
	if casted, ok := structType.(CALDataNormalReplyReply); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataNormalReplyReply); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataNormalReplyReply) GetTypeName() string {
	return "CALDataNormalReplyReply"
}

func (m *_CALDataNormalReplyReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataNormalReplyReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNumber)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CALDataNormalReplyReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataNormalReplyReplyParse(readBuffer utils.ReadBuffer, commandTypeContainer CALCommandTypeContainer) (CALDataNormalReplyReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataNormalReplyReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataNormalReplyReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNumber)
	_paramNumber, _paramNumberErr := readBuffer.ReadUint8("paramNumber", 8)
	if _paramNumberErr != nil {
		return nil, errors.Wrap(_paramNumberErr, "Error parsing 'paramNumber' field of CALDataNormalReplyReply")
	}
	paramNumber := _paramNumber
	// Byte Array field (data)
	numberOfBytesdata := int(commandTypeContainer.NumBytes())
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of CALDataNormalReplyReply")
	}

	if closeErr := readBuffer.CloseContext("CALDataNormalReplyReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataNormalReplyReply")
	}

	// Create a partially initialized instance
	_child := &_CALDataNormalReplyReply{
		ParamNumber:    paramNumber,
		Data:           data,
		_CALDataNormal: &_CALDataNormal{},
	}
	_child._CALDataNormal._CALDataNormalChildRequirements = _child
	return _child, nil
}

func (m *_CALDataNormalReplyReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataNormalReplyReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataNormalReplyReply")
		}

		// Simple Field (paramNumber)
		paramNumber := uint8(m.GetParamNumber())
		_paramNumberErr := writeBuffer.WriteUint8("paramNumber", 8, (paramNumber))
		if _paramNumberErr != nil {
			return errors.Wrap(_paramNumberErr, "Error serializing 'paramNumber' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("CALDataNormalReplyReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataNormalReplyReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataNormalReplyReply) isCALDataNormalReplyReply() bool {
	return true
}

func (m *_CALDataNormalReplyReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}