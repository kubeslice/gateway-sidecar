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

package bootstrap

import (
	"os"
	"os/exec"

	"github.com/kubeslice/gateway-sidecar/pkg/logger"
)

const (
	SERVER string = "SERVER"
	CLIENT string = "CLIENT"
)

type GatewayPod struct {
	mode            string
	mountPath       string
	secretMountPath string
	logger          *logger.Logger
}

func NewGatewayPod(mode string, mountPath, secretMountPath string, logger *logger.Logger) *GatewayPod {
	return &GatewayPod{
		mode:            mode,
		mountPath:       mountPath,
		secretMountPath: secretMountPath,
		logger:          logger,
	}
}

// Process() creates a directory structure as required by openvpn pods
func (gw *GatewayPod) Process() error {
	if gw.mode == SERVER {
		//copy the files from /var/run/vpn/. to /config/
		source := gw.secretMountPath + "."
		dest := "config/"
		cmd := exec.Command("cp", "-R", source, dest)
		stdout, err := cmd.Output()
		if err != nil {
			return err
		}
		gw.logger.Info(string(stdout))

		err = writeFile("config/ovpn_env.sh")
		if err != nil {
			return err
		}

	}
	return nil
}

func writeFile(source string) error {
	f, err := os.Create(source)
	if err != nil {
		return err
	}

	contents := `declare -x OVPN_DEFROUTE=0
declare -x OVPN_DEVICE=tun
declare -x OVPN_NAT=0`

	_, err = f.WriteString(contents)
	if err != nil {
		return err
	}
	defer f.Close()
	return err
}
