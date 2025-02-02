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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const RequestDirectCommandAccess_AT byte = 0x40

// RequestDirectCommandAccess is the corresponding interface of RequestDirectCommandAccess
type RequestDirectCommandAccess interface {
	utils.LengthAware
	utils.Serializable
	Request
	// GetCalData returns CalData (property field)
	GetCalData() CALData
}

// RequestDirectCommandAccessExactly can be used when we want exactly this type and not a type which fulfills RequestDirectCommandAccess.
// This is useful for switch cases.
type RequestDirectCommandAccessExactly interface {
	RequestDirectCommandAccess
	isRequestDirectCommandAccess() bool
}

// _RequestDirectCommandAccess is the data-structure of this message
type _RequestDirectCommandAccess struct {
	*_Request
	CalData CALData

	// Arguments.
	PayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestDirectCommandAccess) InitializeParent(parent Request, peekedByte RequestType, termination RequestTermination) {
	m.PeekedByte = peekedByte
	m.Termination = termination
}

func (m *_RequestDirectCommandAccess) GetParent() Request {
	return m._Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestDirectCommandAccess) GetCalData() CALData {
	return m.CalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestDirectCommandAccess) GetAt() byte {
	return RequestDirectCommandAccess_AT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestDirectCommandAccess factory function for _RequestDirectCommandAccess
func NewRequestDirectCommandAccess(calData CALData, peekedByte RequestType, termination RequestTermination, srchk bool, messageLength uint16, payloadLength uint16) *_RequestDirectCommandAccess {
	_result := &_RequestDirectCommandAccess{
		CalData:  calData,
		_Request: NewRequest(peekedByte, termination, srchk, messageLength),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestDirectCommandAccess(structType interface{}) RequestDirectCommandAccess {
	if casted, ok := structType.(RequestDirectCommandAccess); ok {
		return casted
	}
	if casted, ok := structType.(*RequestDirectCommandAccess); ok {
		return *casted
	}
	return nil
}

func (m *_RequestDirectCommandAccess) GetTypeName() string {
	return "RequestDirectCommandAccess"
}

func (m *_RequestDirectCommandAccess) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestDirectCommandAccess) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (at)
	lengthInBits += 8

	// Manual Field (calData)
	lengthInBits += uint16(int32(m.GetLengthInBytes()) * int32(int32(2)))

	return lengthInBits
}

func (m *_RequestDirectCommandAccess) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestDirectCommandAccessParse(readBuffer utils.ReadBuffer, srchk bool, messageLength uint16, payloadLength uint16) (RequestDirectCommandAccess, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestDirectCommandAccess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestDirectCommandAccess")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (at)
	at, _atErr := readBuffer.ReadByte("at")
	if _atErr != nil {
		return nil, errors.Wrap(_atErr, "Error parsing 'at' field of RequestDirectCommandAccess")
	}
	if at != RequestDirectCommandAccess_AT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestDirectCommandAccess_AT) + " but got " + fmt.Sprintf("%d", at))
	}

	// Manual Field (calData)
	_calData, _calDataErr := ReadCALData(readBuffer, payloadLength)
	if _calDataErr != nil {
		return nil, errors.Wrap(_calDataErr, "Error parsing 'calData' field of RequestDirectCommandAccess")
	}
	calData := _calData.(CALData)

	if closeErr := readBuffer.CloseContext("RequestDirectCommandAccess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestDirectCommandAccess")
	}

	// Create a partially initialized instance
	_child := &_RequestDirectCommandAccess{
		CalData: calData,
		_Request: &_Request{
			Srchk:         srchk,
			MessageLength: messageLength,
		},
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestDirectCommandAccess) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestDirectCommandAccess"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestDirectCommandAccess")
		}

		// Const Field (at)
		_atErr := writeBuffer.WriteByte("at", 0x40)
		if _atErr != nil {
			return errors.Wrap(_atErr, "Error serializing 'at' field")
		}

		// Manual Field (calData)
		_calDataErr := WriteCALData(writeBuffer, m.GetCalData())
		if _calDataErr != nil {
			return errors.Wrap(_calDataErr, "Error serializing 'calData' field")
		}

		if popErr := writeBuffer.PopContext("RequestDirectCommandAccess"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestDirectCommandAccess")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_RequestDirectCommandAccess) isRequestDirectCommandAccess() bool {
	return true
}

func (m *_RequestDirectCommandAccess) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
