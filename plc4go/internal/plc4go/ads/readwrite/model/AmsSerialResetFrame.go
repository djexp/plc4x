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
type AmsSerialResetFrame struct {
	MagicCookie        uint16
	TransmitterAddress int8
	ReceiverAddress    int8
	FragmentNumber     int8
	Length             int8
	Crc                uint16
}

// The corresponding interface
type IAmsSerialResetFrame interface {
	// GetMagicCookie returns MagicCookie
	GetMagicCookie() uint16
	// GetTransmitterAddress returns TransmitterAddress
	GetTransmitterAddress() int8
	// GetReceiverAddress returns ReceiverAddress
	GetReceiverAddress() int8
	// GetFragmentNumber returns FragmentNumber
	GetFragmentNumber() int8
	// GetLength returns Length
	GetLength() int8
	// GetCrc returns Crc
	GetCrc() uint16
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
func (m *AmsSerialResetFrame) GetMagicCookie() uint16 {
	return m.MagicCookie
}

func (m *AmsSerialResetFrame) GetTransmitterAddress() int8 {
	return m.TransmitterAddress
}

func (m *AmsSerialResetFrame) GetReceiverAddress() int8 {
	return m.ReceiverAddress
}

func (m *AmsSerialResetFrame) GetFragmentNumber() int8 {
	return m.FragmentNumber
}

func (m *AmsSerialResetFrame) GetLength() int8 {
	return m.Length
}

func (m *AmsSerialResetFrame) GetCrc() uint16 {
	return m.Crc
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAmsSerialResetFrame factory function for AmsSerialResetFrame
func NewAmsSerialResetFrame(magicCookie uint16, transmitterAddress int8, receiverAddress int8, fragmentNumber int8, length int8, crc uint16) *AmsSerialResetFrame {
	return &AmsSerialResetFrame{MagicCookie: magicCookie, TransmitterAddress: transmitterAddress, ReceiverAddress: receiverAddress, FragmentNumber: fragmentNumber, Length: length, Crc: crc}
}

func CastAmsSerialResetFrame(structType interface{}) *AmsSerialResetFrame {
	castFunc := func(typ interface{}) *AmsSerialResetFrame {
		if casted, ok := typ.(AmsSerialResetFrame); ok {
			return &casted
		}
		if casted, ok := typ.(*AmsSerialResetFrame); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AmsSerialResetFrame) GetTypeName() string {
	return "AmsSerialResetFrame"
}

func (m *AmsSerialResetFrame) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AmsSerialResetFrame) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (magicCookie)
	lengthInBits += 16

	// Simple field (transmitterAddress)
	lengthInBits += 8

	// Simple field (receiverAddress)
	lengthInBits += 8

	// Simple field (fragmentNumber)
	lengthInBits += 8

	// Simple field (length)
	lengthInBits += 8

	// Simple field (crc)
	lengthInBits += 16

	return lengthInBits
}

func (m *AmsSerialResetFrame) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AmsSerialResetFrameParse(readBuffer utils.ReadBuffer) (*AmsSerialResetFrame, error) {
	if pullErr := readBuffer.PullContext("AmsSerialResetFrame"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (magicCookie)
	_magicCookie, _magicCookieErr := readBuffer.ReadUint16("magicCookie", 16)
	if _magicCookieErr != nil {
		return nil, errors.Wrap(_magicCookieErr, "Error parsing 'magicCookie' field")
	}
	magicCookie := _magicCookie

	// Simple Field (transmitterAddress)
	_transmitterAddress, _transmitterAddressErr := readBuffer.ReadInt8("transmitterAddress", 8)
	if _transmitterAddressErr != nil {
		return nil, errors.Wrap(_transmitterAddressErr, "Error parsing 'transmitterAddress' field")
	}
	transmitterAddress := _transmitterAddress

	// Simple Field (receiverAddress)
	_receiverAddress, _receiverAddressErr := readBuffer.ReadInt8("receiverAddress", 8)
	if _receiverAddressErr != nil {
		return nil, errors.Wrap(_receiverAddressErr, "Error parsing 'receiverAddress' field")
	}
	receiverAddress := _receiverAddress

	// Simple Field (fragmentNumber)
	_fragmentNumber, _fragmentNumberErr := readBuffer.ReadInt8("fragmentNumber", 8)
	if _fragmentNumberErr != nil {
		return nil, errors.Wrap(_fragmentNumberErr, "Error parsing 'fragmentNumber' field")
	}
	fragmentNumber := _fragmentNumber

	// Simple Field (length)
	_length, _lengthErr := readBuffer.ReadInt8("length", 8)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}
	length := _length

	// Simple Field (crc)
	_crc, _crcErr := readBuffer.ReadUint16("crc", 16)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field")
	}
	crc := _crc

	if closeErr := readBuffer.CloseContext("AmsSerialResetFrame"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAmsSerialResetFrame(magicCookie, transmitterAddress, receiverAddress, fragmentNumber, length, crc), nil
}

func (m *AmsSerialResetFrame) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AmsSerialResetFrame"); pushErr != nil {
		return pushErr
	}

	// Simple Field (magicCookie)
	magicCookie := uint16(m.MagicCookie)
	_magicCookieErr := writeBuffer.WriteUint16("magicCookie", 16, (magicCookie))
	if _magicCookieErr != nil {
		return errors.Wrap(_magicCookieErr, "Error serializing 'magicCookie' field")
	}

	// Simple Field (transmitterAddress)
	transmitterAddress := int8(m.TransmitterAddress)
	_transmitterAddressErr := writeBuffer.WriteInt8("transmitterAddress", 8, (transmitterAddress))
	if _transmitterAddressErr != nil {
		return errors.Wrap(_transmitterAddressErr, "Error serializing 'transmitterAddress' field")
	}

	// Simple Field (receiverAddress)
	receiverAddress := int8(m.ReceiverAddress)
	_receiverAddressErr := writeBuffer.WriteInt8("receiverAddress", 8, (receiverAddress))
	if _receiverAddressErr != nil {
		return errors.Wrap(_receiverAddressErr, "Error serializing 'receiverAddress' field")
	}

	// Simple Field (fragmentNumber)
	fragmentNumber := int8(m.FragmentNumber)
	_fragmentNumberErr := writeBuffer.WriteInt8("fragmentNumber", 8, (fragmentNumber))
	if _fragmentNumberErr != nil {
		return errors.Wrap(_fragmentNumberErr, "Error serializing 'fragmentNumber' field")
	}

	// Simple Field (length)
	length := int8(m.Length)
	_lengthErr := writeBuffer.WriteInt8("length", 8, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (crc)
	crc := uint16(m.Crc)
	_crcErr := writeBuffer.WriteUint16("crc", 16, (crc))
	if _crcErr != nil {
		return errors.Wrap(_crcErr, "Error serializing 'crc' field")
	}

	if popErr := writeBuffer.PopContext("AmsSerialResetFrame"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AmsSerialResetFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
