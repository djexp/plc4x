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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "s7_data_alarm_message.h"

// Code generated by code-generation. DO NOT EDIT.

// Array of discriminator values that match the enum type constants.
// (The order is identical to the enum constants so we can use the
// enum constant to directly access a given types discriminator values)
const plc4c_s7_read_write_s7_data_alarm_message_discriminator plc4c_s7_read_write_s7_data_alarm_message_discriminators[] = {
  {/* plc4c_s7_read_write_s7_message_object_request */
   .cpuFunctionType = DefaultHexadecimalLiteral{hexString=0x04} },
  {/* plc4c_s7_read_write_s7_message_object_response */
   .cpuFunctionType = DefaultHexadecimalLiteral{hexString=0x08} }

};

// Function returning the discriminator values for a given type constant.
plc4c_s7_read_write_s7_data_alarm_message_discriminator plc4c_s7_read_write_s7_data_alarm_message_get_discriminator(plc4c_s7_read_write_s7_data_alarm_message_type type) {
  return plc4c_s7_read_write_s7_data_alarm_message_discriminators[type];
}

// Create an empty NULL-struct
static const plc4c_s7_read_write_s7_data_alarm_message plc4c_s7_read_write_s7_data_alarm_message_null_const;

plc4c_s7_read_write_s7_data_alarm_message plc4c_s7_read_write_s7_data_alarm_message_null() {
  return plc4c_s7_read_write_s7_data_alarm_message_null_const;
}


// Constant values.
static const uint8_t PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_FUNCTION_ID_const = 0x00;
uint8_t PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_FUNCTION_ID() {
  return PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_FUNCTION_ID_const;
}
static const uint8_t PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_NUMBER_MESSAGE_OBJ_const = 0x01;
uint8_t PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_NUMBER_MESSAGE_OBJ() {
  return PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_NUMBER_MESSAGE_OBJ_const;
}
static const uint8_t PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_VARIABLE_SPEC_const = 0x12;
uint8_t PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_VARIABLE_SPEC() {
  return PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_VARIABLE_SPEC_const;
}
static const uint8_t PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_LENGTH_const = 0x08;
uint8_t PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_LENGTH() {
  return PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_LENGTH_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_s7_data_alarm_message_parse(plc4c_spi_read_buffer* readBuffer, uint8_t cpuFunctionType, plc4c_s7_read_write_s7_data_alarm_message** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_s7_data_alarm_message));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Const Field (functionId)
  uint8_t functionId = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &functionId);
  if(_res != OK) {
    return _res;
  }
  if(functionId != PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_FUNCTION_ID()) {
    return PARSE_ERROR;
    // throw new ParseException("Expected constant value " + PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_FUNCTION_ID + " but got " + functionId);
  }

  // Const Field (numberMessageObj)
  uint8_t numberMessageObj = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &numberMessageObj);
  if(_res != OK) {
    return _res;
  }
  if(numberMessageObj != PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_NUMBER_MESSAGE_OBJ()) {
    return PARSE_ERROR;
    // throw new ParseException("Expected constant value " + PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_NUMBER_MESSAGE_OBJ + " but got " + numberMessageObj);
  }

  // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
  { /* S7MessageObjectRequest */
    (*_message)->_type = plc4c_s7_read_write_s7_data_alarm_message_type_plc4c_s7_read_write_s7_message_object_request;
                    
    // Const Field (variableSpec)
    uint8_t variableSpec = 0;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &variableSpec);
    if(_res != OK) {
      return _res;
    }
    if(variableSpec != PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_VARIABLE_SPEC()) {
      return PARSE_ERROR;
      // throw new ParseException("Expected constant value " + PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_VARIABLE_SPEC + " but got " + variableSpec);
    }


                    
    // Const Field (length)
    uint8_t length = 0;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &length);
    if(_res != OK) {
      return _res;
    }
    if(length != PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_LENGTH()) {
      return PARSE_ERROR;
      // throw new ParseException("Expected constant value " + PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_LENGTH + " but got " + length);
    }


                    
    // Simple Field (syntaxId)
    plc4c_s7_read_write_syntax_id_type* syntaxId;
    _res = plc4c_s7_read_write_syntax_id_type_parse(readBuffer, (void*) &syntaxId);
    if(_res != OK) {
      return _res;
    }
    (*_message)->s7_message_object_request_syntax_id = *syntaxId;


                    
    // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
    {
      uint8_t _reserved = 0;
      _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &_reserved);
      if(_res != OK) {
        return _res;
      }
      if(_reserved != 0x00) {
        printf("Expected constant value '%d' but got '%d' for reserved field.", 0x00, _reserved);
      }
    }


                    
    // Simple Field (queryType)
    plc4c_s7_read_write_query_type* queryType;
    _res = plc4c_s7_read_write_query_type_parse(readBuffer, (void*) &queryType);
    if(_res != OK) {
      return _res;
    }
    (*_message)->s7_message_object_request_query_type = *queryType;


                    
    // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
    {
      uint8_t _reserved = 0;
      _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &_reserved);
      if(_res != OK) {
        return _res;
      }
      if(_reserved != 0x34) {
        printf("Expected constant value '%d' but got '%d' for reserved field.", 0x34, _reserved);
      }
    }


                    
    // Simple Field (alarmType)
    plc4c_s7_read_write_alarm_type* alarmType;
    _res = plc4c_s7_read_write_alarm_type_parse(readBuffer, (void*) &alarmType);
    if(_res != OK) {
      return _res;
    }
    (*_message)->s7_message_object_request_alarm_type = *alarmType;

  } else 
  { /* S7MessageObjectResponse */
    (*_message)->_type = plc4c_s7_read_write_s7_data_alarm_message_type_plc4c_s7_read_write_s7_message_object_response;
                    
    // Simple Field (returnCode)
    plc4c_s7_read_write_data_transport_error_code* returnCode;
    _res = plc4c_s7_read_write_data_transport_error_code_parse(readBuffer, (void*) &returnCode);
    if(_res != OK) {
      return _res;
    }
    (*_message)->s7_message_object_response_return_code = *returnCode;


                    
    // Simple Field (transportSize)
    plc4c_s7_read_write_data_transport_size* transportSize;
    _res = plc4c_s7_read_write_data_transport_size_parse(readBuffer, (void*) &transportSize);
    if(_res != OK) {
      return _res;
    }
    (*_message)->s7_message_object_response_transport_size = *transportSize;


                    
    // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
    {
      uint8_t _reserved = 0;
      _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &_reserved);
      if(_res != OK) {
        return _res;
      }
      if(_reserved != 0x00) {
        printf("Expected constant value '%d' but got '%d' for reserved field.", 0x00, _reserved);
      }
    }

  }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_s7_data_alarm_message_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_s7_data_alarm_message* _message) {
  plc4c_return_code _res = OK;

  // Const Field (functionId)
  plc4c_spi_write_unsigned_byte(writeBuffer, 8, PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_FUNCTION_ID());

  // Const Field (numberMessageObj)
  plc4c_spi_write_unsigned_byte(writeBuffer, 8, PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_NUMBER_MESSAGE_OBJ());

  // Switch Field (Depending of the current type, serialize the sub-type elements)
  switch(_message->_type) {
    case plc4c_s7_read_write_s7_data_alarm_message_type_plc4c_s7_read_write_s7_message_object_request: {

      // Const Field (variableSpec)
      plc4c_spi_write_unsigned_byte(writeBuffer, 8, PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_VARIABLE_SPEC());

      // Const Field (length)
      plc4c_spi_write_unsigned_byte(writeBuffer, 8, PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_LENGTH());

      // Simple Field (syntaxId)
      _res = plc4c_s7_read_write_syntax_id_type_serialize(writeBuffer, &_message->s7_message_object_request_syntax_id);
      if(_res != OK) {
        return _res;
      }

      // Reserved Field
      _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, 0x00);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (queryType)
      _res = plc4c_s7_read_write_query_type_serialize(writeBuffer, &_message->s7_message_object_request_query_type);
      if(_res != OK) {
        return _res;
      }

      // Reserved Field
      _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, 0x34);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (alarmType)
      _res = plc4c_s7_read_write_alarm_type_serialize(writeBuffer, &_message->s7_message_object_request_alarm_type);
      if(_res != OK) {
        return _res;
      }

      break;
    }
    case plc4c_s7_read_write_s7_data_alarm_message_type_plc4c_s7_read_write_s7_message_object_response: {

      // Simple Field (returnCode)
      _res = plc4c_s7_read_write_data_transport_error_code_serialize(writeBuffer, &_message->s7_message_object_response_return_code);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (transportSize)
      _res = plc4c_s7_read_write_data_transport_size_serialize(writeBuffer, &_message->s7_message_object_response_transport_size);
      if(_res != OK) {
        return _res;
      }

      // Reserved Field
      _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, 0x00);
      if(_res != OK) {
        return _res;
      }

      break;
    }
  }

  return OK;
}

