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

// BACnetLiftCarDoorCommand is an enum
type BACnetLiftCarDoorCommand uint8

type IBACnetLiftCarDoorCommand interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetLiftCarDoorCommand_NONE  BACnetLiftCarDoorCommand = 0
	BACnetLiftCarDoorCommand_OPEN  BACnetLiftCarDoorCommand = 1
	BACnetLiftCarDoorCommand_CLOSE BACnetLiftCarDoorCommand = 2
)

var BACnetLiftCarDoorCommandValues []BACnetLiftCarDoorCommand

func init() {
	_ = errors.New
	BACnetLiftCarDoorCommandValues = []BACnetLiftCarDoorCommand{
		BACnetLiftCarDoorCommand_NONE,
		BACnetLiftCarDoorCommand_OPEN,
		BACnetLiftCarDoorCommand_CLOSE,
	}
}

func BACnetLiftCarDoorCommandByValue(value uint8) BACnetLiftCarDoorCommand {
	switch value {
	case 0:
		return BACnetLiftCarDoorCommand_NONE
	case 1:
		return BACnetLiftCarDoorCommand_OPEN
	case 2:
		return BACnetLiftCarDoorCommand_CLOSE
	}
	return 0
}

func BACnetLiftCarDoorCommandByName(value string) (enum BACnetLiftCarDoorCommand, ok bool) {
	ok = true
	switch value {
	case "NONE":
		enum = BACnetLiftCarDoorCommand_NONE
	case "OPEN":
		enum = BACnetLiftCarDoorCommand_OPEN
	case "CLOSE":
		enum = BACnetLiftCarDoorCommand_CLOSE
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetLiftCarDoorCommandKnows(value uint8) bool {
	for _, typeValue := range BACnetLiftCarDoorCommandValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLiftCarDoorCommand(structType interface{}) BACnetLiftCarDoorCommand {
	castFunc := func(typ interface{}) BACnetLiftCarDoorCommand {
		if sBACnetLiftCarDoorCommand, ok := typ.(BACnetLiftCarDoorCommand); ok {
			return sBACnetLiftCarDoorCommand
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLiftCarDoorCommand) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetLiftCarDoorCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLiftCarDoorCommandParse(readBuffer utils.ReadBuffer) (BACnetLiftCarDoorCommand, error) {
	val, err := readBuffer.ReadUint8("BACnetLiftCarDoorCommand", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetLiftCarDoorCommandByValue(val), nil
}

func (e BACnetLiftCarDoorCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetLiftCarDoorCommand", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLiftCarDoorCommand) PLC4XEnumName() string {
	switch e {
	case BACnetLiftCarDoorCommand_NONE:
		return "NONE"
	case BACnetLiftCarDoorCommand_OPEN:
		return "OPEN"
	case BACnetLiftCarDoorCommand_CLOSE:
		return "CLOSE"
	}
	return ""
}

func (e BACnetLiftCarDoorCommand) String() string {
	return e.PLC4XEnumName()
}
