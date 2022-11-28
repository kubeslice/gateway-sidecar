//  Copyright (c) 2022 Avesha, Inc. All rights reserved.
//
//  SPDX-License-Identifier: Apache-2.0
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: pkg/sidecar/sidecarpb/gw_sidecar.proto

package sidecar

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// slice gateway-host-type
type SliceGwHostType int32

const (
	SliceGwHostType_SLICE_GW_SERVER SliceGwHostType = 0
	SliceGwHostType_SLICE_GW_CLIENT SliceGwHostType = 1
)

// Enum value maps for SliceGwHostType.
var (
	SliceGwHostType_name = map[int32]string{
		0: "SLICE_GW_SERVER",
		1: "SLICE_GW_CLIENT",
	}
	SliceGwHostType_value = map[string]int32{
		"SLICE_GW_SERVER": 0,
		"SLICE_GW_CLIENT": 1,
	}
)

func (x SliceGwHostType) Enum() *SliceGwHostType {
	p := new(SliceGwHostType)
	*p = x
	return p
}

func (x SliceGwHostType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SliceGwHostType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_enumTypes[0].Descriptor()
}

func (SliceGwHostType) Type() protoreflect.EnumType {
	return &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_enumTypes[0]
}

func (x SliceGwHostType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SliceGwHostType.Descriptor instead.
func (SliceGwHostType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{0}
}

// TcType represents Traffic Control Type.
type TcType int32

const (
	TcType_BANDWIDTH_CONTROL TcType = 0
)

// Enum value maps for TcType.
var (
	TcType_name = map[int32]string{
		0: "BANDWIDTH_CONTROL",
	}
	TcType_value = map[string]int32{
		"BANDWIDTH_CONTROL": 0,
	}
)

func (x TcType) Enum() *TcType {
	p := new(TcType)
	*p = x
	return p
}

func (x TcType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TcType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_enumTypes[1].Descriptor()
}

func (TcType) Type() protoreflect.EnumType {
	return &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_enumTypes[1]
}

func (x TcType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TcType.Descriptor instead.
func (TcType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{1}
}

type ClassType int32

const (
	ClassType_HTB ClassType = 0
	ClassType_TBF ClassType = 1
)

// Enum value maps for ClassType.
var (
	ClassType_name = map[int32]string{
		0: "HTB",
		1: "TBF",
	}
	ClassType_value = map[string]int32{
		"HTB": 0,
		"TBF": 1,
	}
)

func (x ClassType) Enum() *ClassType {
	p := new(ClassType)
	*p = x
	return p
}

func (x ClassType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClassType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_enumTypes[2].Descriptor()
}

func (ClassType) Type() protoreflect.EnumType {
	return &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_enumTypes[2]
}

func (x ClassType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClassType.Descriptor instead.
func (ClassType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{2}
}

// SidecarResponse represents the Sidecar response format.
type SidecarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusMsg string `protobuf:"bytes,1,opt,name=statusMsg,proto3" json:"statusMsg,omitempty"`
}

func (x *SidecarResponse) Reset() {
	*x = SidecarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SidecarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SidecarResponse) ProtoMessage() {}

func (x *SidecarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SidecarResponse.ProtoReflect.Descriptor instead.
func (*SidecarResponse) Descriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{0}
}

func (x *SidecarResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

// Slice QoS Profile
type SliceQosProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the slice
	SliceName string `protobuf:"bytes,1,opt,name=sliceName,proto3" json:"sliceName,omitempty"`
	// Slice Identifier
	SliceId string `protobuf:"bytes,2,opt,name=sliceId,proto3" json:"sliceId,omitempty"`
	// Name of the QoS profile attached to the slice
	QosProfileName string `protobuf:"bytes,3,opt,name=qosProfileName,proto3" json:"qosProfileName,omitempty"`
	// TC type -  Bandwidth control
	TcType TcType `protobuf:"varint,4,opt,name=tcType,proto3,enum=sidecar.TcType" json:"tcType,omitempty"`
	// ClassType - HTB   ( HTB)
	ClassType ClassType `protobuf:"varint,5,opt,name=ClassType,proto3,enum=sidecar.ClassType" json:"ClassType,omitempty"`
	// Bandwidth Ceiling in Mbps  - 5 Mbps (100k - 100 Mbps)
	BwCeiling uint32 `protobuf:"varint,6,opt,name=bwCeiling,proto3" json:"bwCeiling,omitempty"`
	// Bandwidth Guaranteed -  1 Mbps ( 100k- 100 Mbps)
	BwGuaranteed uint32 `protobuf:"varint,7,opt,name=bwGuaranteed,proto3" json:"bwGuaranteed,omitempty"`
	// Priority - 2 (Number 0-3)
	Priority uint32 `protobuf:"varint,8,opt,name=priority,proto3" json:"priority,omitempty"`
	// Dscp class to mark inter cluster traffic
	DscpClass string `protobuf:"bytes,9,opt,name=dscpClass,proto3" json:"dscpClass,omitempty"`
}

func (x *SliceQosProfile) Reset() {
	*x = SliceQosProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliceQosProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceQosProfile) ProtoMessage() {}

func (x *SliceQosProfile) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceQosProfile.ProtoReflect.Descriptor instead.
func (*SliceQosProfile) Descriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{1}
}

func (x *SliceQosProfile) GetSliceName() string {
	if x != nil {
		return x.SliceName
	}
	return ""
}

func (x *SliceQosProfile) GetSliceId() string {
	if x != nil {
		return x.SliceId
	}
	return ""
}

func (x *SliceQosProfile) GetQosProfileName() string {
	if x != nil {
		return x.QosProfileName
	}
	return ""
}

func (x *SliceQosProfile) GetTcType() TcType {
	if x != nil {
		return x.TcType
	}
	return TcType_BANDWIDTH_CONTROL
}

func (x *SliceQosProfile) GetClassType() ClassType {
	if x != nil {
		return x.ClassType
	}
	return ClassType_HTB
}

func (x *SliceQosProfile) GetBwCeiling() uint32 {
	if x != nil {
		return x.BwCeiling
	}
	return 0
}

func (x *SliceQosProfile) GetBwGuaranteed() uint32 {
	if x != nil {
		return x.BwGuaranteed
	}
	return 0
}

func (x *SliceQosProfile) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *SliceQosProfile) GetDscpClass() string {
	if x != nil {
		return x.DscpClass
	}
	return ""
}

