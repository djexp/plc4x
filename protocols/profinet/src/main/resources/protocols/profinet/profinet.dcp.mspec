//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

[type 'EthernetFrame'
    [simple MacAddress 'destination' ]
    [simple MacAddress 'source'      ]
    [simple uint 16    'ethernetType']
    [simple ProfinetFrame 'payload'  ]
]

[type 'ProfinetFrame'
    [enum FrameType 'frameType'                ]
    [simple ProfinetData 'frame'  ['frameType']]
    [padding uint 8 'alignment' '0x00' '40 - frame.lengthInBytes']
]

[discriminatedType 'ProfinetData' [FrameType 'frameType']
    [typeSwitch 'frameType'
        ['FrameType.IDENTIFY_MULTICAST_REQUEST' DcpIdentRequestPDU
            [enum DCPServiceID 'serviceId'                 ]
            [enum DCPServiceType 'serviceType'             ]
            [simple uint 32 'xid'                          ]
            [simple uint 16 'responseDelay'                ]
            [simple uint 16 'dcpDataLength'                ]
            [array DCPBlock 'blocks' length 'dcpDataLength']
        ]
        ['FrameType.IDENTIFY_RESPONSE' DcpIdentResponsePDU
            [enum DCPServiceID 'serviceId'                 ]
            [enum DCPServiceType 'serviceType'             ]
            [simple uint 32 'xid'                          ]
            [simple uint 16 'responseDelay'                ]
            [simple uint 16 'dcpDataLength'                ]
            [array DCPBlock 'blocks' length 'dcpDataLength']
        ]
    ]
]

[discriminatedType 'DCPBlock'
    [enum DCPBlockOption 'blockType']
    [typeSwitch 'blockType'
        ['DCPBlockOption.IP' IP
            [enum   IpSubOption 'subOption']
            [simple uint 16 'length'   ]
            [simple uint 16 'info'     ]
            [simple IPv4Address 'ipAddress'      ]
            [simple IPv4Address 'subnetMask'     ]
            [simple IPv4Address 'standardGateway']
        ]
        ['DCPBlockOption.DEVICE_PROPERTIES' DeviceProperties
            [enum   DevicePropertiesSubOption 'subOption' ]
            [simple DCPDeviceProperties       'properties' ['subOption']]
        ]
        ['DCPBlockOption.ALL_SELECTOR' AllSelector
            [const uint 8  'subOption'  '0xFF' ]
            [const uint 16 'length'     '0x00' ]
        ]
    ]
]

[discriminatedType 'DCPDeviceProperties' [DevicePropertiesSubOption 'subOptionType']
    [simple uint 16 'length' ]
    [simple uint 16 'info'   ]
    [typeSwitch 'subOptionType'
        ['DevicePropertiesSubOption.STATION_TYPE' StationType [uint 16 'length']
            [simple Text 'vendorNameForDevice' ['length - 2']]
        ]
        ['DevicePropertiesSubOption.STATION_NAME' StationName [uint 16 'length']
            [simple Text 'name' ['length - 2']]
        ]
        ['DevicePropertiesSubOption.DEVICE_ID' DeviceId
            [simple uint 16 'vendorId'         ]
            [simple uint 16 'deviceId'         ]
        ]
        ['DevicePropertiesSubOption.DEVICE_ROLE' DeviceRole
            [simple uint 8   'role'       ]
            [reserved uint 8 '0x00'       ]
        ]
        ['DevicePropertiesSubOption.DEVICE_OPTIONS' DeviceOptions
            [enum DCPBlockOption            'blockType'         ]
            [enum DevicePropertiesSubOption 'subOption'         ]
        ]
        ['DevicePropertiesSubOption.DEVICE_INSTANCE' DeviceInstance
            [simple uint 8 'instanceHigh'        ]
            [simple uint 8 'instanceLow'         ]
        ]
    ]
    // Padding using 0..1 old format
    //[padding uint 8                 'pad' '0x00' '(length % 2 == 0 ? 0 : 1)']
    // Padding using 0..N new format
    [padding uint 8                 'pad' '0x00' 'length % 2']
]

