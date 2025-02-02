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

// BACnetConstructedDataSupportedSecurityAlgorithms is the corresponding interface of BACnetConstructedDataSupportedSecurityAlgorithms
type BACnetConstructedDataSupportedSecurityAlgorithms interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSupportedSecurityAlgorithms returns SupportedSecurityAlgorithms (property field)
	GetSupportedSecurityAlgorithms() []BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataSupportedSecurityAlgorithmsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSupportedSecurityAlgorithms.
// This is useful for switch cases.
type BACnetConstructedDataSupportedSecurityAlgorithmsExactly interface {
	BACnetConstructedDataSupportedSecurityAlgorithms
	isBACnetConstructedDataSupportedSecurityAlgorithms() bool
}

// _BACnetConstructedDataSupportedSecurityAlgorithms is the data-structure of this message
type _BACnetConstructedDataSupportedSecurityAlgorithms struct {
	*_BACnetConstructedData
	SupportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUPPORTED_SECURITY_ALGORITHMS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetSupportedSecurityAlgorithms() []BACnetApplicationTagUnsignedInteger {
	return m.SupportedSecurityAlgorithms
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSupportedSecurityAlgorithms factory function for _BACnetConstructedDataSupportedSecurityAlgorithms
func NewBACnetConstructedDataSupportedSecurityAlgorithms(supportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSupportedSecurityAlgorithms {
	_result := &_BACnetConstructedDataSupportedSecurityAlgorithms{
		SupportedSecurityAlgorithms: supportedSecurityAlgorithms,
		_BACnetConstructedData:      NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSupportedSecurityAlgorithms(structType interface{}) BACnetConstructedDataSupportedSecurityAlgorithms {
	if casted, ok := structType.(BACnetConstructedDataSupportedSecurityAlgorithms); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSupportedSecurityAlgorithms); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetTypeName() string {
	return "BACnetConstructedDataSupportedSecurityAlgorithms"
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.SupportedSecurityAlgorithms) > 0 {
		for _, element := range m.SupportedSecurityAlgorithms {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSupportedSecurityAlgorithmsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSupportedSecurityAlgorithms, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSupportedSecurityAlgorithms")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (supportedSecurityAlgorithms)
	if pullErr := readBuffer.PullContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for supportedSecurityAlgorithms")
	}
	// Terminated array
	var supportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'supportedSecurityAlgorithms' field of BACnetConstructedDataSupportedSecurityAlgorithms")
			}
			supportedSecurityAlgorithms = append(supportedSecurityAlgorithms, _item.(BACnetApplicationTagUnsignedInteger))

		}
	}
	if closeErr := readBuffer.CloseContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for supportedSecurityAlgorithms")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSupportedSecurityAlgorithms"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSupportedSecurityAlgorithms")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSupportedSecurityAlgorithms{
		SupportedSecurityAlgorithms: supportedSecurityAlgorithms,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSupportedSecurityAlgorithms")
		}

		// Array Field (supportedSecurityAlgorithms)
		if pushErr := writeBuffer.PushContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for supportedSecurityAlgorithms")
		}
		for _, _element := range m.GetSupportedSecurityAlgorithms() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'supportedSecurityAlgorithms' field")
			}
		}
		if popErr := writeBuffer.PopContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for supportedSecurityAlgorithms")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSupportedSecurityAlgorithms"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSupportedSecurityAlgorithms")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) isBACnetConstructedDataSupportedSecurityAlgorithms() bool {
	return true
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