// TunnelInterfaceStatus represents Tunnel Interface Status.
type TunnelInterfaceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Net Interface
	NetInterface string `protobuf:"bytes,1,opt,name=netInterface,proto3" json:"netInterface,omitempty"`
	// Local IP
	LocalIP string `protobuf:"bytes,2,opt,name=localIP,proto3" json:"localIP,omitempty"`
	// Peer IP
	PeerIP string `protobuf:"bytes,3,opt,name=peerIP,proto3" json:"peerIP,omitempty"`
	// Latency
	Latency uint64 `protobuf:"varint,4,opt,name=latency,proto3" json:"latency,omitempty"`
	// Transmit Rate
	TxRate uint64 `protobuf:"varint,5,opt,name=txRate,proto3" json:"txRate,omitempty"`
	// Receive Rate
	RxRate uint64 `protobuf:"varint,6,opt,name=rxRate,proto3" json:"rxRate,omitempty"`
}

func (x *TunnelInterfaceStatus) Reset() {
	*x = TunnelInterfaceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelInterfaceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelInterfaceStatus) ProtoMessage() {}

func (x *TunnelInterfaceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelInterfaceStatus.ProtoReflect.Descriptor instead.
func (*TunnelInterfaceStatus) Descriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{2}
}

func (x *TunnelInterfaceStatus) GetNetInterface() string {
	if x != nil {
		return x.NetInterface
	}
	return ""
}

func (x *TunnelInterfaceStatus) GetLocalIP() string {
	if x != nil {
		return x.LocalIP
	}
	return ""
}

func (x *TunnelInterfaceStatus) GetPeerIP() string {
	if x != nil {
		return x.PeerIP
	}
	return ""
}

func (x *TunnelInterfaceStatus) GetLatency() uint64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

