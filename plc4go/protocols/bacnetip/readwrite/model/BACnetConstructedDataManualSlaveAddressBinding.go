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

// BACnetConstructedDataManualSlaveAddressBinding is the corresponding interface of BACnetConstructedDataManualSlaveAddressBinding
type BACnetConstructedDataManualSlaveAddressBinding interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetManualSlaveAddressBinding returns ManualSlaveAddressBinding (property field)
	GetManualSlaveAddressBinding() []BACnetAddressBinding
}

// BACnetConstructedDataManualSlaveAddressBindingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataManualSlaveAddressBinding.
// This is useful for switch cases.
type BACnetConstructedDataManualSlaveAddressBindingExactly interface {
	BACnetConstructedDataManualSlaveAddressBinding
	isBACnetConstructedDataManualSlaveAddressBinding() bool
}

// _BACnetConstructedDataManualSlaveAddressBinding is the data-structure of this message
type _BACnetConstructedDataManualSlaveAddressBinding struct {
	*_BACnetConstructedData
	ManualSlaveAddressBinding []BACnetAddressBinding
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MANUAL_SLAVE_ADDRESS_BINDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetManualSlaveAddressBinding() []BACnetAddressBinding {
	return m.ManualSlaveAddressBinding
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataManualSlaveAddressBinding factory function for _BACnetConstructedDataManualSlaveAddressBinding
func NewBACnetConstructedDataManualSlaveAddressBinding(manualSlaveAddressBinding []BACnetAddressBinding, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataManualSlaveAddressBinding {
	_result := &_BACnetConstructedDataManualSlaveAddressBinding{
		ManualSlaveAddressBinding: manualSlaveAddressBinding,
		_BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataManualSlaveAddressBinding(structType interface{}) BACnetConstructedDataManualSlaveAddressBinding {
	if casted, ok := structType.(BACnetConstructedDataManualSlaveAddressBinding); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataManualSlaveAddressBinding); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetTypeName() string {
	return "BACnetConstructedDataManualSlaveAddressBinding"
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.ManualSlaveAddressBinding) > 0 {
		for _, element := range m.ManualSlaveAddressBinding {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataManualSlaveAddressBindingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataManualSlaveAddressBinding, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataManualSlaveAddressBinding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataManualSlaveAddressBinding")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (manualSlaveAddressBinding)
	if pullErr := readBuffer.PullContext("manualSlaveAddressBinding", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for manualSlaveAddressBinding")
	}
	// Terminated array
	var manualSlaveAddressBinding []BACnetAddressBinding
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAddressBindingParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'manualSlaveAddressBinding' field of BACnetConstructedDataManualSlaveAddressBinding")
			}
			manualSlaveAddressBinding = append(manualSlaveAddressBinding, _item.(BACnetAddressBinding))

		}
	}
	if closeErr := readBuffer.CloseContext("manualSlaveAddressBinding", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for manualSlaveAddressBinding")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataManualSlaveAddressBinding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataManualSlaveAddressBinding")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataManualSlaveAddressBinding{
		ManualSlaveAddressBinding: manualSlaveAddressBinding,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataManualSlaveAddressBinding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataManualSlaveAddressBinding")
		}

		// Array Field (manualSlaveAddressBinding)
		if pushErr := writeBuffer.PushContext("manualSlaveAddressBinding", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for manualSlaveAddressBinding")
		}
		for _, _element := range m.GetManualSlaveAddressBinding() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'manualSlaveAddressBinding' field")
			}
		}
		if popErr := writeBuffer.PopContext("manualSlaveAddressBinding", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for manualSlaveAddressBinding")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataManualSlaveAddressBinding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataManualSlaveAddressBinding")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) isBACnetConstructedDataManualSlaveAddressBinding() bool {
	return true
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
