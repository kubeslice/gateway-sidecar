/*  Copyright (c) 2022 Avesha, Inc. All rights reserved.
 *
 *  SPDX-License-Identifier: Apache-2.0
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package sidecar

import (
	"context"
	"fmt"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

var qosProfile *SliceQosProfile

func dialer() func(context.Context, string) (net.Conn, error) {

	listner := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()

	RegisterGwSidecarServiceServer(server, &GwSidecar{})

	go func() {
		if err := server.Serve(listner); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listner.Dial()
	}
}

func TestSliceQosProfile(t *testing.T) {

	tests := []struct {
		testName string
		tcInfo   *TcInfo
		req      *SliceQosProfile
		res      *SidecarResponse

		errCode   codes.Code
		errMsg    string
		ctxCancel bool
	}{
		{
			"It should sucessfully update slice Qos profile",
			&TcInfo{class: "class", bwCeiling: 1, bwGuaranteed: 1, priority: 1},
			&SliceQosProfile{SliceName: "SliceName", SliceId: "SliceId", TcType: TcType_BANDWIDTH_CONTROL, ClassType: ClassType_HTB, BwCeiling: 1, BwGuaranteed: 1, Priority: 1, DscpClass: "DscpClass"},
			&SidecarResponse{StatusMsg: "Slice QoS policy enforced successfully"},
			codes.OK,
			"",
			false,
		},
		{
			"Test for cancelled context",
			&TcInfo{},
			&SliceQosProfile{SliceName: "SliceName", SliceId: "SliceId", TcType: TcType_BANDWIDTH_CONTROL, ClassType: ClassType_HTB, BwCeiling: 1, BwGuaranteed: 1, Priority: 1, DscpClass: "DscpClass"},
			&SidecarResponse{StatusMsg: ""},
			codes.Canceled,
			"context canceled",
			true,
		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	conn, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewGwSidecarServiceClient(conn)

	request := SliceQosProfile{}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			request = *tt.req

			if tt.ctxCancel {
				cancel()
			}

			fmt.Println(qosProfile)

			if tt.req == nil {
				fmt.Println(true)
			}

			response, err := client.UpdateSliceQosProfile(ctx, &request)

			if response != nil {
				if response.StatusMsg != tt.res.StatusMsg {
					t.Error("response: expected", tt.res, "received", response)
				}
			}
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}
