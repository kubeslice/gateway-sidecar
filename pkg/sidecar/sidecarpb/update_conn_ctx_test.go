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
	"testing"
	"time"

	"github.com/kubeslice/gateway-sidecar/pkg/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	st "google.golang.org/grpc/status"
)

func TestUpdateConnCtx(t *testing.T) {

	var tunCheck *status.TunnelChecker
	statusMonitor := status.NewMonitor(log)
	tunCheck = status.NewTunnelChecker(log).(*status.TunnelChecker)
	mod, err := statusMonitor.RegisterCheck(&status.Config{
		Name:     "TestTunnCheck",
		Checker:  tunCheck,
		Interval: time.Second * 6,
	})

	if err != nil {
		fmt.Println("failed to register tunnel checker", err.Error())
	}

	SetStatusMonitor(statusMonitor)
	tunCheck.UpdateExecModule(mod)

	tests := []struct {
		testName        string
		req             *SliceGwConnectionContext
		res             *SidecarResponse
		errCode         codes.Code
		errMsg          string
		ctxCancel       bool
		isGwInitialized bool
	}{
		{
			"It should sucessfull update connection context",
			&SliceGwConnectionContext{SliceId: "SliceId", LocalSliceGwId: "LocalSliceGwId", LocalSliceGwVpnIP: "LocalSliceGwVpnIP", LocalSliceGwHostType: SliceGwHostType_SLICE_GW_CLIENT, LocalSliceGwNsmSubnet: "LocalSliceGwNsmSubnet", LocalSliceGwNodeIP: "LocalSliceGwNodeIP", LocalSliceGwNodePort: "LocalSliceGwNodePort", RemoteSliceGwId: "RemoteSliceGwId", RemoteSliceGwVpnIP: "1.1.1.1", RemoteSliceGwHostType: SliceGwHostType_SLICE_GW_SERVER, RemoteSliceGwNodeIP: "RemoteSliceGwNodeIP", RemoteSliceGwNsmSubnet: "10.12.11.1/16", RemoteSliceGwNodePort: "RemoteSliceGwNodePort"},
			&SidecarResponse{StatusMsg: "Connection Context Updated Successfully"},
			codes.OK,
			"",
			false,
			true,
		},
		{
			"Check For Invalid Remote Slice Gateway VPN IP",
			&SliceGwConnectionContext{},
			&SidecarResponse{StatusMsg: ""},
			codes.InvalidArgument,
			"Invalid Remote Slice Gateway VPN IP",
			false,
			true,
		},
		{
			"Check For Invalid Remote Slice Gateway Nsm Subnet",
			&SliceGwConnectionContext{RemoteSliceGwVpnIP: "1.1.1.1"},
			&SidecarResponse{StatusMsg: ""},
			codes.InvalidArgument,
			"Invalid Remote Slice Gateway Subnet",
			false,
			true,
		},
		{
			"Check For Error while Parsing CIDR: local nsm subnet",
			&SliceGwConnectionContext{RemoteSliceGwVpnIP: "1.1.1.1", RemoteSliceGwNsmSubnet: "10.12.11.1"},
			&SidecarResponse{StatusMsg: ""},
			codes.InvalidArgument,
			"Error in Parsing CIDR",
			false,
			true,
		},
		{
			"Test for cancelled context",
			&SliceGwConnectionContext{SliceId: "SliceId", LocalSliceGwId: "LocalSliceGwId", LocalSliceGwVpnIP: "LocalSliceGwVpnIP", LocalSliceGwHostType: SliceGwHostType_SLICE_GW_CLIENT, LocalSliceGwNsmSubnet: "LocalSliceGwNsmSubnet", LocalSliceGwNodeIP: "LocalSliceGwNodeIP", LocalSliceGwNodePort: "LocalSliceGwNodePort", RemoteSliceGwId: "RemoteSliceGwId", RemoteSliceGwVpnIP: "1.1.1.1", RemoteSliceGwHostType: SliceGwHostType_SLICE_GW_SERVER, RemoteSliceGwNodeIP: "RemoteSliceGwNodeIP", RemoteSliceGwNsmSubnet: "10.12.11.1/16", RemoteSliceGwNodePort: "RemoteSliceGwNodePort"},
			&SidecarResponse{StatusMsg: ""},
			codes.Canceled,
			"context canceled",
			true,
			true,
		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewGwSidecarServiceClient(conn)
	var errVal error = nil
	fmt.Println(errVal)

	request := SliceGwConnectionContext{}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			request = *tt.req

			if tt.req == nil {
				t.Error("Connection Context is Empty")
			}

			if tt.ctxCancel {
				cancel()
			}

			response, err := client.UpdateConnectionContext(ctx, &request)

			if response != nil {
				if response.StatusMsg != tt.res.StatusMsg {
					t.Error("response: expected", tt.res, "received", response)
				}
			}
			if err != nil {
				if er, ok := st.FromError(err); ok {
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
