/*
##########################################################

slicegateway_handler.go

Avesha LLC
February 2021

Copyright (c) Avesha LLC. 2021

Avesha Slice Gateway sidecar related routines

##########################################################
*/

package sidecar

import (
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/logger"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

var (
	sliceGwTc                         = TcInfo{}
	interClusterDscpClass      string = "Default"
	dscpClsStringToDscpCodeMap        = map[string]string{
		"QOS_PROFILE_DSCP_CLASS_DEFAULT": "Default",
		"QOS_PROFILE_DSCP_CLASS_AF11":    "AF11",
		"QOS_PROFILE_DSCP_CLASS_AF12":    "AF12",
		"QOS_PROFILE_DSCP_CLASS_AF13":    "AF13",
		"QOS_PROFILE_DSCP_CLASS_AF21":    "AF21",
		"QOS_PROFILE_DSCP_CLASS_AF22":    "AF22",
		"QOS_PROFILE_DSCP_CLASS_AF23":    "AF23",
		"QOS_PROFILE_DSCP_CLASS_AF31":    "AF31",
		"QOS_PROFILE_DSCP_CLASS_AF32":    "AF32",
		"QOS_PROFILE_DSCP_CLASS_AF33":    "AF33",
		"QOS_PROFILE_DSCP_CLASS_AF41":    "AF41",
		"QOS_PROFILE_DSCP_CLASS_AF42":    "AF42",
		"QOS_PROFILE_DSCP_CLASS_AF43":    "AF43",
		"QOS_PROFILE_DSCP_CLASS_EF":      "EF",
	}
	log = logger.NewLogger()
)

func sliceGwSetInterClusterDscpConfig(dscpClass string) error {
	_, valid := dscpClsStringToDscpCodeMap[dscpClass]
	if !valid {
		log.Infof("Dscp class string is not valid: %v", dscpClass)
		dscpClass = "QOS_PROFILE_DSCP_CLASS_DEFAULT"
	}

	portFilter := ""
	if os.Getenv("OPEN_VPN_MODE") == "CLIENT" {
		if SliceGwRemoteClusterNodePort == "" {
			log.Infof("Waiting for remote cluster node port to set the dscp config")
			return nil
		}
		portFilter = "--destination-port " + SliceGwRemoteClusterNodePort
	} else {
		portFilter = "--source-port 11194"
	}

	if interClusterDscpClass == dscpClsStringToDscpCodeMap[dscpClass] {
		log.Infof("No change in DSCP marking needed: %v", interClusterDscpClass)
		return nil
	}
	// Delete existing DSCP config before adding a new one
	if interClusterDscpClass != "Default" {
		ipTablesCmd := fmt.Sprintf("iptables -t mangle -D POSTROUTING -p udp %s -j DSCP --set-dscp-class %s",
			portFilter, interClusterDscpClass)
		_, err := runCommand(ipTablesCmd)
		if err != nil {
			log.Errorf("Could not remove existing DSCP config: %v. DSCP class in use: %v", err, interClusterDscpClass)
			return err
		}
	}

	ipTablesCmd := fmt.Sprintf("iptables -t mangle -A POSTROUTING -p udp %s -j DSCP --set-dscp-class %s",
		portFilter, dscpClsStringToDscpCodeMap[dscpClass])
	_, err := runCommand(ipTablesCmd)
	if err != nil {
		log.Errorf("DSCP marking failed: %v. DSCP class in use: %v", err, interClusterDscpClass)
	} else {
		log.Infof("Updating DSCP marking from %v to %v", interClusterDscpClass, dscpClsStringToDscpCodeMap[dscpClass])
		interClusterDscpClass = dscpClsStringToDscpCodeMap[dscpClass]
	}

	return err
}

func sliceGwGetInterClusterDscpConfig() (string, error) {
	ipTablesCmd := "iptables -t mangle -n -L POSTROUTING"
	return runCommand(ipTablesCmd)
}

func (s *GwSidecar) enforceSliceGwTc(newTc TcInfo) error {
	if sliceGwTc == newTc {
		log.Infof("No change in TC params, ignoring update")
		return nil
	} else {
		log.Info("TC params updated. Old: %v, New: %v", sliceGwTc, newTc)
		tcCmd := fmt.Sprintf("tc qdisc delete dev tun0 root tbf rate %dkbit burst 32kbit latency 500ms", sliceGwTc.bwCeiling)
		_, err := runTcCommand(tcCmd)
		if err != nil {
			return status.Errorf(codes.Internal, "tc command %v execution failed: %v", tcCmd, err)
		}
	}

	// Add follow TC command
	// tc qdisc add dev tun0 root tbf rate 5mbit burst 32kbit latency 500ms
	tcCmd := fmt.Sprintf("tc qdisc add dev tun0 root tbf rate %dkbit burst 32kbit latency 500ms", newTc.bwCeiling)
	cmdOut, err := runTcCommand(tcCmd)
	if err != nil {
		return status.Errorf(codes.Internal, "tc command %v execution failed: %v", tcCmd, err)
	}
	sliceGwTc = newTc
	log.Infof("tc Command %v output :%v", tcCmd, cmdOut)

	tcCmd = "tc qdisc show dev tun0"
	cmdOut, err = runTcCommand(tcCmd)
	log.Infof("tc Command %v output :%v", tcCmd, cmdOut)

	return nil
}

func (s *GwSidecar) enforceInterClusterQosPolicy(dscpClass string) error {
	err := sliceGwSetInterClusterDscpConfig(dscpClass)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to configure DSCP on inter cluster traffic: %v", err)
	}

	dscpConfig, err := sliceGwGetInterClusterDscpConfig()
	log.Infof("DSCP setting for inter cluster traffic: %v", dscpConfig)

	return nil
}

func (s *GwSidecar) enforceSliceQosPolicy(qosProfile *SliceQosProfile) error {
	err := s.enforceSliceGwTc(TcInfo{
		class:        classType(qosProfile.GetClassType().String()),
		bwCeiling:    qosProfile.BwCeiling,
		bwGuaranteed: qosProfile.BwGuaranteed,
		priority:     qosProfile.Priority,
	})
	if err != nil {
		log.Errorf("Failed to enforce TC settings on sliceGw. err: %v", err)
	}

	err = s.enforceInterClusterQosPolicy(qosProfile.DscpClass)
	if err != nil {
		log.Errorf("Failed to enforce Inter Cluster QoS policy. err: %v", err)
	}

	return nil
}
