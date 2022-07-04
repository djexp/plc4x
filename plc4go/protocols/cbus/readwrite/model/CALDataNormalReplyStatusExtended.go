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

// CALDataNormalReplyStatusExtended is the corresponding interface of CALDataNormalReplyStatusExtended
type CALDataNormalReplyStatusExtended interface {
	utils.LengthAware
	utils.Serializable
	CALDataNormal
	// GetEncoding returns Encoding (property field)
	GetEncoding() uint8
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetBlockStart returns BlockStart (property field)
	GetBlockStart() uint8
	// GetData returns Data (property field)
	GetData() []byte
}

// CALDataNormalReplyStatusExtendedExactly can be used when we want exactly this type and not a type which fulfills CALDataNormalReplyStatusExtended.
// This is useful for switch cases.
type CALDataNormalReplyStatusExtendedExactly interface {
	CALDataNormalReplyStatusExtended
	isCALDataNormalReplyStatusExtended() bool
}

// _CALDataNormalReplyStatusExtended is the data-structure of this message
type _CALDataNormalReplyStatusExtended struct {
	*_CALDataNormal
	Encoding    uint8
	Application ApplicationIdContainer
	BlockStart  uint8
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

func (m *_CALDataNormalReplyStatusExtended) InitializeParent(parent CALDataNormal, commandTypeContainer CALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_CALDataNormalReplyStatusExtended) GetParent() CALDataNormal {
	return m._CALDataNormal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataNormalReplyStatusExtended) GetEncoding() uint8 {
	return m.Encoding
}

func (m *_CALDataNormalReplyStatusExtended) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CALDataNormalReplyStatusExtended) GetBlockStart() uint8 {
	return m.BlockStart
}

func (m *_CALDataNormalReplyStatusExtended) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataNormalReplyStatusExtended factory function for _CALDataNormalReplyStatusExtended
func NewCALDataNormalReplyStatusExtended(encoding uint8, application ApplicationIdContainer, blockStart uint8, data []byte, commandTypeContainer CALCommandTypeContainer) *_CALDataNormalReplyStatusExtended {
	_result := &_CALDataNormalReplyStatusExtended{
		Encoding:       encoding,
		Application:    application,
		BlockStart:     blockStart,
		Data:           data,
		_CALDataNormal: NewCALDataNormal(commandTypeContainer),
	}
	_result._CALDataNormal._CALDataNormalChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataNormalReplyStatusExtended(structType interface{}) CALDataNormalReplyStatusExtended {
	if casted, ok := structType.(CALDataNormalReplyStatusExtended); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataNormalReplyStatusExtended); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataNormalReplyStatusExtended) GetTypeName() string {
	return "CALDataNormalReplyStatusExtended"
}

func (m *_CALDataNormalReplyStatusExtended) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataNormalReplyStatusExtended) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (encoding)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Simple field (blockStart)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CALDataNormalReplyStatusExtended) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataNormalReplyStatusExtendedParse(readBuffer utils.ReadBuffer, commandTypeContainer CALCommandTypeContainer) (CALDataNormalReplyStatusExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataNormalReplyStatusExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataNormalReplyStatusExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (encoding)
	_encoding, _encodingErr := readBuffer.ReadUint8("encoding", 8)
	if _encodingErr != nil {
		return nil, errors.Wrap(_encodingErr, "Error parsing 'encoding' field of CALDataNormalReplyStatusExtended")
	}
	encoding := _encoding

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParse(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of CALDataNormalReplyStatusExtended")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Simple Field (blockStart)
	_blockStart, _blockStartErr := readBuffer.ReadUint8("blockStart", 8)
	if _blockStartErr != nil {
		return nil, errors.Wrap(_blockStartErr, "Error parsing 'blockStart' field of CALDataNormalReplyStatusExtended")
	}
	blockStart := _blockStart
	// Byte Array field (data)
	numberOfBytesdata := int(commandTypeContainer.NumBytes())
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of CALDataNormalReplyStatusExtended")
	}

	if closeErr := readBuffer.CloseContext("CALDataNormalReplyStatusExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataNormalReplyStatusExtended")
	}

	// Create a partially initialized instance
	_child := &_CALDataNormalReplyStatusExtended{
		Encoding:       encoding,
		Application:    application,
		BlockStart:     blockStart,
		Data:           data,
		_CALDataNormal: &_CALDataNormal{},
	}
	_child._CALDataNormal._CALDataNormalChildRequirements = _child
	return _child, nil
}

func (m *_CALDataNormalReplyStatusExtended) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataNormalReplyStatusExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataNormalReplyStatusExtended")
		}

		// Simple Field (encoding)
		encoding := uint8(m.GetEncoding())
		_encodingErr := writeBuffer.WriteUint8("encoding", 8, (encoding))
		if _encodingErr != nil {
			return errors.Wrap(_encodingErr, "Error serializing 'encoding' field")
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Simple Field (blockStart)
		blockStart := uint8(m.GetBlockStart())
		_blockStartErr := writeBuffer.WriteUint8("blockStart", 8, (blockStart))
		if _blockStartErr != nil {
			return errors.Wrap(_blockStartErr, "Error serializing 'blockStart' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("CALDataNormalReplyStatusExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataNormalReplyStatusExtended")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataNormalReplyStatusExtended) isCALDataNormalReplyStatusExtended() bool {
	return true
}

func (m *_CALDataNormalReplyStatusExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}