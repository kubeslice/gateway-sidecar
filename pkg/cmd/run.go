// run.go
//
// Avesha LLC
// Sept 2020
//
// Copyright (c) Avesha LLC. 2020
//
// Module: Avesha Sidecar - Command Run Helper Module

package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/google/shlex"
)

// runCommand runs the command string
func Run(cmdString string) (string, error) {
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
