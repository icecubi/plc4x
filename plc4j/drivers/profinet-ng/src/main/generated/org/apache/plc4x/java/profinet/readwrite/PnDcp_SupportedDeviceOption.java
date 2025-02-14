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
package org.apache.plc4x.java.profinet.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class PnDcp_SupportedDeviceOption implements Message {

  // Properties.
  protected final PnDcp_BlockOptions option;
  protected final short suboption;

  public PnDcp_SupportedDeviceOption(PnDcp_BlockOptions option, short suboption) {
    super();
    this.option = option;
    this.suboption = suboption;
  }

  public PnDcp_BlockOptions getOption() {
    return option;
  }

  public short getSuboption() {
    return suboption;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnDcp_SupportedDeviceOption");

    // Simple Field (option)
    writeSimpleEnumField(
        "option",
        "PnDcp_BlockOptions",
        option,
        new DataWriterEnumDefault<>(
            PnDcp_BlockOptions::getValue,
            PnDcp_BlockOptions::name,
            writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (suboption)
    writeSimpleField(
        "suboption",
        suboption,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnDcp_SupportedDeviceOption");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PnDcp_SupportedDeviceOption _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (option)
    lengthInBits += 8;

    // Simple field (suboption)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static PnDcp_SupportedDeviceOption staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static PnDcp_SupportedDeviceOption staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnDcp_SupportedDeviceOption");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PnDcp_BlockOptions option =
        readEnumField(
            "option",
            "PnDcp_BlockOptions",
            new DataReaderEnumDefault<>(
                PnDcp_BlockOptions::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short suboption =
        readSimpleField(
            "suboption",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnDcp_SupportedDeviceOption");
    // Create the instance
    PnDcp_SupportedDeviceOption _pnDcp_SupportedDeviceOption;
    _pnDcp_SupportedDeviceOption = new PnDcp_SupportedDeviceOption(option, suboption);
    return _pnDcp_SupportedDeviceOption;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_SupportedDeviceOption)) {
      return false;
    }
    PnDcp_SupportedDeviceOption that = (PnDcp_SupportedDeviceOption) o;
    return (getOption() == that.getOption()) && (getSuboption() == that.getSuboption()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOption(), getSuboption());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
