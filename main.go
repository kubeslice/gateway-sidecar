package main

import (
	"os"
	"github.com/lorenzosaino/go-sysctl"
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/logger"
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/bootstrap"
)

// bootstrapGwPod shall bootstrap the Gateway Pod sidecar service.
// it creates the required directory structure for openvpn pods  
func bootstrapGwPod(log *logger.Logger) error{
	gwPod := bootstrap.NewGatewayPod(os.Getenv("OPEN_VPN_MODE"),os.Getenv("MOUNT_PATH"),os.Getenv("SECRET_MOUNT_PATH"),log)

	if err:= gwPod.Process();err!=nil{
		return err
	}
	return nil
}

func main(){
	var logLevel,logPath string
	logLevel = os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO"
	}
	if logPath == "" {
		logPath = "avesha-sidecar.log"
	}
	//create a new logger module
	log := logger.NewLogger(logLevel,logPath)

	// Get value of a net.ipv4.ip_forward using sysctl
	val,err := sysctl.Get("net.ipv4.ip_forward")
	if err != nil {
		log.Fatalf("Retrive of ipv4.ip_forward errored %v", err)
	}
	if val != "1" {
		// Set value of a net.ipv4.ip_forward to 1 using sysctl
		err = sysctl.Set("net.ipv4.ip_forward", "1")
		if err != nil {
			log.Fatalf("Set of ipv4.ip_forward errored %v", err)
		}
	}
	go bootstrapGwPod(log)

	
}