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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type MQTT_PropertyType uint8

type IMQTT_PropertyType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const(
	MQTT_PropertyType_PAYLOAD_FORMAT_INDICATOR MQTT_PropertyType = 0x01
	MQTT_PropertyType_MESSAGE_EXPIRY_INTERVAL MQTT_PropertyType = 0x02
	MQTT_PropertyType_CONTENT_TYPE MQTT_PropertyType = 0x03
	MQTT_PropertyType_RESPONSE_TOPIC MQTT_PropertyType = 0x08
	MQTT_PropertyType_CORRELATION_DATA MQTT_PropertyType = 0x09
	MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER MQTT_PropertyType = 0x0B
	MQTT_PropertyType_SESSION_EXPIRY_INTERVAL MQTT_PropertyType = 0x11
	MQTT_PropertyType_ASSIGNED_CLIENT_IDENTIFIER MQTT_PropertyType = 0x12
	MQTT_PropertyType_SERVER_KEEP_ALIVE MQTT_PropertyType = 0x13
	MQTT_PropertyType_AUTHENTICATION_METHOD MQTT_PropertyType = 0x15
	MQTT_PropertyType_AUTHENTICATION_DATA MQTT_PropertyType = 0x16
	MQTT_PropertyType_REQUEST_PROBLEM_INFORMATION MQTT_PropertyType = 0x17
	MQTT_PropertyType_WILL_DELAY_INTERVAL MQTT_PropertyType = 0x18
	MQTT_PropertyType_REQUEST_RESPONSE_INFORMATION MQTT_PropertyType = 0x19
	MQTT_PropertyType_RESPONSE_INFORMATION MQTT_PropertyType = 0x1A
	MQTT_PropertyType_SERVER_REFERENCE MQTT_PropertyType = 0x1C
	MQTT_PropertyType_REASON_STRING MQTT_PropertyType = 0x1F
	MQTT_PropertyType_RECEIVE_MAXIMUM MQTT_PropertyType = 0x21
	MQTT_PropertyType_TOPIC_ALIAS_MAXIMUM MQTT_PropertyType = 0x22
	MQTT_PropertyType_TOPIC_ALIAS MQTT_PropertyType = 0x23
	MQTT_PropertyType_MAXIMUM_QOS MQTT_PropertyType = 0x24
	MQTT_PropertyType_RETAIN_AVAILABLE MQTT_PropertyType = 0x25
	MQTT_PropertyType_USER_PROPERTY MQTT_PropertyType = 0x26
	MQTT_PropertyType_MAXIMUM_PACKET_SIZE MQTT_PropertyType = 0x27
	MQTT_PropertyType_WILDCARD_SUBSCRIPTION_AVAILABLE MQTT_PropertyType = 0x28
	MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER_AVAILABLE MQTT_PropertyType = 0x29
	MQTT_PropertyType_SHARED_SUBSCRIPTION_AVAILABLE MQTT_PropertyType = 0x2A
)

var MQTT_PropertyTypeValues []MQTT_PropertyType

func init() {
	_ = errors.New
	MQTT_PropertyTypeValues = []MQTT_PropertyType {
		MQTT_PropertyType_PAYLOAD_FORMAT_INDICATOR,
		MQTT_PropertyType_MESSAGE_EXPIRY_INTERVAL,
		MQTT_PropertyType_CONTENT_TYPE,
		MQTT_PropertyType_RESPONSE_TOPIC,
		MQTT_PropertyType_CORRELATION_DATA,
		MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER,
		MQTT_PropertyType_SESSION_EXPIRY_INTERVAL,
		MQTT_PropertyType_ASSIGNED_CLIENT_IDENTIFIER,
		MQTT_PropertyType_SERVER_KEEP_ALIVE,
		MQTT_PropertyType_AUTHENTICATION_METHOD,
		MQTT_PropertyType_AUTHENTICATION_DATA,
		MQTT_PropertyType_REQUEST_PROBLEM_INFORMATION,
		MQTT_PropertyType_WILL_DELAY_INTERVAL,
		MQTT_PropertyType_REQUEST_RESPONSE_INFORMATION,
		MQTT_PropertyType_RESPONSE_INFORMATION,
		MQTT_PropertyType_SERVER_REFERENCE,
		MQTT_PropertyType_REASON_STRING,
		MQTT_PropertyType_RECEIVE_MAXIMUM,
		MQTT_PropertyType_TOPIC_ALIAS_MAXIMUM,
		MQTT_PropertyType_TOPIC_ALIAS,
		MQTT_PropertyType_MAXIMUM_QOS,
		MQTT_PropertyType_RETAIN_AVAILABLE,
		MQTT_PropertyType_USER_PROPERTY,
		MQTT_PropertyType_MAXIMUM_PACKET_SIZE,
		MQTT_PropertyType_WILDCARD_SUBSCRIPTION_AVAILABLE,
		MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER_AVAILABLE,
		MQTT_PropertyType_SHARED_SUBSCRIPTION_AVAILABLE,
	}
}

