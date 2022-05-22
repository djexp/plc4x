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

// BACnetProgramError is an enum
type BACnetProgramError uint16

type IBACnetProgramError interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetProgramError_NORMAL                   BACnetProgramError = 0
	BACnetProgramError_LOAD_FAILED              BACnetProgramError = 1
	BACnetProgramError_INTERNAL                 BACnetProgramError = 2
	BACnetProgramError_PROGRAM                  BACnetProgramError = 3
	BACnetProgramError_OTHER                    BACnetProgramError = 4
	BACnetProgramError_VENDOR_PROPRIETARY_VALUE BACnetProgramError = 0xFFFF
)

var BACnetProgramErrorValues []BACnetProgramError

func init() {
	_ = errors.New
	BACnetProgramErrorValues = []BACnetProgramError{
		BACnetProgramError_NORMAL,
		BACnetProgramError_LOAD_FAILED,
		BACnetProgramError_INTERNAL,
		BACnetProgramError_PROGRAM,
		BACnetProgramError_OTHER,
		BACnetProgramError_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetProgramErrorByValue(value uint16) BACnetProgramError {
	switch value {
	case 0:
		return BACnetProgramError_NORMAL
	case 0xFFFF:
		return BACnetProgramError_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetProgramError_LOAD_FAILED
	case 2:
		return BACnetProgramError_INTERNAL
	case 3:
		return BACnetProgramError_PROGRAM
	case 4:
		return BACnetProgramError_OTHER
	}
	return 0
}

func BACnetProgramErrorByName(value string) BACnetProgramError {
	switch value {
	case "NORMAL":
		return BACnetProgramError_NORMAL
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetProgramError_VENDOR_PROPRIETARY_VALUE
	case "LOAD_FAILED":
		return BACnetProgramError_LOAD_FAILED
	case "INTERNAL":
		return BACnetProgramError_INTERNAL
	case "PROGRAM":
		return BACnetProgramError_PROGRAM
	case "OTHER":
		return BACnetProgramError_OTHER
	}
	return 0
}

func BACnetProgramErrorKnows(value uint16) bool {
	for _, typeValue := range BACnetProgramErrorValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetProgramError(structType interface{}) BACnetProgramError {
	castFunc := func(typ interface{}) BACnetProgramError {
		if sBACnetProgramError, ok := typ.(BACnetProgramError); ok {
			return sBACnetProgramError
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetProgramError) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetProgramError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetProgramErrorParse(readBuffer utils.ReadBuffer) (BACnetProgramError, error) {
	val, err := readBuffer.ReadUint16("BACnetProgramError", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetProgramErrorByValue(val), nil
}

func (e BACnetProgramError) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetProgramError", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetProgramError) name() string {
	switch e {
	case BACnetProgramError_NORMAL:
		return "NORMAL"
	case BACnetProgramError_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetProgramError_LOAD_FAILED:
		return "LOAD_FAILED"
	case BACnetProgramError_INTERNAL:
		return "INTERNAL"
	case BACnetProgramError_PROGRAM:
		return "PROGRAM"
	case BACnetProgramError_OTHER:
		return "OTHER"
	}
	return ""
}

func (e BACnetProgramError) String() string {
	return e.name()
}