func (x *TunnelInterfaceStatus) GetTxRate() uint64 {
	if x != nil {
		return x.TxRate
	}
	return 0
}

func (x *TunnelInterfaceStatus) GetRxRate() uint64 {
	if x != nil {
		return x.RxRate
	}
	return 0
}

// NsmInterfaceStatus represents the status of NSM Interface to Slice Router.
type NsmInterfaceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NSM Interface Name
	NsmInterfaceName string `protobuf:"bytes,1,opt,name=nsmInterfaceName,proto3" json:"nsmInterfaceName,omitempty"`
	// NSM IP
	NsmIP string `protobuf:"bytes,2,opt,name=nsmIP,proto3" json:"nsmIP,omitempty"`
}

func (x *NsmInterfaceStatus) Reset() {
	*x = NsmInterfaceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NsmInterfaceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NsmInterfaceStatus) ProtoMessage() {}

func (x *NsmInterfaceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NsmInterfaceStatus.ProtoReflect.Descriptor instead.
func (*NsmInterfaceStatus) Descriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{3}
}

func (x *NsmInterfaceStatus) GetNsmInterfaceName() string {
	if x != nil {
		return x.NsmInterfaceName
	}
	return ""
}

func (x *NsmInterfaceStatus) GetNsmIP() string {
	if x != nil {
		return x.NsmIP
	}
	return ""
}

// GwPodStatus represents overall status of Pod.
type GwPodStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node IP.
	NodeIP string `protobuf:"bytes,1,opt,name=nodeIP,proto3" json:"nodeIP,omitempty"`
	// Gateway Pod IP.
	GatewayPodIP string `protobuf:"bytes,2,opt,name=gatewayPodIP,proto3" json:"gatewayPodIP,omitempty"`
	// Gateway Pod Name
	GatewayPodName string `protobuf:"bytes,3,opt,name=gatewayPodName,proto3" json:"gatewayPodName,omitempty"`
	// The Tunnel Interface Status.
	TunnelStatus *TunnelInterfaceStatus `protobuf:"bytes,4,opt,name=tunnelStatus,proto3" json:"tunnelStatus,omitempty"`
	// NSM Interface Status
	NsmIntfStatus *NsmInterfaceStatus `protobuf:"bytes,5,opt,name=nsmIntfStatus,proto3" json:"nsmIntfStatus,omitempty"`
}

func (x *GwPodStatus) Reset() {
	*x = GwPodStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GwPodStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GwPodStatus) ProtoMessage() {}

func (x *GwPodStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GwPodStatus.ProtoReflect.Descriptor instead.
func (*GwPodStatus) Descriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{4}
}

func (x *GwPodStatus) GetNodeIP() string {
	if x != nil {
		return x.NodeIP
	}
	return ""
}

func (x *GwPodStatus) GetGatewayPodIP() string {
	if x != nil {
		return x.GatewayPodIP
	}
	return ""
}

func (x *GwPodStatus) GetGatewayPodName() string {
	if x != nil {
		return x.GatewayPodName
	}
	return ""
}

func (x *GwPodStatus) GetTunnelStatus() *TunnelInterfaceStatus {
	if x != nil {
		return x.TunnelStatus
	}
	return nil
}

func (x *GwPodStatus) GetNsmIntfStatus() *NsmInterfaceStatus {
	if x != nil {
		return x.NsmIntfStatus
	}
	return nil
}

