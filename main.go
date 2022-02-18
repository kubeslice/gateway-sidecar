package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"github.com/lorenzosaino/go-sysctl"
	"google.golang.org/grpc"
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/bootstrap"
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/logger"
	sidecar "bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/sidecar/sidecarpb"
	
)

const (
	SECRET_MOUNT_PATH = "/var/run/vpn"
)

var (
	log *logger.Logger = logger.NewLogger()
)

// bootstrapGwPod shall bootstrap the Gateway Pod sidecar service.
// it creates the required directory structure for openvpn pods
func bootstrapGwPod(wg *sync.WaitGroup) error{
	gwPod := bootstrap.NewGatewayPod(os.Getenv("OPEN_VPN_MODE"),os.Getenv("MOUNT_PATH"),SECRET_MOUNT_PATH,log)

	if err:= gwPod.Process();err!=nil{
		log.Errorf("Error bootstraping gw pod",err.Error())
		return err
	}
	//TODO: register status checks
	wg.Done()
	log.Info("finished bootstraping gw pod")
	return nil
}

func startGrpcServer(grpcPort string) error {
	address := fmt.Sprintf(":%s",grpcPort)
	log.Infof("Starting GRPC Server for %v Pod at %v", "GW-Sidecar", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("Unable to connect to Server: %v", err.Error())
		return err
	}
	srv := grpc.NewServer()
	sidecar.RegisterGwSidecarServiceServer(srv,&sidecar.GwSidecar{})

	err = srv.Serve(lis)
	if err != nil {
		log.Errorf("Start GRPC Server Failed with %v", err.Error())
		return err
	}
	log.Infof("GRPC Server exited gracefully")

	return nil
}

// shutdownHandler triggers application shutdown.
func shutdownHandler(wg *sync.WaitGroup) {
	// signChan channel is used to transmit signal notifications.
	signChan := make(chan os.Signal, 1)
	// Catch and relay certain signal(s) to signChan channel.
	signal.Notify(signChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Blocking until a signal is sent over signChan channel. Progress to
	// next line after signal
	sig := <-signChan
	log.Infof("Teardown started with ", sig, "signal")
	
	wg.Done()
	//os.Exit(1)
}


func main(){
	var grpcPort string = "5000"
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
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go bootstrapGwPod(wg)

	// Start the GRPC Server to communicate with slice controller.
	go startGrpcServer(grpcPort)

	go shutdownHandler(wg)

	wg.Wait()
	log.Infof("Avesha Sidecar exited")
	
}