func MQTT_PropertyTypeByValue(value uint8) MQTT_PropertyType {
	switch value {
		case 0x01:
			return MQTT_PropertyType_PAYLOAD_FORMAT_INDICATOR
		case 0x02:
			return MQTT_PropertyType_MESSAGE_EXPIRY_INTERVAL
		case 0x03:
			return MQTT_PropertyType_CONTENT_TYPE
		case 0x08:
			return MQTT_PropertyType_RESPONSE_TOPIC
		case 0x09:
			return MQTT_PropertyType_CORRELATION_DATA
		case 0x0B:
			return MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER
		case 0x11:
			return MQTT_PropertyType_SESSION_EXPIRY_INTERVAL
		case 0x12:
			return MQTT_PropertyType_ASSIGNED_CLIENT_IDENTIFIER
		case 0x13:
			return MQTT_PropertyType_SERVER_KEEP_ALIVE
		case 0x15:
			return MQTT_PropertyType_AUTHENTICATION_METHOD
		case 0x16:
			return MQTT_PropertyType_AUTHENTICATION_DATA
		case 0x17:
			return MQTT_PropertyType_REQUEST_PROBLEM_INFORMATION
		case 0x18:
			return MQTT_PropertyType_WILL_DELAY_INTERVAL
		case 0x19:
			return MQTT_PropertyType_REQUEST_RESPONSE_INFORMATION
		case 0x1A:
			return MQTT_PropertyType_RESPONSE_INFORMATION
		case 0x1C:
			return MQTT_PropertyType_SERVER_REFERENCE
		case 0x1F:
			return MQTT_PropertyType_REASON_STRING
		case 0x21:
			return MQTT_PropertyType_RECEIVE_MAXIMUM
		case 0x22:
			return MQTT_PropertyType_TOPIC_ALIAS_MAXIMUM
		case 0x23:
			return MQTT_PropertyType_TOPIC_ALIAS
		case 0x24:
			return MQTT_PropertyType_MAXIMUM_QOS
		case 0x25:
			return MQTT_PropertyType_RETAIN_AVAILABLE
		case 0x26:
			return MQTT_PropertyType_USER_PROPERTY
		case 0x27:
			return MQTT_PropertyType_MAXIMUM_PACKET_SIZE
		case 0x28:
			return MQTT_PropertyType_WILDCARD_SUBSCRIPTION_AVAILABLE
		case 0x29:
			return MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER_AVAILABLE
		case 0x2A:
			return MQTT_PropertyType_SHARED_SUBSCRIPTION_AVAILABLE
	}
	return 0
}

func MQTT_PropertyTypeByName(value string) MQTT_PropertyType {
	switch value {
	case "PAYLOAD_FORMAT_INDICATOR":
		return MQTT_PropertyType_PAYLOAD_FORMAT_INDICATOR
	case "MESSAGE_EXPIRY_INTERVAL":
		return MQTT_PropertyType_MESSAGE_EXPIRY_INTERVAL
	case "CONTENT_TYPE":
		return MQTT_PropertyType_CONTENT_TYPE
	case "RESPONSE_TOPIC":
		return MQTT_PropertyType_RESPONSE_TOPIC
	case "CORRELATION_DATA":
		return MQTT_PropertyType_CORRELATION_DATA
	case "SUBSCRIPTION_IDENTIFIER":
		return MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER
	case "SESSION_EXPIRY_INTERVAL":
		return MQTT_PropertyType_SESSION_EXPIRY_INTERVAL
	case "ASSIGNED_CLIENT_IDENTIFIER":
		return MQTT_PropertyType_ASSIGNED_CLIENT_IDENTIFIER
	case "SERVER_KEEP_ALIVE":
		return MQTT_PropertyType_SERVER_KEEP_ALIVE
	case "AUTHENTICATION_METHOD":
		return MQTT_PropertyType_AUTHENTICATION_METHOD
	case "AUTHENTICATION_DATA":
		return MQTT_PropertyType_AUTHENTICATION_DATA
	case "REQUEST_PROBLEM_INFORMATION":
		return MQTT_PropertyType_REQUEST_PROBLEM_INFORMATION
	case "WILL_DELAY_INTERVAL":
		return MQTT_PropertyType_WILL_DELAY_INTERVAL
	case "REQUEST_RESPONSE_INFORMATION":
		return MQTT_PropertyType_REQUEST_RESPONSE_INFORMATION
	case "RESPONSE_INFORMATION":
		return MQTT_PropertyType_RESPONSE_INFORMATION
	case "SERVER_REFERENCE":
		return MQTT_PropertyType_SERVER_REFERENCE
	case "REASON_STRING":
		return MQTT_PropertyType_REASON_STRING
	case "RECEIVE_MAXIMUM":
		return MQTT_PropertyType_RECEIVE_MAXIMUM
	case "TOPIC_ALIAS_MAXIMUM":
		return MQTT_PropertyType_TOPIC_ALIAS_MAXIMUM
	case "TOPIC_ALIAS":
		return MQTT_PropertyType_TOPIC_ALIAS
	case "MAXIMUM_QOS":
		return MQTT_PropertyType_MAXIMUM_QOS
	case "RETAIN_AVAILABLE":
		return MQTT_PropertyType_RETAIN_AVAILABLE
	case "USER_PROPERTY":
		return MQTT_PropertyType_USER_PROPERTY
	case "MAXIMUM_PACKET_SIZE":
		return MQTT_PropertyType_MAXIMUM_PACKET_SIZE
	case "WILDCARD_SUBSCRIPTION_AVAILABLE":
		return MQTT_PropertyType_WILDCARD_SUBSCRIPTION_AVAILABLE
	case "SUBSCRIPTION_IDENTIFIER_AVAILABLE":
		return MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER_AVAILABLE
	case "SHARED_SUBSCRIPTION_AVAILABLE":
		return MQTT_PropertyType_SHARED_SUBSCRIPTION_AVAILABLE
	}
	return 0
}