uint16_t plc4c_s7_read_write_s7_data_alarm_message_length_in_bytes(plc4c_s7_read_write_s7_data_alarm_message* _message) {
  return plc4c_s7_read_write_s7_data_alarm_message_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_s7_data_alarm_message_length_in_bits(plc4c_s7_read_write_s7_data_alarm_message* _message) {
  uint16_t lengthInBits = 0;

  // Const Field (functionId)
  lengthInBits += 8;

  // Const Field (numberMessageObj)
  lengthInBits += 8;

  // Depending of the current type, add the length of sub-type elements ...
  switch(_message->_type) {
    case plc4c_s7_read_write_s7_data_alarm_message_type_plc4c_s7_read_write_s7_message_object_request: {

      // Const Field (variableSpec)
      lengthInBits += 8;


      // Const Field (length)
      lengthInBits += 8;


      // Simple field (syntaxId)
      lengthInBits += plc4c_s7_read_write_syntax_id_type_length_in_bits(&_message->s7_message_object_request_syntax_id);


      // Reserved Field (reserved)
      lengthInBits += 8;


      // Simple field (queryType)
      lengthInBits += plc4c_s7_read_write_query_type_length_in_bits(&_message->s7_message_object_request_query_type);


      // Reserved Field (reserved)
      lengthInBits += 8;


      // Simple field (alarmType)
      lengthInBits += plc4c_s7_read_write_alarm_type_length_in_bits(&_message->s7_message_object_request_alarm_type);

      break;
    }
    case plc4c_s7_read_write_s7_data_alarm_message_type_plc4c_s7_read_write_s7_message_object_response: {

      // Simple field (returnCode)
      lengthInBits += plc4c_s7_read_write_data_transport_error_code_length_in_bits(&_message->s7_message_object_response_return_code);


      // Simple field (transportSize)
      lengthInBits += plc4c_s7_read_write_data_transport_size_length_in_bits(&_message->s7_message_object_response_transport_size);


      // Reserved Field (reserved)
      lengthInBits += 8;

      break;
    }
  }

  return lengthInBits;
}

