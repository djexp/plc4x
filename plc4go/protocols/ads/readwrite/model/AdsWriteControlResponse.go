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

// AdsWriteControlResponse is the data-structure of this message
type AdsWriteControlResponse struct {
	*AdsData
	Result ReturnCode
}

// IAdsWriteControlResponse is the corresponding interface of AdsWriteControlResponse
type IAdsWriteControlResponse interface {
	IAdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
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

func (m *AdsWriteControlResponse) GetCommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *AdsWriteControlResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *AdsWriteControlResponse) InitializeParent(parent *AdsData) {}

func (m *AdsWriteControlResponse) GetParent() *AdsData {
	return m.AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AdsWriteControlResponse) GetResult() ReturnCode {
	return m.Result
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsWriteControlResponse factory function for AdsWriteControlResponse
func NewAdsWriteControlResponse(result ReturnCode) *AdsWriteControlResponse {
	_result := &AdsWriteControlResponse{
		Result:  result,
		AdsData: NewAdsData(),
	}
	_result.Child = _result
	return _result
}

func CastAdsWriteControlResponse(structType interface{}) *AdsWriteControlResponse {
	if casted, ok := structType.(AdsWriteControlResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsWriteControlResponse); ok {
		return casted
	}
	if casted, ok := structType.(AdsData); ok {
		return CastAdsWriteControlResponse(casted.Child)
	}
	if casted, ok := structType.(*AdsData); ok {
		return CastAdsWriteControlResponse(casted.Child)
	}
	return nil
}

func (m *AdsWriteControlResponse) GetTypeName() string {
	return "AdsWriteControlResponse"
}

func (m *AdsWriteControlResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsWriteControlResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsWriteControlResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsWriteControlResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsWriteControlResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsWriteControlResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, pullErr
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AdsWriteControlResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsWriteControlResponse{
		Result:  result,
		AdsData: &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child, nil
}

func (m *AdsWriteControlResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteControlResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return pushErr
		}
		_resultErr := m.Result.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return popErr
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		if popErr := writeBuffer.PopContext("AdsWriteControlResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsWriteControlResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}