// SliceGwConnectionContext - Slice Gateway Connection Context.
type SliceGwConnectionContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Slice-Id
	SliceId string `protobuf:"bytes,1,opt,name=sliceId,proto3" json:"sliceId,omitempty"`
	// Local slice gateway ID
	LocalSliceGwId string `protobuf:"bytes,2,opt,name=localSliceGwId,proto3" json:"localSliceGwId,omitempty"`
	// Local slice gateway VPN IP
	LocalSliceGwVpnIP string `protobuf:"bytes,3,opt,name=localSliceGwVpnIP,proto3" json:"localSliceGwVpnIP,omitempty"`
	// Local slice gateway-host-type  -  client/server
	LocalSliceGwHostType SliceGwHostType `protobuf:"varint,4,opt,name=localSliceGwHostType,proto3,enum=sidecar.SliceGwHostType" json:"localSliceGwHostType,omitempty"`
	// Local slice gateway NSM Subnet
	LocalSliceGwNsmSubnet string `protobuf:"bytes,5,opt,name=localSliceGwNsmSubnet,proto3" json:"localSliceGwNsmSubnet,omitempty"`
	// Local slice gateway Node IP
	LocalSliceGwNodeIP string `protobuf:"bytes,6,opt,name=localSliceGwNodeIP,proto3" json:"localSliceGwNodeIP,omitempty"`
	// Local slice gateway Node Port
	LocalSliceGwNodePort string `protobuf:"bytes,7,opt,name=localSliceGwNodePort,proto3" json:"localSliceGwNodePort,omitempty"`
	// Remote slice gateway ID
	RemoteSliceGwId string `protobuf:"bytes,8,opt,name=remoteSliceGwId,proto3" json:"remoteSliceGwId,omitempty"`
	// Remote slice gateway VPN IP
	RemoteSliceGwVpnIP string `protobuf:"bytes,9,opt,name=remoteSliceGwVpnIP,proto3" json:"remoteSliceGwVpnIP,omitempty"`
	// Remote-slice gateway-host-type client or server
	RemoteSliceGwHostType SliceGwHostType `protobuf:"varint,10,opt,name=remoteSliceGwHostType,proto3,enum=sidecar.SliceGwHostType" json:"remoteSliceGwHostType,omitempty"`
	// Remote slice gateway NSM subnet
	RemoteSliceGwNsmSubnet string `protobuf:"bytes,11,opt,name=remoteSliceGwNsmSubnet,proto3" json:"remoteSliceGwNsmSubnet,omitempty"`
	// Remote slice gateway Node IP
	RemoteSliceGwNodeIP string `protobuf:"bytes,12,opt,name=remoteSliceGwNodeIP,proto3" json:"remoteSliceGwNodeIP,omitempty"`
	// Remote slice gateway Node Port
	RemoteSliceGwNodePort string `protobuf:"bytes,13,opt,name=remoteSliceGwNodePort,proto3" json:"remoteSliceGwNodePort,omitempty"`
}

func (x *SliceGwConnectionContext) Reset() {
	*x = SliceGwConnectionContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliceGwConnectionContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceGwConnectionContext) ProtoMessage() {}

func (x *SliceGwConnectionContext) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceGwConnectionContext.ProtoReflect.Descriptor instead.
func (*SliceGwConnectionContext) Descriptor() ([]byte, []int) {
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP(), []int{5}
}

func (x *SliceGwConnectionContext) GetSliceId() string {
	if x != nil {
		return x.SliceId
	}
	return ""
}

func (x *SliceGwConnectionContext) GetLocalSliceGwId() string {
	if x != nil {
		return x.LocalSliceGwId
	}
	return ""
}

func (x *SliceGwConnectionContext) GetLocalSliceGwVpnIP() string {
	if x != nil {
		return x.LocalSliceGwVpnIP
	}
	return ""
}

func (x *SliceGwConnectionContext) GetLocalSliceGwHostType() SliceGwHostType {
	if x != nil {
		return x.LocalSliceGwHostType
	}
	return SliceGwHostType_SLICE_GW_SERVER
}

func (x *SliceGwConnectionContext) GetLocalSliceGwNsmSubnet() string {
	if x != nil {
		return x.LocalSliceGwNsmSubnet
	}
	return ""
}

func (x *SliceGwConnectionContext) GetLocalSliceGwNodeIP() string {
	if x != nil {
		return x.LocalSliceGwNodeIP
	}
	return ""
}

func (x *SliceGwConnectionContext) GetLocalSliceGwNodePort() string {
	if x != nil {
		return x.LocalSliceGwNodePort
	}
	return ""
}

func (x *SliceGwConnectionContext) GetRemoteSliceGwId() string {
	if x != nil {
		return x.RemoteSliceGwId
	}
	return ""
}

