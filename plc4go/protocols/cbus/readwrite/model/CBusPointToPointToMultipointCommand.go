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

// CBusPointToPointToMultipointCommand is the corresponding interface of CBusPointToPointToMultipointCommand
type CBusPointToPointToMultipointCommand interface {
	utils.LengthAware
	utils.Serializable
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() NetworkRoute
	// GetPeekedApplication returns PeekedApplication (property field)
	GetPeekedApplication() byte
}

// CBusPointToPointToMultipointCommandExactly can be used when we want exactly this type and not a type which fulfills CBusPointToPointToMultipointCommand.
// This is useful for switch cases.
type CBusPointToPointToMultipointCommandExactly interface {
	CBusPointToPointToMultipointCommand
	isCBusPointToPointToMultipointCommand() bool
}

// _CBusPointToPointToMultipointCommand is the data-structure of this message
type _CBusPointToPointToMultipointCommand struct {
	_CBusPointToPointToMultipointCommandChildRequirements
	BridgeAddress     BridgeAddress
	NetworkRoute      NetworkRoute
	PeekedApplication byte

	// Arguments.
	Srchk bool
}

type _CBusPointToPointToMultipointCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type CBusPointToPointToMultipointCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child CBusPointToPointToMultipointCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type CBusPointToPointToMultipointCommandChild interface {
	utils.Serializable
	InitializeParent(parent CBusPointToPointToMultipointCommand, bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte)
	GetParent() *CBusPointToPointToMultipointCommand

	GetTypeName() string
	CBusPointToPointToMultipointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointToMultipointCommand) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_CBusPointToPointToMultipointCommand) GetNetworkRoute() NetworkRoute {
	return m.NetworkRoute
}

func (m *_CBusPointToPointToMultipointCommand) GetPeekedApplication() byte {
	return m.PeekedApplication
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointToMultipointCommand factory function for _CBusPointToPointToMultipointCommand
func NewCBusPointToPointToMultipointCommand(bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte, srchk bool) *_CBusPointToPointToMultipointCommand {
	return &_CBusPointToPointToMultipointCommand{BridgeAddress: bridgeAddress, NetworkRoute: networkRoute, PeekedApplication: peekedApplication, Srchk: srchk}
}

// Deprecated: use the interface for direct cast
func CastCBusPointToPointToMultipointCommand(structType interface{}) CBusPointToPointToMultipointCommand {
	if casted, ok := structType.(CBusPointToPointToMultipointCommand); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointToMultipointCommand); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointToMultipointCommand) GetTypeName() string {
	return "CBusPointToPointToMultipointCommand"
}

func (m *_CBusPointToPointToMultipointCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (bridgeAddress)
	lengthInBits += m.BridgeAddress.GetLengthInBits()

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits()

	return lengthInBits
}

func (m *_CBusPointToPointToMultipointCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusPointToPointToMultipointCommandParse(readBuffer utils.ReadBuffer, srchk bool) (CBusPointToPointToMultipointCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointToMultipointCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointToMultipointCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bridgeAddress)
	if pullErr := readBuffer.PullContext("bridgeAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bridgeAddress")
	}
	_bridgeAddress, _bridgeAddressErr := BridgeAddressParse(readBuffer)
	if _bridgeAddressErr != nil {
		return nil, errors.Wrap(_bridgeAddressErr, "Error parsing 'bridgeAddress' field of CBusPointToPointToMultipointCommand")
	}
	bridgeAddress := _bridgeAddress.(BridgeAddress)
	if closeErr := readBuffer.CloseContext("bridgeAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bridgeAddress")
	}

	// Simple Field (networkRoute)
	if pullErr := readBuffer.PullContext("networkRoute"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkRoute")
	}
	_networkRoute, _networkRouteErr := NetworkRouteParse(readBuffer)
	if _networkRouteErr != nil {
		return nil, errors.Wrap(_networkRouteErr, "Error parsing 'networkRoute' field of CBusPointToPointToMultipointCommand")
	}
	networkRoute := _networkRoute.(NetworkRoute)
	if closeErr := readBuffer.CloseContext("networkRoute"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkRoute")
	}

	// Peek Field (peekedApplication)
	currentPos = positionAware.GetPos()
	peekedApplication, _err := readBuffer.ReadByte("peekedApplication")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekedApplication' field of CBusPointToPointToMultipointCommand")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CBusPointToPointToMultipointCommandChildSerializeRequirement interface {
		CBusPointToPointToMultipointCommand
		InitializeParent(CBusPointToPointToMultipointCommand, BridgeAddress, NetworkRoute, byte)
		GetParent() CBusPointToPointToMultipointCommand
	}
	var _childTemp interface{}
	var _child CBusPointToPointToMultipointCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedApplication == 0xFF: // CBusCommandPointToPointToMultiPointStatus
		_childTemp, typeSwitchError = CBusCommandPointToPointToMultiPointStatusParse(readBuffer, srchk)
	case true: // CBusCommandPointToPointToMultiPointNormal
		_childTemp, typeSwitchError = CBusCommandPointToPointToMultiPointNormalParse(readBuffer, srchk)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedApplication=%v]", peekedApplication)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CBusPointToPointToMultipointCommand")
	}
	_child = _childTemp.(CBusPointToPointToMultipointCommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CBusPointToPointToMultipointCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointToMultipointCommand")
	}

	// Finish initializing
	_child.InitializeParent(_child, bridgeAddress, networkRoute, peekedApplication)
	return _child, nil
}

func (pm *_CBusPointToPointToMultipointCommand) SerializeParent(writeBuffer utils.WriteBuffer, child CBusPointToPointToMultipointCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CBusPointToPointToMultipointCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusPointToPointToMultipointCommand")
	}

	// Simple Field (bridgeAddress)
	if pushErr := writeBuffer.PushContext("bridgeAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for bridgeAddress")
	}
	_bridgeAddressErr := writeBuffer.WriteSerializable(m.GetBridgeAddress())
	if popErr := writeBuffer.PopContext("bridgeAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for bridgeAddress")
	}
	if _bridgeAddressErr != nil {
		return errors.Wrap(_bridgeAddressErr, "Error serializing 'bridgeAddress' field")
	}

	// Simple Field (networkRoute)
	if pushErr := writeBuffer.PushContext("networkRoute"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for networkRoute")
	}
	_networkRouteErr := writeBuffer.WriteSerializable(m.GetNetworkRoute())
	if popErr := writeBuffer.PopContext("networkRoute"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for networkRoute")
	}
	if _networkRouteErr != nil {
		return errors.Wrap(_networkRouteErr, "Error serializing 'networkRoute' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusPointToPointToMultipointCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusPointToPointToMultipointCommand")
	}
	return nil
}

func (m *_CBusPointToPointToMultipointCommand) isCBusPointToPointToMultipointCommand() bool {
	return true
}

func (m *_CBusPointToPointToMultipointCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