func CastMQTT_PropertyType(structType interface{}) MQTT_PropertyType {
	castFunc := func(typ interface{}) MQTT_PropertyType {
		if sMQTT_PropertyType, ok := typ.(MQTT_PropertyType); ok {
			return sMQTT_PropertyType
		}
		return 0
	}
	return castFunc(structType)
}

func (m MQTT_PropertyType) LengthInBits() uint16 {
	return 8
}

func (m MQTT_PropertyType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MQTT_PropertyTypeParse(readBuffer utils.ReadBuffer) (MQTT_PropertyType, error) {
	val, err := readBuffer.ReadUint8("MQTT_PropertyType", 8)
	if err != nil {
		return 0, nil
	}
	return MQTT_PropertyTypeByValue(val), nil
}

func (e MQTT_PropertyType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("MQTT_PropertyType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e MQTT_PropertyType) name() string {
	switch e {
	case MQTT_PropertyType_PAYLOAD_FORMAT_INDICATOR:
		return "PAYLOAD_FORMAT_INDICATOR"
	case MQTT_PropertyType_MESSAGE_EXPIRY_INTERVAL:
		return "MESSAGE_EXPIRY_INTERVAL"
	case MQTT_PropertyType_CONTENT_TYPE:
		return "CONTENT_TYPE"
	case MQTT_PropertyType_RESPONSE_TOPIC:
		return "RESPONSE_TOPIC"
	case MQTT_PropertyType_CORRELATION_DATA:
		return "CORRELATION_DATA"
	case MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER:
		return "SUBSCRIPTION_IDENTIFIER"
	case MQTT_PropertyType_SESSION_EXPIRY_INTERVAL:
		return "SESSION_EXPIRY_INTERVAL"
	case MQTT_PropertyType_ASSIGNED_CLIENT_IDENTIFIER:
		return "ASSIGNED_CLIENT_IDENTIFIER"
	case MQTT_PropertyType_SERVER_KEEP_ALIVE:
		return "SERVER_KEEP_ALIVE"
	case MQTT_PropertyType_AUTHENTICATION_METHOD:
		return "AUTHENTICATION_METHOD"
	case MQTT_PropertyType_AUTHENTICATION_DATA:
		return "AUTHENTICATION_DATA"
	case MQTT_PropertyType_REQUEST_PROBLEM_INFORMATION:
		return "REQUEST_PROBLEM_INFORMATION"
	case MQTT_PropertyType_WILL_DELAY_INTERVAL:
		return "WILL_DELAY_INTERVAL"
	case MQTT_PropertyType_REQUEST_RESPONSE_INFORMATION:
		return "REQUEST_RESPONSE_INFORMATION"
	case MQTT_PropertyType_RESPONSE_INFORMATION:
		return "RESPONSE_INFORMATION"
	case MQTT_PropertyType_SERVER_REFERENCE:
		return "SERVER_REFERENCE"
	case MQTT_PropertyType_REASON_STRING:
		return "REASON_STRING"
	case MQTT_PropertyType_RECEIVE_MAXIMUM:
		return "RECEIVE_MAXIMUM"
	case MQTT_PropertyType_TOPIC_ALIAS_MAXIMUM:
		return "TOPIC_ALIAS_MAXIMUM"
	case MQTT_PropertyType_TOPIC_ALIAS:
		return "TOPIC_ALIAS"
	case MQTT_PropertyType_MAXIMUM_QOS:
		return "MAXIMUM_QOS"
	case MQTT_PropertyType_RETAIN_AVAILABLE:
		return "RETAIN_AVAILABLE"
	case MQTT_PropertyType_USER_PROPERTY:
		return "USER_PROPERTY"
	case MQTT_PropertyType_MAXIMUM_PACKET_SIZE:
		return "MAXIMUM_PACKET_SIZE"
	case MQTT_PropertyType_WILDCARD_SUBSCRIPTION_AVAILABLE:
		return "WILDCARD_SUBSCRIPTION_AVAILABLE"
	case MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER_AVAILABLE:
		return "SUBSCRIPTION_IDENTIFIER_AVAILABLE"
	case MQTT_PropertyType_SHARED_SUBSCRIPTION_AVAILABLE:
		return "SHARED_SUBSCRIPTION_AVAILABLE"
	}
	return ""
}

func (e MQTT_PropertyType) String() string {
	return e.name()
}