func (x *SliceGwConnectionContext) GetRemoteSliceGwVpnIP() string {
	if x != nil {
		return x.RemoteSliceGwVpnIP
	}
	return ""
}

func (x *SliceGwConnectionContext) GetRemoteSliceGwHostType() SliceGwHostType {
	if x != nil {
		return x.RemoteSliceGwHostType
	}
	return SliceGwHostType_SLICE_GW_SERVER
}

func (x *SliceGwConnectionContext) GetRemoteSliceGwNsmSubnet() string {
	if x != nil {
		return x.RemoteSliceGwNsmSubnet
	}
	return ""
}

func (x *SliceGwConnectionContext) GetRemoteSliceGwNodeIP() string {
	if x != nil {
		return x.RemoteSliceGwNodeIP
	}
	return ""
}

func (x *SliceGwConnectionContext) GetRemoteSliceGwNodePort() string {
	if x != nil {
		return x.RemoteSliceGwNodePort
	}
	return ""
}

var File_pkg_sidecar_sidecarpb_gw_sidecar_proto protoreflect.FileDescriptor

var file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x73, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x70, 0x62, 0x2f, 0x67, 0x77, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f,
	0x0a, 0x0f, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22,
	0xc8, 0x02, 0x0a, 0x0f, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x51, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x71,
	0x6f, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x74, 0x63, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x54, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x74, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x09,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x77, 0x43, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x62, 0x77, 0x43, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c,
	0x62, 0x77, 0x47, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x62, 0x77, 0x47, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x73, 0x63, 0x70, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x73, 0x63, 0x70, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x15, 0x54,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x49, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x52, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x78, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x78, 0x52, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x78,
	0x52, 0x61, 0x74, 0x65, 0x22, 0x56, 0x0a, 0x12, 0x4e, 0x73, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x73,
	0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x73, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x73, 0x6d, 0x49, 0x50, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x73, 0x6d, 0x49, 0x50, 0x22, 0xf8, 0x01, 0x0a,
	0x0b, 0x47, 0x77, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x50, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50,
	0x6f, 0x64, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x50, 0x6f, 0x64, 0x49, 0x50, 0x12, 0x26, 0x0a, 0x0e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x42, 0x0a, 0x0c, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x6e, 0x73, 0x6d, 0x49, 0x6e, 0x74, 0x66, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x4e, 0x73, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x6e, 0x73, 0x6d, 0x49, 0x6e, 0x74,
	0x66, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xbc, 0x05, 0x0a, 0x18, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x47, 0x77, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x47, 0x77, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53,
	0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x56, 0x70, 0x6e, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x56,
	0x70, 0x6e, 0x49, 0x50, 0x12, 0x4c, 0x0a, 0x14, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x47, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x47, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65,
	0x47, 0x77, 0x4e, 0x73, 0x6d, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x4e,
	0x73, 0x6d, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65,
	0x47, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x12, 0x32, 0x0a, 0x14, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x47, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x49, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x47, 0x77, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x56, 0x70, 0x6e, 0x49, 0x50, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47,
	0x77, 0x56, 0x70, 0x6e, 0x49, 0x50, 0x12, 0x4e, 0x0a, 0x15, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x15, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x48, 0x6f,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x4e, 0x73, 0x6d, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6c,
	0x69, 0x63, 0x65, 0x47, 0x77, 0x4e, 0x73, 0x6d, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x30,
	0x0a, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x50, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50,
	0x12, 0x34, 0x0a, 0x15, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47,
	0x77, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x4e, 0x6f,
	0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x2a, 0x3b, 0x0a, 0x0f, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47,
	0x77, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x4c, 0x49,
	0x43, 0x45, 0x5f, 0x47, 0x57, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x4c, 0x49, 0x43, 0x45, 0x5f, 0x47, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x2a, 0x1f, 0x0a, 0x06, 0x54, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a,
	0x11, 0x42, 0x41, 0x4e, 0x44, 0x57, 0x49, 0x44, 0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52,
	0x4f, 0x4c, 0x10, 0x00, 0x2a, 0x1d, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x54, 0x42, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x42,
	0x46, 0x10, 0x01, 0x32, 0xf8, 0x01, 0x0a, 0x10, 0x47, 0x77, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e,
	0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x47, 0x77, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x21, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65,
	0x47, 0x77, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x53, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4d, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x51, 0x6f,
	0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x51, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x53, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x3b, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescOnce sync.Once
	file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescData = file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDesc
)

