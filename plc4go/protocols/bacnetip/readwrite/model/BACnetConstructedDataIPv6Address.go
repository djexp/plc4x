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

// BACnetConstructedDataIPv6Address is the data-structure of this message
type BACnetConstructedDataIPv6Address struct {
	*BACnetConstructedData
	Ipv6Address *BACnetApplicationTagOctetString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataIPv6Address is the corresponding interface of BACnetConstructedDataIPv6Address
type IBACnetConstructedDataIPv6Address interface {
	IBACnetConstructedData
	// GetIpv6Address returns Ipv6Address (property field)
	GetIpv6Address() *BACnetApplicationTagOctetString
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

func (m *BACnetConstructedDataIPv6Address) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIPv6Address) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIPv6Address) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIPv6Address) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIPv6Address) GetIpv6Address() *BACnetApplicationTagOctetString {
	return m.Ipv6Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6Address factory function for BACnetConstructedDataIPv6Address
func NewBACnetConstructedDataIPv6Address(ipv6Address *BACnetApplicationTagOctetString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataIPv6Address {
	_result := &BACnetConstructedDataIPv6Address{
		Ipv6Address:           ipv6Address,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIPv6Address(structType interface{}) *BACnetConstructedDataIPv6Address {
	if casted, ok := structType.(BACnetConstructedDataIPv6Address); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6Address); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPv6Address(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPv6Address(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIPv6Address) GetTypeName() string {
	return "BACnetConstructedDataIPv6Address"
}

func (m *BACnetConstructedDataIPv6Address) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIPv6Address) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipv6Address)
	lengthInBits += m.Ipv6Address.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataIPv6Address) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPv6AddressParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataIPv6Address, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6Address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6Address")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6Address)
	if pullErr := readBuffer.PullContext("ipv6Address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6Address")
	}
	_ipv6Address, _ipv6AddressErr := BACnetApplicationTagParse(readBuffer)
	if _ipv6AddressErr != nil {
		return nil, errors.Wrap(_ipv6AddressErr, "Error parsing 'ipv6Address' field")
	}
	ipv6Address := CastBACnetApplicationTagOctetString(_ipv6Address)
	if closeErr := readBuffer.CloseContext("ipv6Address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6Address")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6Address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6Address")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIPv6Address{
		Ipv6Address:           CastBACnetApplicationTagOctetString(ipv6Address),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIPv6Address) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6Address"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6Address")
		}

		// Simple Field (ipv6Address)
		if pushErr := writeBuffer.PushContext("ipv6Address"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipv6Address")
		}
		_ipv6AddressErr := m.Ipv6Address.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("ipv6Address"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipv6Address")
		}
		if _ipv6AddressErr != nil {
			return errors.Wrap(_ipv6AddressErr, "Error serializing 'ipv6Address' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6Address"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6Address")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIPv6Address) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}