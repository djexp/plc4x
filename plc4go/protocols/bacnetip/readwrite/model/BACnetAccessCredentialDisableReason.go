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

// BACnetAccessCredentialDisableReason is an enum
type BACnetAccessCredentialDisableReason uint16

type IBACnetAccessCredentialDisableReason interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetAccessCredentialDisableReason_DISABLED                    BACnetAccessCredentialDisableReason = 0
	BACnetAccessCredentialDisableReason_DISABLED_NEEDS_PROVISIONING BACnetAccessCredentialDisableReason = 1
	BACnetAccessCredentialDisableReason_DISABLED_UNASSIGNED         BACnetAccessCredentialDisableReason = 2
	BACnetAccessCredentialDisableReason_DISABLED_NOT_YET_ACTIVE     BACnetAccessCredentialDisableReason = 3
	BACnetAccessCredentialDisableReason_DISABLED_EXPIRED            BACnetAccessCredentialDisableReason = 4
	BACnetAccessCredentialDisableReason_DISABLED_LOCKOUT            BACnetAccessCredentialDisableReason = 5
	BACnetAccessCredentialDisableReason_DISABLED_MAX_DAYS           BACnetAccessCredentialDisableReason = 6
	BACnetAccessCredentialDisableReason_DISABLED_MAX_USES           BACnetAccessCredentialDisableReason = 7
	BACnetAccessCredentialDisableReason_DISABLED_INACTIVITY         BACnetAccessCredentialDisableReason = 8
	BACnetAccessCredentialDisableReason_DISABLED_MANUAL             BACnetAccessCredentialDisableReason = 9
	BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE    BACnetAccessCredentialDisableReason = 0xFFFF
)

var BACnetAccessCredentialDisableReasonValues []BACnetAccessCredentialDisableReason

