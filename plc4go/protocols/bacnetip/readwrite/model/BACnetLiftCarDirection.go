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

// BACnetLiftCarDirection is an enum
type BACnetLiftCarDirection uint16

type IBACnetLiftCarDirection interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetLiftCarDirection_UNKNOWN                  BACnetLiftCarDirection = 0
	BACnetLiftCarDirection_NONE                     BACnetLiftCarDirection = 1
	BACnetLiftCarDirection_STOPPED                  BACnetLiftCarDirection = 2
	BACnetLiftCarDirection_UP                       BACnetLiftCarDirection = 3
	BACnetLiftCarDirection_DOWN                     BACnetLiftCarDirection = 4
	BACnetLiftCarDirection_UP_AND_DOWN              BACnetLiftCarDirection = 5
	BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE BACnetLiftCarDirection = 0xFFFF
)

var BACnetLiftCarDirectionValues []BACnetLiftCarDirection

func init() {
	_ = errors.New
	BACnetLiftCarDirectionValues = []BACnetLiftCarDirection{
		BACnetLiftCarDirection_UNKNOWN,
		BACnetLiftCarDirection_NONE,
		BACnetLiftCarDirection_STOPPED,
		BACnetLiftCarDirection_UP,
		BACnetLiftCarDirection_DOWN,
		BACnetLiftCarDirection_UP_AND_DOWN,
		BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLiftCarDirectionByValue(value uint16) BACnetLiftCarDirection {
	switch value {
	case 0:
		return BACnetLiftCarDirection_UNKNOWN
	case 0xFFFF:
		return BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetLiftCarDirection_NONE
	case 2:
		return BACnetLiftCarDirection_STOPPED
	case 3:
		return BACnetLiftCarDirection_UP
	case 4:
		return BACnetLiftCarDirection_DOWN
	case 5:
		return BACnetLiftCarDirection_UP_AND_DOWN
	}
	return 0
}

func BACnetLiftCarDirectionByName(value string) (enum BACnetLiftCarDirection, ok bool) {
	ok = true
	switch value {
	case "UNKNOWN":
		enum = BACnetLiftCarDirection_UNKNOWN
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE
	case "NONE":
		enum = BACnetLiftCarDirection_NONE
	case "STOPPED":
		enum = BACnetLiftCarDirection_STOPPED
	case "UP":
		enum = BACnetLiftCarDirection_UP
	case "DOWN":
		enum = BACnetLiftCarDirection_DOWN
	case "UP_AND_DOWN":
		enum = BACnetLiftCarDirection_UP_AND_DOWN
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetLiftCarDirectionKnows(value uint16) bool {
	for _, typeValue := range BACnetLiftCarDirectionValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLiftCarDirection(structType interface{}) BACnetLiftCarDirection {
	castFunc := func(typ interface{}) BACnetLiftCarDirection {
		if sBACnetLiftCarDirection, ok := typ.(BACnetLiftCarDirection); ok {
			return sBACnetLiftCarDirection
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLiftCarDirection) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetLiftCarDirection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLiftCarDirectionParse(readBuffer utils.ReadBuffer) (BACnetLiftCarDirection, error) {
	val, err := readBuffer.ReadUint16("BACnetLiftCarDirection", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetLiftCarDirectionByValue(val), nil
}

func (e BACnetLiftCarDirection) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetLiftCarDirection", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLiftCarDirection) PLC4XEnumName() string {
	switch e {
	case BACnetLiftCarDirection_UNKNOWN:
		return "UNKNOWN"
	case BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLiftCarDirection_NONE:
		return "NONE"
	case BACnetLiftCarDirection_STOPPED:
		return "STOPPED"
	case BACnetLiftCarDirection_UP:
		return "UP"
	case BACnetLiftCarDirection_DOWN:
		return "DOWN"
	case BACnetLiftCarDirection_UP_AND_DOWN:
		return "UP_AND_DOWN"
	}
	return ""
}

func (e BACnetLiftCarDirection) String() string {
	return e.PLC4XEnumName()
}
