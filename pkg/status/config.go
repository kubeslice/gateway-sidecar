package status

import "time"

// Config defines a health Check and it's scheduling timing requirements.
type Config struct {
	// Name of the status check
	Name string
	// Check is the status Check to be scheduled for execution.
	Checker Check
	// Interval is the time between successive executions.
	Interval time.Duration
	// InitialDelay is the time to delay first execution; defaults to zero.
	InitialDelay time.Duration
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
