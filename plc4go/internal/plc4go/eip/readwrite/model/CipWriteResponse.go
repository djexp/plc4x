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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CipWriteResponse struct {
	*CipService
	Status    uint8
	ExtStatus uint8

	// Arguments.
	ServiceLen uint16
}

// The corresponding interface
type ICipWriteResponse interface {
	// GetStatus returns Status
	GetStatus() uint8
	// GetExtStatus returns ExtStatus
	GetExtStatus() uint8
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
func (m *CipWriteResponse) Service() uint8 {
	return 0xCD
}

func (m *CipWriteResponse) GetService() uint8 {
	return 0xCD
}

func (m *CipWriteResponse) InitializeParent(parent *CipService) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CipWriteResponse) GetStatus() uint8 {
	return m.Status
}

func (m *CipWriteResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCipWriteResponse factory function for CipWriteResponse
func NewCipWriteResponse(status uint8, extStatus uint8, serviceLen uint16) *CipService {
	child := &CipWriteResponse{
		Status:     status,
		ExtStatus:  extStatus,
		CipService: NewCipService(serviceLen),
	}
	child.Child = child
	return child.CipService
}

func CastCipWriteResponse(structType interface{}) *CipWriteResponse {
	castFunc := func(typ interface{}) *CipWriteResponse {
		if casted, ok := typ.(CipWriteResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*CipWriteResponse); ok {
			return casted
		}
		if casted, ok := typ.(CipService); ok {
			return CastCipWriteResponse(casted.Child)
		}
		if casted, ok := typ.(*CipService); ok {
			return CastCipWriteResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CipWriteResponse) GetTypeName() string {
	return "CipWriteResponse"
}

func (m *CipWriteResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CipWriteResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	return lengthInBits
}

func (m *CipWriteResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CipWriteResponseParse(readBuffer utils.ReadBuffer, serviceLen uint16) (*CipService, error) {
	if pullErr := readBuffer.PullContext("CipWriteResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := _status

	// Simple Field (extStatus)
	_extStatus, _extStatusErr := readBuffer.ReadUint8("extStatus", 8)
	if _extStatusErr != nil {
		return nil, errors.Wrap(_extStatusErr, "Error parsing 'extStatus' field")
	}
	extStatus := _extStatus

	if closeErr := readBuffer.CloseContext("CipWriteResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CipWriteResponse{
		Status:     status,
		ExtStatus:  extStatus,
		CipService: &CipService{},
	}
	_child.CipService.Child = _child
	return _child.CipService, nil
}

func (m *CipWriteResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipWriteResponse"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (status)
		status := uint8(m.Status)
		_statusErr := writeBuffer.WriteUint8("status", 8, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (extStatus)
		extStatus := uint8(m.ExtStatus)
		_extStatusErr := writeBuffer.WriteUint8("extStatus", 8, (extStatus))
		if _extStatusErr != nil {
			return errors.Wrap(_extStatusErr, "Error serializing 'extStatus' field")
		}

		if popErr := writeBuffer.PopContext("CipWriteResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CipWriteResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
