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
	"runtime/debug"
	"testing"

	nt "github.com/kubeslice/gateway-sidecar/pkg/nettools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	st "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	NodeIP  = "156.178.1.1"
	NsmIP   = "192.178.1.1"
	PeerIP  = "192.168.0.2"
	LocalIP = "192.168.0.1"
)

func getTheIpAndName() (string, string) {

	interfaceNames := make([]string, 2)

	allInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err while getting all the interfaces")
	}

	for _, a := range allInterfaces {
		interfaceNames = append(interfaceNames, a.Name)
	}

	infos, err := nt.GetInterfaceInfos(interfaceNames[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	return infos[0].Name, infos[0].IP
}

func TestGwStatus(t *testing.T) {

	_, InfIP := getTheIpAndName()

	tests := []struct {
		testName  string
		res       *GwPodStatus
		errCode   codes.Code
		errMsg    string
		ctxCancel bool
	}{
		{
			"It should pass",
			&GwPodStatus{NodeIP: NodeIP, GatewayPodIP: InfIP,
				NsmIntfStatus: &NsmInterfaceStatus{NsmInterfaceName: "nsm0", NsmIP: NsmIP},
				TunnelStatus:  &TunnelInterfaceStatus{NetInterface: "veth0", LocalIP: LocalIP, PeerIP: PeerIP, Latency: 1, RxRate: 1, TxRate: 1}},
			codes.OK,
			"",
			false,
		},
		{
			"Test for cancelled context",
			&GwPodStatus{NodeIP: NodeIP, GatewayPodIP: InfIP,
				NsmIntfStatus: &NsmInterfaceStatus{NsmInterfaceName: "nsm0", NsmIP: NsmIP},
				TunnelStatus:  &TunnelInterfaceStatus{NetInterface: "veth0", LocalIP: LocalIP, PeerIP: PeerIP, Latency: 1, RxRate: 1, TxRate: 1}},
			codes.Canceled,
			"Client cancelled, abandoning.",
			true,
		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer()))
	AssertNoError(t, err)

	defer conn.Close()

	client := NewGwSidecarServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			response, err := client.GetStatus(ctx, &emptypb.Empty{})

			if tt.ctxCancel {
				cancel()
			}

			response.NodeIP = NodeIP
			response.GetNsmIntfStatus().NsmIP = NsmIP

			tunnleStatus := TunnelInterfaceStatus{
				NetInterface: "veth0",
				Latency:      1,
				RxRate:       1,
				TxRate:       1,
				LocalIP:      LocalIP,
				PeerIP:       PeerIP,
			}
			response.TunnelStatus = &tunnleStatus

			if response != nil {
				if response.GetNodeIP() != tt.res.NodeIP {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetGatewayPodIP() != tt.res.GatewayPodIP {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetNsmIntfStatus().NsmInterfaceName != tt.res.NsmIntfStatus.NsmInterfaceName {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetNsmIntfStatus().NsmIP != tt.res.NsmIntfStatus.NsmIP {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetTunnelStatus().NetInterface != tt.res.TunnelStatus.NetInterface {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetTunnelStatus().Latency != tt.res.TunnelStatus.Latency {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetTunnelStatus().RxRate != tt.res.TunnelStatus.RxRate {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetTunnelStatus().TxRate != tt.res.TunnelStatus.TxRate {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetTunnelStatus().LocalIP != tt.res.TunnelStatus.LocalIP {
					t.Error("response: expected", tt.res, "received", response)
				}
				if response.GetTunnelStatus().PeerIP != tt.res.TunnelStatus.PeerIP {
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

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected No Error but got %s, Stack:\n%s", err, string(debug.Stack()))
	}
}
