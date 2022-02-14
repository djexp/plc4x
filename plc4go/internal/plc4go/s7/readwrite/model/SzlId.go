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
type SzlId struct {
	TypeClass      SzlModuleTypeClass
	SublistExtract uint8
	SublistList    SzlSublist
}

// The corresponding interface
type ISzlId interface {
	// GetTypeClass returns TypeClass
	GetTypeClass() SzlModuleTypeClass
	// GetSublistExtract returns SublistExtract
	GetSublistExtract() uint8
	// GetSublistList returns SublistList
	GetSublistList() SzlSublist
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
func (m *SzlId) GetTypeClass() SzlModuleTypeClass {
	return m.TypeClass
}

func (m *SzlId) GetSublistExtract() uint8 {
	return m.SublistExtract
}

func (m *SzlId) GetSublistList() SzlSublist {
	return m.SublistList
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewSzlId factory function for SzlId
func NewSzlId(typeClass SzlModuleTypeClass, sublistExtract uint8, sublistList SzlSublist) *SzlId {
	return &SzlId{TypeClass: typeClass, SublistExtract: sublistExtract, SublistList: sublistList}
}

func CastSzlId(structType interface{}) *SzlId {
	castFunc := func(typ interface{}) *SzlId {
		if casted, ok := typ.(SzlId); ok {
			return &casted
		}
		if casted, ok := typ.(*SzlId); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SzlId) GetTypeName() string {
	return "SzlId"
}

func (m *SzlId) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SzlId) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (typeClass)
	lengthInBits += 4

	// Simple field (sublistExtract)
	lengthInBits += 4

	// Simple field (sublistList)
	lengthInBits += 8

	return lengthInBits
}

func (m *SzlId) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SzlIdParse(readBuffer utils.ReadBuffer) (*SzlId, error) {
	if pullErr := readBuffer.PullContext("SzlId"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (typeClass)
	if pullErr := readBuffer.PullContext("typeClass"); pullErr != nil {
		return nil, pullErr
	}
	_typeClass, _typeClassErr := SzlModuleTypeClassParse(readBuffer)
	if _typeClassErr != nil {
		return nil, errors.Wrap(_typeClassErr, "Error parsing 'typeClass' field")
	}
	typeClass := _typeClass
	if closeErr := readBuffer.CloseContext("typeClass"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (sublistExtract)
	_sublistExtract, _sublistExtractErr := readBuffer.ReadUint8("sublistExtract", 4)
	if _sublistExtractErr != nil {
		return nil, errors.Wrap(_sublistExtractErr, "Error parsing 'sublistExtract' field")
	}
	sublistExtract := _sublistExtract

	// Simple Field (sublistList)
	if pullErr := readBuffer.PullContext("sublistList"); pullErr != nil {
		return nil, pullErr
	}
	_sublistList, _sublistListErr := SzlSublistParse(readBuffer)
	if _sublistListErr != nil {
		return nil, errors.Wrap(_sublistListErr, "Error parsing 'sublistList' field")
	}
	sublistList := _sublistList
	if closeErr := readBuffer.CloseContext("sublistList"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("SzlId"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewSzlId(typeClass, sublistExtract, sublistList), nil
}

func (m *SzlId) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("SzlId"); pushErr != nil {
		return pushErr
	}

	// Simple Field (typeClass)
	if pushErr := writeBuffer.PushContext("typeClass"); pushErr != nil {
		return pushErr
	}
	_typeClassErr := m.TypeClass.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("typeClass"); popErr != nil {
		return popErr
	}
	if _typeClassErr != nil {
		return errors.Wrap(_typeClassErr, "Error serializing 'typeClass' field")
	}

	// Simple Field (sublistExtract)
	sublistExtract := uint8(m.SublistExtract)
	_sublistExtractErr := writeBuffer.WriteUint8("sublistExtract", 4, (sublistExtract))
	if _sublistExtractErr != nil {
		return errors.Wrap(_sublistExtractErr, "Error serializing 'sublistExtract' field")
	}

	// Simple Field (sublistList)
	if pushErr := writeBuffer.PushContext("sublistList"); pushErr != nil {
		return pushErr
	}
	_sublistListErr := m.SublistList.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("sublistList"); popErr != nil {
		return popErr
	}
	if _sublistListErr != nil {
		return errors.Wrap(_sublistListErr, "Error serializing 'sublistList' field")
	}

	if popErr := writeBuffer.PopContext("SzlId"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *SzlId) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
