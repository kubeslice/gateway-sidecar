// response.go
//
// Avesha LLC
// Sept 2020
//
// Copyright (c) Avesha LLC. 2020
//
// Module: Avesha Sidecar - Status Module - Pod status to Backend.

package status

// Direction - Custom type to hold value for week day ranging from 1-4
type PodStatus int

// Declare related constants for each direction starting with index 1
const (
	POD_STATUS_INVALID      PodStatus = iota // EnumIndex = 0
	POD_STATUS_INITIALIZING                  // EnumIndex = 1
	POD_STATUS_HEALTHY                       // EnumIndex = 2
	POD_STATUS_UNHEALTHY                     // EnumIndex = 3
	POD_STATUS_TERMINATED                    // EnumIndex = 4
)

// String - Creating common behavior - give the type a String function
func (d PodStatus) String() string {
	return [...]string{"POD_STATUS_INVALID", "POD_STATUS_INITIALIZING", "POD_STATUS_HEALTHY", "POD_STATUS_UNHEALTHY"}[d]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (d PodStatus) EnumIndex() int {
	return int(d)
}

// AppPodStatus represents Application Pod Status.
type AppPodStatus struct {
	PodName        string `json:"podName,omitempty"`
	NsmInterface   string `json:"nsmInterface,omitempty"`
	NsmIP          string `json:"nsmIp,omitempty"`
	PodIP          string `json:"podIp,omitempty"`
	NsmGwInterface string `json:"nsmGwInterface,omitempty"`
	NsmGwIP        string `json:"nsmGwIp,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
}

// TunnelInterfaceStatus represents Tunnel Interface Status
type TunnelInterfaceStatus struct {
	NetInterface string `json:"netInterface,omitempty"`
	LocalIP      string `json:"localIp,omitempty"`
	PeerIP       string `json:"peerIp,omitempty"`
	Latency      uint64 `json:"latency,omitempty"`
	TxRate       uint64 `json:"txRate,omitempty"`
	RxRate       uint64 `json:"rxRate,omitempty"`
}

// SliceGatewayHealth represents generic details of the Pod.
type SliceGatewayHealth struct {
	NodeIP       string                `json:"nodeIp,omitempty"`
	GatewayPodIP string                `json:"gatewayPodIp,omitempty"`
	PodStatus    PodStatus             `json:"podStatus,omitempty"`
	TunnelStatus TunnelInterfaceStatus `json:"tunnelStatus,omitempty"`
	AppStatus    []AppPodStatus        `json:"appStatus,omitempty"`
}

// Response represents overall status of the Pod
type Response struct {
	Health SliceGatewayHealth `json:"health,omitempty"`
}
