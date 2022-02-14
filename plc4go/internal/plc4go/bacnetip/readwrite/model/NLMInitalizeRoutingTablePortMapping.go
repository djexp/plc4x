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
type NLMInitalizeRoutingTablePortMapping struct {
	DestinationNetworkAddress uint16
	PortId                    uint8
	PortInfoLength            uint8
	PortInfo                  []byte
}

// The corresponding interface
type INLMInitalizeRoutingTablePortMapping interface {
	// GetDestinationNetworkAddress returns DestinationNetworkAddress
	GetDestinationNetworkAddress() uint16
	// GetPortId returns PortId
	GetPortId() uint8
	// GetPortInfoLength returns PortInfoLength
	GetPortInfoLength() uint8
	// GetPortInfo returns PortInfo
	GetPortInfo() []byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *NLMInitalizeRoutingTablePortMapping) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

func (m *NLMInitalizeRoutingTablePortMapping) GetPortId() uint8 {
	return m.PortId
}

func (m *NLMInitalizeRoutingTablePortMapping) GetPortInfoLength() uint8 {
	return m.PortInfoLength
}

func (m *NLMInitalizeRoutingTablePortMapping) GetPortInfo() []byte {
	return m.PortInfo
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewNLMInitalizeRoutingTablePortMapping factory function for NLMInitalizeRoutingTablePortMapping
func NewNLMInitalizeRoutingTablePortMapping(destinationNetworkAddress uint16, portId uint8, portInfoLength uint8, portInfo []byte) *NLMInitalizeRoutingTablePortMapping {
	return &NLMInitalizeRoutingTablePortMapping{DestinationNetworkAddress: destinationNetworkAddress, PortId: portId, PortInfoLength: portInfoLength, PortInfo: portInfo}
}

func CastNLMInitalizeRoutingTablePortMapping(structType interface{}) *NLMInitalizeRoutingTablePortMapping {
	castFunc := func(typ interface{}) *NLMInitalizeRoutingTablePortMapping {
		if casted, ok := typ.(NLMInitalizeRoutingTablePortMapping); ok {
			return &casted
		}
		if casted, ok := typ.(*NLMInitalizeRoutingTablePortMapping); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *NLMInitalizeRoutingTablePortMapping) GetTypeName() string {
	return "NLMInitalizeRoutingTablePortMapping"
}

func (m *NLMInitalizeRoutingTablePortMapping) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NLMInitalizeRoutingTablePortMapping) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	// Simple field (portId)
	lengthInBits += 8

	// Simple field (portInfoLength)
	lengthInBits += 8

	// Array field
	if len(m.PortInfo) > 0 {
		lengthInBits += 8 * uint16(len(m.PortInfo))
	}

	return lengthInBits
}

func (m *NLMInitalizeRoutingTablePortMapping) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMInitalizeRoutingTablePortMappingParse(readBuffer utils.ReadBuffer) (*NLMInitalizeRoutingTablePortMapping, error) {
	if pullErr := readBuffer.PullContext("NLMInitalizeRoutingTablePortMapping"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	// Simple Field (portId)
	_portId, _portIdErr := readBuffer.ReadUint8("portId", 8)
	if _portIdErr != nil {
		return nil, errors.Wrap(_portIdErr, "Error parsing 'portId' field")
	}
	portId := _portId

	// Simple Field (portInfoLength)
	_portInfoLength, _portInfoLengthErr := readBuffer.ReadUint8("portInfoLength", 8)
	if _portInfoLengthErr != nil {
		return nil, errors.Wrap(_portInfoLengthErr, "Error parsing 'portInfoLength' field")
	}
	portInfoLength := _portInfoLength
	// Byte Array field (portInfo)
	numberOfBytesportInfo := int(portInfoLength)
	portInfo, _readArrayErr := readBuffer.ReadByteArray("portInfo", numberOfBytesportInfo)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'portInfo' field")
	}

	if closeErr := readBuffer.CloseContext("NLMInitalizeRoutingTablePortMapping"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewNLMInitalizeRoutingTablePortMapping(destinationNetworkAddress, portId, portInfoLength, portInfo), nil
}

func (m *NLMInitalizeRoutingTablePortMapping) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("NLMInitalizeRoutingTablePortMapping"); pushErr != nil {
		return pushErr
	}

	// Simple Field (destinationNetworkAddress)
	destinationNetworkAddress := uint16(m.DestinationNetworkAddress)
	_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
	if _destinationNetworkAddressErr != nil {
		return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
	}

	// Simple Field (portId)
	portId := uint8(m.PortId)
	_portIdErr := writeBuffer.WriteUint8("portId", 8, (portId))
	if _portIdErr != nil {
		return errors.Wrap(_portIdErr, "Error serializing 'portId' field")
	}

	// Simple Field (portInfoLength)
	portInfoLength := uint8(m.PortInfoLength)
	_portInfoLengthErr := writeBuffer.WriteUint8("portInfoLength", 8, (portInfoLength))
	if _portInfoLengthErr != nil {
		return errors.Wrap(_portInfoLengthErr, "Error serializing 'portInfoLength' field")
	}

	// Array Field (portInfo)
	if m.PortInfo != nil {
		// Byte Array field (portInfo)
		_writeArrayErr := writeBuffer.WriteByteArray("portInfo", m.PortInfo)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'portInfo' field")
		}
	}

	if popErr := writeBuffer.PopContext("NLMInitalizeRoutingTablePortMapping"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *NLMInitalizeRoutingTablePortMapping) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
