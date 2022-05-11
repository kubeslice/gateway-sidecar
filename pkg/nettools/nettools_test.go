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
		fmt.Println(a.Name)
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
