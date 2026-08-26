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

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/vishvananda/netlink"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type GwSidecar struct {
	UnimplementedGwSidecarServiceServer
}

// Stores the node port of the remote cluster.
// We need this to set the dscp config when the mode of the Gw is of type CLIENT, in
// which case the remote cluster would be the SERVER and the client is connected to
// it over this node port.
var SliceGwRemoteClusterNodePort string = ""

// Checks if vpphost interface is present on the host.
// The vpphost interface is used to set up a network connection between the host kernel stack
// and the vpp data plane stack.
func checkIfVppIntfPresent() bool {
	vppInterface, err := net.InterfaceByName("vpphost")
	if err != nil {
		return false
	}
	return vppInterface != nil
}

// GetSliceGwRemotePodName get the remote GwPodName
func (s *GwSidecar) GetSliceGwRemotePodName(ctx context.Context, remoteGwVpnIP *RemoteGwVpnIP) (*GwPodStatus, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
	if remoteGwVpnIP == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Remote Gateway VPN IP is nil")
	}
	if remoteGwVpnIP.GetRemoteGwVpnIP() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Remote Slice Gateway VPN IP")
	}

	// Call the GRPC client to get the RemoteGW PodName
	address := remoteGwVpnIP.GetRemoteGwVpnIP() + ":5000"
	log.Infof("Attempting to dial remote gateway at: %s", address)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("Failed to dial remote gateway: %v", err)
		return nil, status.Errorf(codes.Unavailable, "Unable to connect to remote gateway pod")
	}
	defer conn.Close()

	client := NewGwSidecarServiceClient(conn)
	res, err := client.GetStatus(context.Background(), &empty.Empty{})
	if err != nil {
		log.Errorf("Failed to get remote pod status: %v", err)
		return nil, status.Errorf(codes.Unavailable, "Unable to get the remote pod status")
	}

	if res == nil {
		log.Warn("Received nil response from remote cluster")
		return nil, status.Errorf(codes.Internal, "Remote cluster returned empty response")
	}

	log.Infof("Received response from remote cluster: %v", res)
	return res, nil
}

// GetStatus get the status of sidecar.
func (s *GwSidecar) GetStatus(ctx context.Context, in *empty.Empty) (*GwPodStatus, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
	podStatus, err := getGwPodStatus()
	return podStatus, err
}

func (s *GwSidecar) UpdateConnectionContext(ctx context.Context, conContext *SliceGwConnectionContext) (*SidecarResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
	if conContext == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Connection Context is Empty")
	}
	if conContext.GetRemoteSliceGwVpnIP() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Remote Slice Gateway VPN IP")
	}
	log.Infof("conContext : %v", conContext)
	err := updateGwStatusWithConContext(conContext)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Failed to update the Connection context,tunnel is not up yet!")
	}
	if conContext.GetRemoteSliceGwNsmSubnet() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Remote Slice Gateway Subnet")
	}

	if SliceGwRemoteClusterNodePort == "" {
		SliceGwRemoteClusterNodePort = conContext.GetRemoteSliceGwNodePort()
	}

	// Add Gateway Route as follows
	// route add -net  <remote-subnet> netmask <255.255.255.0> gw <remove-vpn-ip>
	_, dstIPNet, err := net.ParseCIDR(conContext.GetRemoteSliceGwNsmSubnet())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error in Parsing CIDR")
	}
	gwIP := net.ParseIP(conContext.GetRemoteSliceGwVpnIP())

	// Program the tunnel route as two halves of the requested subnet (one bit more
	// specific) rather than the subnet itself. NSM continuously reconciles a route
	// for the whole slice subnet via the nsm interface using the SAME prefix we
	// want (e.g. 10.11.0.0/16); a RouteReplace on that exact prefix only wins until
	// NSM re-asserts, so the route FLAPS between nsm0 and tun0 and spoke-to-spoke
	// traffic intermittently black-holes. Installing e.g. 10.11.0.0/17 + 10.11.128.0/17
	// makes the tunnel routes strictly more specific than NSM's /16, so longest-prefix
	// match always selects the tunnel and NSM can never overwrite them. The local
	// subnet route (a /20 that NSM owns) stays more specific still, so local delivery
	// is unaffected. RouteReplace keeps each write idempotent across reconciles.
	for _, half := range moreSpecificHalves(dstIPNet) {
		route := netlink.Route{Dst: half, Gw: gwIP}
		log.Infof("RouteReplace args %v, %v ", half, gwIP)
		if err := netlink.RouteReplace(&route); err != nil {
			log.Errorf("Gateway Pod RouteReplace Failed : %v", err)
		}
	}

	if checkIfVppIntfPresent() {
		vppGwIP := net.ParseIP("10.255.255.254")
		_, localGwNsmSubnetIP, err := net.ParseCIDR(conContext.GetLocalSliceGwNsmSubnet())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Error in Parsing CIDR: local nsm subnet")
		}
		vppGwRoute := netlink.Route{Dst: localGwNsmSubnetIP, Gw: vppGwIP}
		log.Infof("RouteAdd args %v, %v ", localGwNsmSubnetIP, vppGwIP)
		if err := netlink.RouteAdd(&vppGwRoute); err != nil {
			log.Errorf("VPP Gateway RouteAdd Failed : %v", err)
		}
	}
	// Clamp the TCP MSS of connections crossing the tunnel to the path MTU. The
	// pod-side NSM interface is MTU 1500 while the OpenVPN tunnel (tun0) is smaller,
	// so full-size segments would otherwise be silently dropped (no "fragmentation
	// needed" ICMP -> PMTU black hole) and TCP would stall. Required for
	// spoke-to-spoke traffic (which crosses two tunnels), and also fixes plain
	// hub-to-spoke transfers.
	ensureTunnelMSSClamp()

	log.Infof("Connection Context Updated Successfully")

	return &SidecarResponse{StatusMsg: "Connection Context Updated Successfully"}, nil
}

