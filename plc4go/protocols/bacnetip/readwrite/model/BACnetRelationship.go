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

// BACnetRelationship is an enum
type BACnetRelationship uint16

type IBACnetRelationship interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetRelationship_UNKNOWN                  BACnetRelationship = 0
	BACnetRelationship_DEFAULT                  BACnetRelationship = 1
	BACnetRelationship_CONTAINS                 BACnetRelationship = 2
	BACnetRelationship_CONTAINED_BY             BACnetRelationship = 3
	BACnetRelationship_USES                     BACnetRelationship = 4
	BACnetRelationship_USED_BY                  BACnetRelationship = 5
	BACnetRelationship_COMMANDS                 BACnetRelationship = 6
	BACnetRelationship_COMMANDED_BY             BACnetRelationship = 7
	BACnetRelationship_ADJUSTS                  BACnetRelationship = 8
	BACnetRelationship_ADJUSTED_BY              BACnetRelationship = 9
	BACnetRelationship_INGRESS                  BACnetRelationship = 10
	BACnetRelationship_EGRESS                   BACnetRelationship = 11
	BACnetRelationship_SUPPLIES_AIR             BACnetRelationship = 12
	BACnetRelationship_RECEIVES_AIR             BACnetRelationship = 13
	BACnetRelationship_SUPPLIES_HOT_AIR         BACnetRelationship = 14
	BACnetRelationship_RECEIVES_HOT_AIR         BACnetRelationship = 15
	BACnetRelationship_SUPPLIES_COOL_AIR        BACnetRelationship = 16
	BACnetRelationship_RECEIVES_COOL_AIR        BACnetRelationship = 17
	BACnetRelationship_SUPPLIES_POWER           BACnetRelationship = 18
	BACnetRelationship_RECEIVES_POWER           BACnetRelationship = 19
	BACnetRelationship_SUPPLIES_GAS             BACnetRelationship = 20
	BACnetRelationship_RECEIVES_GAS             BACnetRelationship = 21
	BACnetRelationship_SUPPLIES_WATER           BACnetRelationship = 22
	BACnetRelationship_RECEIVES_WATER           BACnetRelationship = 23
	BACnetRelationship_SUPPLIES_HOT_WATER       BACnetRelationship = 24
	BACnetRelationship_RECEIVES_HOT_WATER       BACnetRelationship = 25
	BACnetRelationship_SUPPLIES_COOL_WATER      BACnetRelationship = 26
	BACnetRelationship_RECEIVES_COOL_WATER      BACnetRelationship = 27
	BACnetRelationship_SUPPLIES_STEAM           BACnetRelationship = 28
	BACnetRelationship_RECEIVES_STEAM           BACnetRelationship = 29
	BACnetRelationship_VENDOR_PROPRIETARY_VALUE BACnetRelationship = 0xFFFF
)

var BACnetRelationshipValues []BACnetRelationship

