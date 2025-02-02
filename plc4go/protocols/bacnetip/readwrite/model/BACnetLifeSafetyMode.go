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

// BACnetLifeSafetyMode is an enum
type BACnetLifeSafetyMode uint16

type IBACnetLifeSafetyMode interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetLifeSafetyMode_OFF                        BACnetLifeSafetyMode = 0
	BACnetLifeSafetyMode_ON                         BACnetLifeSafetyMode = 1
	BACnetLifeSafetyMode_TEST                       BACnetLifeSafetyMode = 2
	BACnetLifeSafetyMode_MANNED                     BACnetLifeSafetyMode = 3
	BACnetLifeSafetyMode_UNMANNED                   BACnetLifeSafetyMode = 4
	BACnetLifeSafetyMode_ARMED                      BACnetLifeSafetyMode = 5
	BACnetLifeSafetyMode_DISARMED                   BACnetLifeSafetyMode = 6
	BACnetLifeSafetyMode_PREARMED                   BACnetLifeSafetyMode = 7
	BACnetLifeSafetyMode_SLOW                       BACnetLifeSafetyMode = 8
	BACnetLifeSafetyMode_FAST                       BACnetLifeSafetyMode = 9
	BACnetLifeSafetyMode_DISCONNECTED               BACnetLifeSafetyMode = 10
	BACnetLifeSafetyMode_ENABLED                    BACnetLifeSafetyMode = 11
	BACnetLifeSafetyMode_DISABLED                   BACnetLifeSafetyMode = 12
	BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED BACnetLifeSafetyMode = 13
	BACnetLifeSafetyMode_DEFAULT                    BACnetLifeSafetyMode = 14
	BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE   BACnetLifeSafetyMode = 0xFFFF
)

var BACnetLifeSafetyModeValues []BACnetLifeSafetyMode

func init() {
	_ = errors.New
	BACnetLifeSafetyModeValues = []BACnetLifeSafetyMode{
		BACnetLifeSafetyMode_OFF,
		BACnetLifeSafetyMode_ON,
		BACnetLifeSafetyMode_TEST,
		BACnetLifeSafetyMode_MANNED,
		BACnetLifeSafetyMode_UNMANNED,
		BACnetLifeSafetyMode_ARMED,
		BACnetLifeSafetyMode_DISARMED,
		BACnetLifeSafetyMode_PREARMED,
		BACnetLifeSafetyMode_SLOW,
		BACnetLifeSafetyMode_FAST,
		BACnetLifeSafetyMode_DISCONNECTED,
		BACnetLifeSafetyMode_ENABLED,
		BACnetLifeSafetyMode_DISABLED,
		BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED,
		BACnetLifeSafetyMode_DEFAULT,
		BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLifeSafetyModeByValue(value uint16) BACnetLifeSafetyMode {
	switch value {
	case 0:
		return BACnetLifeSafetyMode_OFF
	case 0xFFFF:
		return BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetLifeSafetyMode_ON
	case 10:
		return BACnetLifeSafetyMode_DISCONNECTED
	case 11:
		return BACnetLifeSafetyMode_ENABLED
	case 12:
		return BACnetLifeSafetyMode_DISABLED
	case 13:
		return BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED
	case 14:
		return BACnetLifeSafetyMode_DEFAULT
	case 2:
		return BACnetLifeSafetyMode_TEST
	case 3:
		return BACnetLifeSafetyMode_MANNED
	case 4:
		return BACnetLifeSafetyMode_UNMANNED
	case 5:
		return BACnetLifeSafetyMode_ARMED
	case 6:
		return BACnetLifeSafetyMode_DISARMED
	case 7:
		return BACnetLifeSafetyMode_PREARMED
	case 8:
		return BACnetLifeSafetyMode_SLOW
	case 9:
		return BACnetLifeSafetyMode_FAST
	}
	return 0
}

func BACnetLifeSafetyModeByName(value string) (enum BACnetLifeSafetyMode, ok bool) {
	ok = true
	switch value {
	case "OFF":
		enum = BACnetLifeSafetyMode_OFF
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE
	case "ON":
		enum = BACnetLifeSafetyMode_ON
	case "DISCONNECTED":
		enum = BACnetLifeSafetyMode_DISCONNECTED
	case "ENABLED":
		enum = BACnetLifeSafetyMode_ENABLED
	case "DISABLED":
		enum = BACnetLifeSafetyMode_DISABLED
	case "AUTOMATIC_RELEASE_DISABLED":
		enum = BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED
	case "DEFAULT":
		enum = BACnetLifeSafetyMode_DEFAULT
	case "TEST":
		enum = BACnetLifeSafetyMode_TEST
	case "MANNED":
		enum = BACnetLifeSafetyMode_MANNED
	case "UNMANNED":
		enum = BACnetLifeSafetyMode_UNMANNED
	case "ARMED":
		enum = BACnetLifeSafetyMode_ARMED
	case "DISARMED":
		enum = BACnetLifeSafetyMode_DISARMED
	case "PREARMED":
		enum = BACnetLifeSafetyMode_PREARMED
	case "SLOW":
		enum = BACnetLifeSafetyMode_SLOW
	case "FAST":
		enum = BACnetLifeSafetyMode_FAST
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetLifeSafetyModeKnows(value uint16) bool {
	for _, typeValue := range BACnetLifeSafetyModeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLifeSafetyMode(structType interface{}) BACnetLifeSafetyMode {
	castFunc := func(typ interface{}) BACnetLifeSafetyMode {
		if sBACnetLifeSafetyMode, ok := typ.(BACnetLifeSafetyMode); ok {
			return sBACnetLifeSafetyMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLifeSafetyMode) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetLifeSafetyMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLifeSafetyModeParse(readBuffer utils.ReadBuffer) (BACnetLifeSafetyMode, error) {
	val, err := readBuffer.ReadUint16("BACnetLifeSafetyMode", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetLifeSafetyModeByValue(val), nil
}

func (e BACnetLifeSafetyMode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetLifeSafetyMode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLifeSafetyMode) PLC4XEnumName() string {
	switch e {
	case BACnetLifeSafetyMode_OFF:
		return "OFF"
	case BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLifeSafetyMode_ON:
		return "ON"
	case BACnetLifeSafetyMode_DISCONNECTED:
		return "DISCONNECTED"
	case BACnetLifeSafetyMode_ENABLED:
		return "ENABLED"
	case BACnetLifeSafetyMode_DISABLED:
		return "DISABLED"
	case BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED:
		return "AUTOMATIC_RELEASE_DISABLED"
	case BACnetLifeSafetyMode_DEFAULT:
		return "DEFAULT"
	case BACnetLifeSafetyMode_TEST:
		return "TEST"
	case BACnetLifeSafetyMode_MANNED:
		return "MANNED"
	case BACnetLifeSafetyMode_UNMANNED:
		return "UNMANNED"
	case BACnetLifeSafetyMode_ARMED:
		return "ARMED"
	case BACnetLifeSafetyMode_DISARMED:
		return "DISARMED"
	case BACnetLifeSafetyMode_PREARMED:
		return "PREARMED"
	case BACnetLifeSafetyMode_SLOW:
		return "SLOW"
	case BACnetLifeSafetyMode_FAST:
		return "FAST"
	}
	return ""
}

func (e BACnetLifeSafetyMode) String() string {
	return e.PLC4XEnumName()
}