// moreSpecificHalves splits a subnet into its two halves, one bit more specific
// than the input (e.g. 10.11.0.0/16 -> 10.11.0.0/17, 10.11.128.0/17). This is used
// so the tunnel route out-specifies NSM's same-prefix route and cannot be
// overwritten by NSM's reconcile. A host route (or a subnet that cannot be split
// further) is returned unchanged.
func moreSpecificHalves(dst *net.IPNet) []*net.IPNet {
	ones, bits := dst.Mask.Size()
	if bits == 0 || ones >= bits {
		return []*net.IPNet{dst}
	}
	half := net.CIDRMask(ones+1, bits)
	lower := dst.IP.Mask(half)
	// upper half: set the bit at position `ones` in the network address
	upper := make(net.IP, len(lower))
	copy(upper, lower)
	upper[ones/8] |= 1 << uint(7-(ones%8))
	return []*net.IPNet{
		{IP: lower, Mask: half},
		{IP: upper, Mask: half},
	}
}

// ensureTunnelMSSClamp installs, idempotently, an iptables rule that clamps the
// TCP MSS of forwarded connections leaving via the tunnel interface (tun0) to the
// tunnel's path MTU. This is the standard remedy for tunnel MTU mismatches used
// by CNIs such as Flannel and Calico (--clamp-mss-to-pmtu).
// tunnelMSSClampCommands returns the iptables check (-C) and add (-A) commands
// that clamp the TCP MSS of forwarded connections leaving via iface to the path
// MTU. Kept as a pure, interface-parameterized helper so the rule can be unit
// tested (and the interface generalized beyond tun0 later).
func tunnelMSSClampCommands(iface string) (checkCmd, addCmd string) {
	const spec = "-p tcp --tcp-flags SYN,RST SYN -j TCPMSS --clamp-mss-to-pmtu"
	checkCmd = fmt.Sprintf("iptables -t mangle -C FORWARD -o %s %s", iface, spec)
	addCmd = fmt.Sprintf("iptables -t mangle -A FORWARD -o %s %s", iface, spec)
	return checkCmd, addCmd
}

func ensureTunnelMSSClamp() {
	checkCmd, addCmd := tunnelMSSClampCommands("tun0")
	if _, err := runCommand(checkCmd); err == nil {
		return // rule already present
	}
	if out, err := runCommand(addCmd); err != nil {
		log.Errorf("Failed to add TCP MSS clamp rule on tun0: %v (%v)", err, out)
		return
	}
	log.Infof("Installed TCP MSS clamp (clamp-mss-to-pmtu) on tun0")
}

func (s *GwSidecar) UpdateSliceQosProfile(ctx context.Context, qosProfile *SliceQosProfile) (*SidecarResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "Client canceled, ignoring qos update message.")
	}
	if qosProfile == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Qos profile message is empty")
	}

	//settings.Log.Infof("SliceQosProfile : %v", qosProfile)

	err := s.enforceSliceQosPolicy(
		&SliceQosProfile{
			ClassType:    qosProfile.GetClassType(),
			BwCeiling:    qosProfile.GetBwCeiling(),
			BwGuaranteed: qosProfile.GetBwGuaranteed(),
			Priority:     qosProfile.GetPriority(),
			DscpClass:    qosProfile.GetDscpClass(),
		},
	)
	if err != nil {
		//settings.Log.Errorf("Failed to enforce QoS policy: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to enforce QoS policy: %v", err)
	}
	log.Infof("Slice QoS policy enforced successfully")
	return &SidecarResponse{StatusMsg: "Slice QoS policy enforced successfully"}, nil
}