func file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescGZIP() []byte {
	file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescOnce.Do(func() {
		file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescData)
	})
	return file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDescData
}

var file_pkg_sidecar_sidecarpb_gw_sidecar_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_sidecar_sidecarpb_gw_sidecar_proto_goTypes = []interface{}{
	(SliceGwHostType)(0),             // 0: sidecar.SliceGwHostType
	(TcType)(0),                      // 1: sidecar.TcType
	(ClassType)(0),                   // 2: sidecar.ClassType
	(*SidecarResponse)(nil),          // 3: sidecar.SidecarResponse
	(*SliceQosProfile)(nil),          // 4: sidecar.SliceQosProfile
	(*TunnelInterfaceStatus)(nil),    // 5: sidecar.TunnelInterfaceStatus
	(*NsmInterfaceStatus)(nil),       // 6: sidecar.NsmInterfaceStatus
	(*GwPodStatus)(nil),              // 7: sidecar.GwPodStatus
	(*SliceGwConnectionContext)(nil), // 8: sidecar.SliceGwConnectionContext
	(*empty.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_pkg_sidecar_sidecarpb_gw_sidecar_proto_depIdxs = []int32{
	1, // 0: sidecar.SliceQosProfile.tcType:type_name -> sidecar.TcType
	2, // 1: sidecar.SliceQosProfile.ClassType:type_name -> sidecar.ClassType
	5, // 2: sidecar.GwPodStatus.tunnelStatus:type_name -> sidecar.TunnelInterfaceStatus
	6, // 3: sidecar.GwPodStatus.nsmIntfStatus:type_name -> sidecar.NsmInterfaceStatus
	0, // 4: sidecar.SliceGwConnectionContext.localSliceGwHostType:type_name -> sidecar.SliceGwHostType
	0, // 5: sidecar.SliceGwConnectionContext.remoteSliceGwHostType:type_name -> sidecar.SliceGwHostType
	9, // 6: sidecar.GwSidecarService.GetStatus:input_type -> google.protobuf.Empty
	8, // 7: sidecar.GwSidecarService.UpdateConnectionContext:input_type -> sidecar.SliceGwConnectionContext
	4, // 8: sidecar.GwSidecarService.UpdateSliceQosProfile:input_type -> sidecar.SliceQosProfile
	7, // 9: sidecar.GwSidecarService.GetStatus:output_type -> sidecar.GwPodStatus
	3, // 10: sidecar.GwSidecarService.UpdateConnectionContext:output_type -> sidecar.SidecarResponse
	3, // 11: sidecar.GwSidecarService.UpdateSliceQosProfile:output_type -> sidecar.SidecarResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_sidecar_sidecarpb_gw_sidecar_proto_init() }
func file_pkg_sidecar_sidecarpb_gw_sidecar_proto_init() {
	if File_pkg_sidecar_sidecarpb_gw_sidecar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SidecarResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliceQosProfile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelInterfaceStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NsmInterfaceStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GwPodStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliceGwConnectionContext); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_sidecar_sidecarpb_gw_sidecar_proto_goTypes,
		DependencyIndexes: file_pkg_sidecar_sidecarpb_gw_sidecar_proto_depIdxs,
		EnumInfos:         file_pkg_sidecar_sidecarpb_gw_sidecar_proto_enumTypes,
		MessageInfos:      file_pkg_sidecar_sidecarpb_gw_sidecar_proto_msgTypes,
	}.Build()
	File_pkg_sidecar_sidecarpb_gw_sidecar_proto = out.File
	file_pkg_sidecar_sidecarpb_gw_sidecar_proto_rawDesc = nil
	file_pkg_sidecar_sidecarpb_gw_sidecar_proto_goTypes = nil
	file_pkg_sidecar_sidecarpb_gw_sidecar_proto_depIdxs = nil
}