[enum uint 8 'DCPBlockOption'
    ['0x01' IP                          ]
    ['0x02' DEVICE_PROPERTIES           ]
    ['0x03' DHCP                        ]
    // TODO implement DCPBlockOption case
    ['0x05' CONTROL                     ]
    ['0x06' DEVICE_INACTIVE             ]
    ['0xFF' ALL_SELECTOR                ]
]
// TODO implement DCPBlockOption case
[enum uint 8 'ControlOption'
    ['0x00' RESERVED                    ]
    ['0x01' START_TRANSACTION           ]
    ['0x02' END_TRANSACTION             ]
    ['0x03' FLASH_LED                   ]
    ['0x04' RESPONSE                    ]
    //Reset Factory Settings (0x05)
    ['0x05' FACTORY_RESET               ]
    ['0x06' RESET_TO_FACTORY            ]
]
// TODO implement DCPBlockOption case
[enum uint 8 'DeviceInactiveOption'
    ['0x00' RESERVED                    ]
    ['0x01' DEVICE_INITIATIVE           ]
]

// Not implemented yet! Vars from https://github.com/dark-lbp/isf/blob/master/icssploit/protocols/pn_dcp.py
[enum uint 8 'DHCPOptions'
    ['0x0c' HOST_NAME                   ]
    ['0x2b' VENDOR_SPECIFIC             ]
    ['0x36' SERVER_IDENTIFIER           ]
    ['0x37' PARAMETER_REQUEST_LIST      ]
    ['0x3c' CLASS_IDENTIFIER            ]
    ['0x3d' DHCP_CLIENT_IDENTIFIER      ]
    ['0x51' FQDN_FULLY_QUALIFIED_DOMAIN_NAME       ]
    ['0x61' UUID_GUID_BASED_CLIENT      ]
    ['0xFF' CONTROL_DHCP_FOR_ADDRESS_RESOLUTION    ]
]

[enum uint 8 'DevicePropertiesSubOption'
    ['0x01' STATION_TYPE                ]
    ['0x02' STATION_NAME                ]
    ['0x03' DEVICE_ID                   ]
    ['0x04' DEVICE_ROLE                 ]
    ['0x05' DEVICE_OPTIONS              ]
    ['0x06' ALIAS_NAME                  ]
    ['0x07' DEVICE_INSTANCE             ]
    ['0x08' OEM_DEVICE_ID               ]
]

[enum uint 8 'IpSubOption'
    ['0x00' RESERVED                    ]
    ['0x01' MAC_ADDRESS                 ]
    ['0x02' IP_PARAMETER                ]
]

[enum uint 8 'DCPServiceID'
    ['0x03' GET                         ]
    ['0x04' SET                         ]
    ['0x05' IDENTIFY                    ]
    ['0x06' HELLO                       ]
    ['0xFEFE' IDENTIFY_RESPONSE         ]
]

[enum uint 8 'DCPServiceType'
    ['0x00' REQUEST                     ]
    ['0x01' RESPONSE_SUCCESS            ]
]

[enum uint 16 'FrameType'
    ['0xFEFE' IDENTIFY_MULTICAST_REQUEST]
    ['0xFEFF' IDENTIFY_RESPONSE         ]
]

[type 'MacAddress'
    [simple     uint        8   'octet1'            ]
    [simple     uint        8   'octet2'            ]
    [simple     uint        8   'octet3'            ]
    [simple     uint        8   'octet4'            ]
    [simple     uint        8   'octet5'            ]
    [simple     uint        8   'octet6'            ]
]

[type 'IPv4Address'
    [simple     uint        8   'octet1'            ]
    [simple     uint        8   'octet2'            ]
    [simple     uint        8   'octet3'            ]
    [simple     uint        8   'octet4'            ]
]

[type 'Text' [uint 16  'len']
    // fixme _type.encoding doesn't work for readString call, it ends up with &quot; sequence in generated Java
    [manual string 'UTF-8' 'text'
        'STATIC_CALL("org.apache.plc4x.java.profinet.dcp.DCPUtil.readString", io, len)'
        'STATIC_CALL("org.apache.plc4x.java.profinet.dcp.DCPUtil.writeString", io, _value, _type.encoding)'
        'STATIC_CALL("org.apache.plc4x.java.profinet.dcp.DCPUtil.length", text)'
    ]
]

[enum uint 16 'TypeLAN'
    ['0x8100' VLAN             ]
    ['0x8892' PN_DCP           ]
    ['0x88CC' LLDP             ]
]