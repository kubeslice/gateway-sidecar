package main

import (
	"github.com/lorenzosaino/go-sysctl"
)

func main(){
	val,err := sysctl.Get("net.ipv4.ip_forward")
	if err != nil {
		
	}
	if val != "1" {
		// Set value of a net.ipv4.ip_forward to 1 using sysctl
		err = sysctl.Set("net.ipv4.ip_forward", "1")
		if err != nil {
			
		}
	}
}