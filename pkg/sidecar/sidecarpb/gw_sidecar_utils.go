package sidecar

import (
	"os"
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/nettools"
)

const (
	nsmInterfaceName string = "nsm0"
)

func getGwPodStatus() (*GwPodStatus,error) {
	podStatus := &GwPodStatus{}

	podStatus.GatewayPodIP = nettools.GetPodIP()
	podStatus.NodeIP = os.Getenv("NODE_IP")
	podNsmIP, err := nettools.GetInterfaceIP(nsmInterfaceName)
	if err != nil {
		podNsmIP = ""
	}

	nsmIntfStatus := NsmInterfaceStatus{
		NsmInterfaceName: nsmInterfaceName,
		NsmIP:            podNsmIP,
	}

	podStatus.NsmIntfStatus = &nsmIntfStatus

	//TODO:remaining part - tunnelStatus
	
	return podStatus,nil

}