func init() {
	_ = errors.New
	BACnetAccessCredentialDisableReasonValues = []BACnetAccessCredentialDisableReason{
		BACnetAccessCredentialDisableReason_DISABLED,
		BACnetAccessCredentialDisableReason_DISABLED_NEEDS_PROVISIONING,
		BACnetAccessCredentialDisableReason_DISABLED_UNASSIGNED,
		BACnetAccessCredentialDisableReason_DISABLED_NOT_YET_ACTIVE,
		BACnetAccessCredentialDisableReason_DISABLED_EXPIRED,
		BACnetAccessCredentialDisableReason_DISABLED_LOCKOUT,
		BACnetAccessCredentialDisableReason_DISABLED_MAX_DAYS,
		BACnetAccessCredentialDisableReason_DISABLED_MAX_USES,
		BACnetAccessCredentialDisableReason_DISABLED_INACTIVITY,
		BACnetAccessCredentialDisableReason_DISABLED_MANUAL,
		BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAccessCredentialDisableReasonByValue(value uint16) BACnetAccessCredentialDisableReason {
	switch value {
	case 0:
		return BACnetAccessCredentialDisableReason_DISABLED
	case 0xFFFF:
		return BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetAccessCredentialDisableReason_DISABLED_NEEDS_PROVISIONING
	case 2:
		return BACnetAccessCredentialDisableReason_DISABLED_UNASSIGNED
	case 3:
		return BACnetAccessCredentialDisableReason_DISABLED_NOT_YET_ACTIVE
	case 4:
		return BACnetAccessCredentialDisableReason_DISABLED_EXPIRED
	case 5:
		return BACnetAccessCredentialDisableReason_DISABLED_LOCKOUT
	case 6:
		return BACnetAccessCredentialDisableReason_DISABLED_MAX_DAYS
	case 7:
		return BACnetAccessCredentialDisableReason_DISABLED_MAX_USES
	case 8:
		return BACnetAccessCredentialDisableReason_DISABLED_INACTIVITY
	case 9:
		return BACnetAccessCredentialDisableReason_DISABLED_MANUAL
	}
	return 0
}

func BACnetAccessCredentialDisableReasonByName(value string) (enum BACnetAccessCredentialDisableReason, ok bool) {
	ok = true
	switch value {
	case "DISABLED":
		enum = BACnetAccessCredentialDisableReason_DISABLED
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE
	case "DISABLED_NEEDS_PROVISIONING":
		enum = BACnetAccessCredentialDisableReason_DISABLED_NEEDS_PROVISIONING
	case "DISABLED_UNASSIGNED":
		enum = BACnetAccessCredentialDisableReason_DISABLED_UNASSIGNED
	case "DISABLED_NOT_YET_ACTIVE":
		enum = BACnetAccessCredentialDisableReason_DISABLED_NOT_YET_ACTIVE
	case "DISABLED_EXPIRED":
		enum = BACnetAccessCredentialDisableReason_DISABLED_EXPIRED
	case "DISABLED_LOCKOUT":
		enum = BACnetAccessCredentialDisableReason_DISABLED_LOCKOUT
	case "DISABLED_MAX_DAYS":
		enum = BACnetAccessCredentialDisableReason_DISABLED_MAX_DAYS
	case "DISABLED_MAX_USES":
		enum = BACnetAccessCredentialDisableReason_DISABLED_MAX_USES
	case "DISABLED_INACTIVITY":
		enum = BACnetAccessCredentialDisableReason_DISABLED_INACTIVITY
	case "DISABLED_MANUAL":
		enum = BACnetAccessCredentialDisableReason_DISABLED_MANUAL
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetAccessCredentialDisableReasonKnows(value uint16) bool {
	for _, typeValue := range BACnetAccessCredentialDisableReasonValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessCredentialDisableReason(structType interface{}) BACnetAccessCredentialDisableReason {
	castFunc := func(typ interface{}) BACnetAccessCredentialDisableReason {
		if sBACnetAccessCredentialDisableReason, ok := typ.(BACnetAccessCredentialDisableReason); ok {
			return sBACnetAccessCredentialDisableReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessCredentialDisableReason) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetAccessCredentialDisableReason) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccessCredentialDisableReasonParse(readBuffer utils.ReadBuffer) (BACnetAccessCredentialDisableReason, error) {
	val, err := readBuffer.ReadUint16("BACnetAccessCredentialDisableReason", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetAccessCredentialDisableReasonByValue(val), nil
}

func (e BACnetAccessCredentialDisableReason) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetAccessCredentialDisableReason", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessCredentialDisableReason) PLC4XEnumName() string {
	switch e {
	case BACnetAccessCredentialDisableReason_DISABLED:
		return "DISABLED"
	case BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAccessCredentialDisableReason_DISABLED_NEEDS_PROVISIONING:
		return "DISABLED_NEEDS_PROVISIONING"
	case BACnetAccessCredentialDisableReason_DISABLED_UNASSIGNED:
		return "DISABLED_UNASSIGNED"
	case BACnetAccessCredentialDisableReason_DISABLED_NOT_YET_ACTIVE:
		return "DISABLED_NOT_YET_ACTIVE"
	case BACnetAccessCredentialDisableReason_DISABLED_EXPIRED:
		return "DISABLED_EXPIRED"
	case BACnetAccessCredentialDisableReason_DISABLED_LOCKOUT:
		return "DISABLED_LOCKOUT"
	case BACnetAccessCredentialDisableReason_DISABLED_MAX_DAYS:
		return "DISABLED_MAX_DAYS"
	case BACnetAccessCredentialDisableReason_DISABLED_MAX_USES:
		return "DISABLED_MAX_USES"
	case BACnetAccessCredentialDisableReason_DISABLED_INACTIVITY:
		return "DISABLED_INACTIVITY"
	case BACnetAccessCredentialDisableReason_DISABLED_MANUAL:
		return "DISABLED_MANUAL"
	}
	return ""
}

func (e BACnetAccessCredentialDisableReason) String() string {
	return e.PLC4XEnumName()
}
