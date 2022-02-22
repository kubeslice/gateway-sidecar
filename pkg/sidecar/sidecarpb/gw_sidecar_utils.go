package sidecar

import (
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/nettools"
	"bytes"
	"errors"
	"fmt"
	"github.com/google/shlex"
	"os"
	"os/exec"
	"strings"
)

const (
	nsmInterfaceName string = "nsm0"
)

func getGwPodStatus() (*GwPodStatus, error) {
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

	return podStatus, nil

}

// runCommand runs the command string
func runCommand(cmdString string) (string, error) {
	var outb, errb bytes.Buffer

	ss, err := shlex.Split(cmdString)
	if err != nil {
		errMsg := fmt.Sprintf("Command split failed with error : %v", err)
		return "", errors.New(errMsg)
	}
	if len(ss) == 0 {
		errMsg := fmt.Sprintf("No command defined : %v", cmdString)
		return "", errors.New(errMsg)
	}
	cmd := exec.Command(ss[0], ss[1:]...)
	if err != nil {
		errMsg := fmt.Sprintf("Command construction failed with error : %v", err)
		return "", errors.New(errMsg)
	}
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	// Run the command
	err = cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("Could not run cmd: %v", err)
		return errb.String(), errors.New(errMsg)

	}
	return outb.String(), nil
}

func runTcCommand(tcCmd string) (string, error) {
	var errVal error = nil
	var err error = nil
	var cmdOut string = ""
	cmdOut, err = runCommand(tcCmd)
	if err != nil {
		errStr := fmt.Sprintf("tc Command: %v execution failed with err: %v and stderr : %v", tcCmd, err, cmdOut)
		log.Errorf(errStr)

		if strings.Contains(cmdOut, "RTNETLINK answers: File exists") {
			tcDelCmd := strings.Replace(tcCmd, "add", "del", -1)
			cmdOut, err = runCommand(tcDelCmd)
			if err != nil {
				errStr := fmt.Sprintf("tc Command: %v execution failed with err: %v and stderr : %v", tcDelCmd, err, cmdOut)
				log.Errorf(errStr)
				errVal = errors.New(errStr)
			}
			log.Debugf("tc Command: %v output :%v", tcDelCmd, cmdOut)

			// Re run the tc command
			cmdOut, err = runCommand(tcCmd)
			if err != nil {
				errStr := fmt.Sprintf("tc Command: %v execution failed with err: %v and stderr : %v", tcCmd, err, cmdOut)
				errVal = errors.New(errStr)
			}
			log.Infof("tc Command: %v output :%v", tcCmd, cmdOut)
		}
	}
	return cmdOut, errVal
}
