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
package org.apache.plc4x.examples.plc4j.s7event;

import org.apache.commons.codec.binary.Hex;
import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcSubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionResponse;
import org.apache.plc4x.java.s7.events.S7CyclicEvent;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Map;


/**
 * Example for capturing events generated from a Siemens S7-300, S7-400 or VIPA PLC.
 * Support for mode events ("MODE"), system events ("SYS"), user events ("USR")
 * and alarms ("ALM").
 * Each consumer shows the tags and associated values of the "map" containing
 * the event parameters.
 */
public class CycSubscription {

    private static final Logger logger = LoggerFactory.getLogger(CycSubscription.class);

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) throws Exception {

        try (PlcConnection connection = new DefaultPlcDriverManager().getConnection("s7://10.10.1.33?remote-rack=0&remote-slot=3&controller-type=S7_400")) {
            final PlcSubscriptionRequest.Builder subscription = connection.subscriptionRequestBuilder();

            //subscription.addEventTagAddress("myCYC", "CYC(B01SEC:5):%DB9002.DBB0[1]");
            subscription.addEventTagAddress("myCYC", "CYC(B01SEC:5):%MB190:BYTE");

            final PlcSubscriptionRequest sub = subscription.build();
            final PlcSubscriptionResponse subresponse = sub.execute().get();

            //Si todo va bien con la subscripción puedo
//            subresponse
//                .getSubscriptionHandle("myMODE")
//                .register(msg -> {
//                    System.out.println("******** S7ModeEvent ********");
//                    Map<String, Object> map = ((S7ModeEvent) msg).getMap();
//                    map.forEach((x, y) -> {
//                        System.out.println(x + " : " + y);
//                    });
//                    System.out.println("****************************");
//                });
//
//            subresponse
//                .getSubscriptionHandle("mySYS")
//                .register(msg -> {
//                    System.out.println("******** S7SysEvent ********");
//                    Map<String, Object> map = ((S7SysEvent) msg).getMap();
//                    map.forEach((x, y) -> {
//                        if ("INFO1".equals(x)) {
//                            System.out.println(x + " : " + String.format("0x%04X", y));
//                        } else if ("INFO2".equals(x)) {
//                            System.out.println(x + " : " + String.format("0x%08X", y));    
//                        } else System.out.println(x + " : " + y);
//                    });
//                    System.out.println("****************************");
//                });
//
//            subresponse
//                .getSubscriptionHandle("myUSR")
//                .register(msg -> {
//                    System.out.println("******** S7UserEvent *******");
//                    Map<String, Object> map = ((S7UserEvent) msg).getMap();
//                    map.forEach((x, y) -> {
//                        System.out.println(x + " : " + y);
//                    });
//                    System.out.println("****************************");
//                });

            subresponse
                .getSubscriptionHandle("myCYC")
                .register(msg -> {
                    System.out.println("******** CYC Event *********");
                    Map<String, Object> map = ((S7CyclicEvent) msg).getMap();
                    map.forEach((x, y) -> {
                        if (x.startsWith("DATA_", 0)) {
                            System.out.println("Longitud de datos: " + ((byte[]) y).length);
                            System.out.println(x + ": " + Hex.encodeHexString((byte[]) y));
                        } else
                            System.out.println(x + " : " + y);
                    });
                    System.out.println("****************************");
                });

            System.out.println("Waiting for events");

            Thread.sleep(10000);

            System.out.println("Bye...");

        }
    }

}