func init() {
	_ = errors.New
	BACnetRelationshipValues = []BACnetRelationship{
		BACnetRelationship_UNKNOWN,
		BACnetRelationship_DEFAULT,
		BACnetRelationship_CONTAINS,
		BACnetRelationship_CONTAINED_BY,
		BACnetRelationship_USES,
		BACnetRelationship_USED_BY,
		BACnetRelationship_COMMANDS,
		BACnetRelationship_COMMANDED_BY,
		BACnetRelationship_ADJUSTS,
		BACnetRelationship_ADJUSTED_BY,
		BACnetRelationship_INGRESS,
		BACnetRelationship_EGRESS,
		BACnetRelationship_SUPPLIES_AIR,
		BACnetRelationship_RECEIVES_AIR,
		BACnetRelationship_SUPPLIES_HOT_AIR,
		BACnetRelationship_RECEIVES_HOT_AIR,
		BACnetRelationship_SUPPLIES_COOL_AIR,
		BACnetRelationship_RECEIVES_COOL_AIR,
		BACnetRelationship_SUPPLIES_POWER,
		BACnetRelationship_RECEIVES_POWER,
		BACnetRelationship_SUPPLIES_GAS,
		BACnetRelationship_RECEIVES_GAS,
		BACnetRelationship_SUPPLIES_WATER,
		BACnetRelationship_RECEIVES_WATER,
		BACnetRelationship_SUPPLIES_HOT_WATER,
		BACnetRelationship_RECEIVES_HOT_WATER,
		BACnetRelationship_SUPPLIES_COOL_WATER,
		BACnetRelationship_RECEIVES_COOL_WATER,
		BACnetRelationship_SUPPLIES_STEAM,
		BACnetRelationship_RECEIVES_STEAM,
		BACnetRelationship_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetRelationshipByValue(value uint16) BACnetRelationship {
	switch value {
	case 0:
		return BACnetRelationship_UNKNOWN
	case 0xFFFF:
		return BACnetRelationship_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetRelationship_DEFAULT
	case 10:
		return BACnetRelationship_INGRESS
	case 11:
		return BACnetRelationship_EGRESS
	case 12:
		return BACnetRelationship_SUPPLIES_AIR
	case 13:
		return BACnetRelationship_RECEIVES_AIR
	case 14:
		return BACnetRelationship_SUPPLIES_HOT_AIR
	case 15:
		return BACnetRelationship_RECEIVES_HOT_AIR
	case 16:
		return BACnetRelationship_SUPPLIES_COOL_AIR
	case 17:
		return BACnetRelationship_RECEIVES_COOL_AIR
	case 18:
		return BACnetRelationship_SUPPLIES_POWER
	case 19:
		return BACnetRelationship_RECEIVES_POWER
	case 2:
		return BACnetRelationship_CONTAINS
	case 20:
		return BACnetRelationship_SUPPLIES_GAS
	case 21:
		return BACnetRelationship_RECEIVES_GAS
	case 22:
		return BACnetRelationship_SUPPLIES_WATER
	case 23:
		return BACnetRelationship_RECEIVES_WATER
	case 24:
		return BACnetRelationship_SUPPLIES_HOT_WATER
	case 25:
		return BACnetRelationship_RECEIVES_HOT_WATER
	case 26:
		return BACnetRelationship_SUPPLIES_COOL_WATER
	case 27:
		return BACnetRelationship_RECEIVES_COOL_WATER
	case 28:
		return BACnetRelationship_SUPPLIES_STEAM
	case 29:
		return BACnetRelationship_RECEIVES_STEAM
	case 3:
		return BACnetRelationship_CONTAINED_BY
	case 4:
		return BACnetRelationship_USES
	case 5:
		return BACnetRelationship_USED_BY
	case 6:
		return BACnetRelationship_COMMANDS
	case 7:
		return BACnetRelationship_COMMANDED_BY
	case 8:
		return BACnetRelationship_ADJUSTS
	case 9:
		return BACnetRelationship_ADJUSTED_BY
	}
	return 0
}

func BACnetRelationshipByName(value string) (enum BACnetRelationship, ok bool) {
	ok = true
	switch value {
	case "UNKNOWN":
		enum = BACnetRelationship_UNKNOWN
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetRelationship_VENDOR_PROPRIETARY_VALUE
	case "DEFAULT":
		enum = BACnetRelationship_DEFAULT
	case "INGRESS":
		enum = BACnetRelationship_INGRESS
	case "EGRESS":
		enum = BACnetRelationship_EGRESS
	case "SUPPLIES_AIR":
		enum = BACnetRelationship_SUPPLIES_AIR
	case "RECEIVES_AIR":
		enum = BACnetRelationship_RECEIVES_AIR
	case "SUPPLIES_HOT_AIR":
		enum = BACnetRelationship_SUPPLIES_HOT_AIR
	case "RECEIVES_HOT_AIR":
		enum = BACnetRelationship_RECEIVES_HOT_AIR
	case "SUPPLIES_COOL_AIR":
		enum = BACnetRelationship_SUPPLIES_COOL_AIR
	case "RECEIVES_COOL_AIR":
		enum = BACnetRelationship_RECEIVES_COOL_AIR
	case "SUPPLIES_POWER":
		enum = BACnetRelationship_SUPPLIES_POWER
	case "RECEIVES_POWER":
		enum = BACnetRelationship_RECEIVES_POWER
	case "CONTAINS":
		enum = BACnetRelationship_CONTAINS
	case "SUPPLIES_GAS":
		enum = BACnetRelationship_SUPPLIES_GAS
	case "RECEIVES_GAS":
		enum = BACnetRelationship_RECEIVES_GAS
	case "SUPPLIES_WATER":
		enum = BACnetRelationship_SUPPLIES_WATER
	case "RECEIVES_WATER":
		enum = BACnetRelationship_RECEIVES_WATER
	case "SUPPLIES_HOT_WATER":
		enum = BACnetRelationship_SUPPLIES_HOT_WATER
	case "RECEIVES_HOT_WATER":
		enum = BACnetRelationship_RECEIVES_HOT_WATER
	case "SUPPLIES_COOL_WATER":
		enum = BACnetRelationship_SUPPLIES_COOL_WATER
	case "RECEIVES_COOL_WATER":
		enum = BACnetRelationship_RECEIVES_COOL_WATER
	case "SUPPLIES_STEAM":
		enum = BACnetRelationship_SUPPLIES_STEAM
	case "RECEIVES_STEAM":
		enum = BACnetRelationship_RECEIVES_STEAM
	case "CONTAINED_BY":
		enum = BACnetRelationship_CONTAINED_BY
	case "USES":
		enum = BACnetRelationship_USES
	case "USED_BY":
		enum = BACnetRelationship_USED_BY
	case "COMMANDS":
		enum = BACnetRelationship_COMMANDS
	case "COMMANDED_BY":
		enum = BACnetRelationship_COMMANDED_BY
	case "ADJUSTS":
		enum = BACnetRelationship_ADJUSTS
	case "ADJUSTED_BY":
		enum = BACnetRelationship_ADJUSTED_BY
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetRelationshipKnows(value uint16) bool {
	for _, typeValue := range BACnetRelationshipValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetRelationship(structType interface{}) BACnetRelationship {
	castFunc := func(typ interface{}) BACnetRelationship {
		if sBACnetRelationship, ok := typ.(BACnetRelationship); ok {
			return sBACnetRelationship
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetRelationship) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetRelationship) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRelationshipParse(readBuffer utils.ReadBuffer) (BACnetRelationship, error) {
	val, err := readBuffer.ReadUint16("BACnetRelationship", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetRelationshipByValue(val), nil
}

func (e BACnetRelationship) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetRelationship", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetRelationship) PLC4XEnumName() string {
	switch e {
	case BACnetRelationship_UNKNOWN:
		return "UNKNOWN"
	case BACnetRelationship_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetRelationship_DEFAULT:
		return "DEFAULT"
	case BACnetRelationship_INGRESS:
		return "INGRESS"
	case BACnetRelationship_EGRESS:
		return "EGRESS"
	case BACnetRelationship_SUPPLIES_AIR:
		return "SUPPLIES_AIR"
	case BACnetRelationship_RECEIVES_AIR:
		return "RECEIVES_AIR"
	case BACnetRelationship_SUPPLIES_HOT_AIR:
		return "SUPPLIES_HOT_AIR"
	case BACnetRelationship_RECEIVES_HOT_AIR:
		return "RECEIVES_HOT_AIR"
	case BACnetRelationship_SUPPLIES_COOL_AIR:
		return "SUPPLIES_COOL_AIR"
	case BACnetRelationship_RECEIVES_COOL_AIR:
		return "RECEIVES_COOL_AIR"
	case BACnetRelationship_SUPPLIES_POWER:
		return "SUPPLIES_POWER"
	case BACnetRelationship_RECEIVES_POWER:
		return "RECEIVES_POWER"
	case BACnetRelationship_CONTAINS:
		return "CONTAINS"
	case BACnetRelationship_SUPPLIES_GAS:
		return "SUPPLIES_GAS"
	case BACnetRelationship_RECEIVES_GAS:
		return "RECEIVES_GAS"
	case BACnetRelationship_SUPPLIES_WATER:
		return "SUPPLIES_WATER"
	case BACnetRelationship_RECEIVES_WATER:
		return "RECEIVES_WATER"
	case BACnetRelationship_SUPPLIES_HOT_WATER:
		return "SUPPLIES_HOT_WATER"
	case BACnetRelationship_RECEIVES_HOT_WATER:
		return "RECEIVES_HOT_WATER"
	case BACnetRelationship_SUPPLIES_COOL_WATER:
		return "SUPPLIES_COOL_WATER"
	case BACnetRelationship_RECEIVES_COOL_WATER:
		return "RECEIVES_COOL_WATER"
	case BACnetRelationship_SUPPLIES_STEAM:
		return "SUPPLIES_STEAM"
	case BACnetRelationship_RECEIVES_STEAM:
		return "RECEIVES_STEAM"
	case BACnetRelationship_CONTAINED_BY:
		return "CONTAINED_BY"
	case BACnetRelationship_USES:
		return "USES"
	case BACnetRelationship_USED_BY:
		return "USED_BY"
	case BACnetRelationship_COMMANDS:
		return "COMMANDS"
	case BACnetRelationship_COMMANDED_BY:
		return "COMMANDED_BY"
	case BACnetRelationship_ADJUSTS:
		return "ADJUSTS"
	case BACnetRelationship_ADJUSTED_BY:
		return "ADJUSTED_BY"
	}
	return ""
}

func (e BACnetRelationship) String() string {
	return e.PLC4XEnumName()
}
