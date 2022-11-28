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

package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kubeslice/gateway-sidecar/pkg/bootstrap"
	"github.com/kubeslice/gateway-sidecar/pkg/logger"
	"github.com/kubeslice/gateway-sidecar/pkg/metrics"
	"github.com/kubeslice/gateway-sidecar/pkg/nettools"
	sidecar "github.com/kubeslice/gateway-sidecar/pkg/sidecar/sidecarpb"
	"github.com/kubeslice/gateway-sidecar/pkg/status"
	"github.com/lorenzosaino/go-sysctl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	SECRET_MOUNT_PATH = "/var/run/vpn/"
)

var (
	log *logger.Logger = logger.NewLogger()
)

// bootstrapGwPod shall bootstrap the Gateway Pod sidecar service.
// it creates the required directory structure for openvpn pods
func bootstrapGwPod(wg *sync.WaitGroup) error {
	var tunCheck *status.TunnelChecker
	gwPod := bootstrap.NewGatewayPod(os.Getenv("OPEN_VPN_MODE"), os.Getenv("MOUNT_PATH"), SECRET_MOUNT_PATH, log)
	if err := gwPod.Process(); err != nil {
		log.Errorf("Error bootstraping gw pod", err.Error())
		return err
	}

	statusMonitor := status.NewMonitor(log)
	tunCheck = status.NewTunnelChecker(log).(*status.TunnelChecker)
	mod, err := statusMonitor.RegisterCheck(&status.Config{
		Name:     "TunnelCheck",
		Checker:  tunCheck,
		Interval: time.Second * 6,
	})
	if err != nil {
		log.Fatalf("Registering Tunnel check failed with Error : %v", err)
	}
	sidecar.SetStatusMonitor(statusMonitor)
	tunCheck.UpdateExecModule(mod)
	wg.Done()
	log.Info("finished bootstraping gw pod")
	return nil
}

func startGrpcServer(grpcPort string, wg *sync.WaitGroup) error {
	address := fmt.Sprintf("0.0.0.0:%s", grpcPort)
	log.Infof("Starting GRPC Server for %v Pod at %v", "GW-Sidecar", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("Unable to connect to Server: %v", err.Error())
		return err
	}
	srv := grpc.NewServer()
	sidecar.RegisterGwSidecarServiceServer(srv, &sidecar.GwSidecar{})

	// TODO: only for debug purpose please revert
	reflection.Register(srv)

	err = srv.Serve(lis)
	if err != nil {
		log.Errorf("Start GRPC Server Failed with %v", err.Error())
		return err
	}
	wg.Done()
	log.Infof("GRPC Server exited gracefully")

	return nil
}

func getTun0PerrIp() (string, error) {
	time.Sleep(30 * time.Second)
	// Get the tun0 interface IP address
	infIp, err := nettools.GetInterfaceIP("tun0")
	fmt.Println("recieved tun0 from getTun0PerrIP", infIp)
	if infIp == "" {
		fmt.Println("unable to get tun0 IP")
	}
	if err != nil {
		log.Fatalf("Error getting the interface IP address", err)
	}
	address := infIp + ":5000"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	defer conn.Close()
	client := sidecar.NewGwSidecarServiceClient(conn)

	res, err := client.GetStatus(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("res from func", res)
	if res.TunnelStatus.PeerIP == "" {
		time.Sleep(10 * time.Second)
	}
	return res.TunnelStatus.PeerIP, nil
}

func startGrpcClient(grpcPort string) error {
	ip, err := getTun0PerrIp()
	fmt.Println("preeIP", ip)
	if err != nil {
		log.Fatalf("Error getting the interface IP address", err)
	}
	address := ip + ":5000"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	defer conn.Close()
	client := sidecar.NewGwSidecarServiceClient(conn)

	res, err := client.GetStatus(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("res:", res)
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

func main() {
	var grpcPort string = "5000"
	var metricCollectorPort string = "18080"
	// Get value of a net.ipv4.ip_forward using sysctl
	val, err := sysctl.Get("net.ipv4.ip_forward")
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
	wg.Add(3)

	go bootstrapGwPod(wg)

	// Start the GRPC Server to communicate with slice controller.
	go startGrpcServer(grpcPort, wg)

	// Start the GRPC Client to fetch info from remote gw pod
	go startGrpcClient(grpcPort)

	go metrics.StartMetricsCollector(metricCollectorPort)

	go shutdownHandler(wg)

	wg.Wait()
	log.Infof("Gateway Sidecar exited")
}
