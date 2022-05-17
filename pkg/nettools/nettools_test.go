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

package nettools

import (
	"fmt"
	"net"
	"runtime/debug"
	"testing"
)

var err error

func getTheIpAndName() (string, string) {

	interfaceNames := make([]string, 2)

	allInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err while getting all the interfaces")
	}
	fmt.Println(allInterfaces)
	for _, a := range allInterfaces {
		interfaceNames = append(interfaceNames, a.Name)
	}

	infos, err := GetInterfaceInfos(interfaceNames[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	return infos[0].Name, infos[0].IP
}

func TestGetPodIP(t *testing.T) {

	_, InfIP := getTheIpAndName()

	tests := []struct {
		testName string
		addrs    []net.Addr
		expected string
		actual   string
		errMsg   string
	}{
		{
			"Sucessfully got the PodIP",
			[]net.Addr{},
			InfIP,
			"",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			tt.actual = GetPodIP()

			if tt.actual != tt.expected {
				t.Fail()
			}
		})
	}
}

func TestGetInterfaceIps(t *testing.T) {

	InfName, InfIP := getTheIpAndName()

	tests := []struct {
		testName string
		expected string
		actual   string
		errMsg   string
	}{
		{
			"It should pass",
			InfIP,
			"",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			tt.actual, err = GetInterfaceIP(InfName)
			AssertNoError(t, err)

			if tt.actual != tt.expected {
				t.Fail()
			}
		})
	}

}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected No Error but got %s, Stack:\n%s", err, string(debug.Stack()))
	}
}
