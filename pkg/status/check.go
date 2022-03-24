// check.go
//
// Avesha LLC
// Feb 2022
//
// Copyright (c) Avesha LLC. 2022
//
// Module: Gateway Sidecar - Status Moduel Check Interface

package status

type Check interface {
	// Execute runs a single time check, and returns execution status.
	Execute(interface{}) (err error)
	// MessageHandler handles the message and returns message handling status.
	MessageHandler(msg interface{}) (err error)
	// Status to provide the status of the check
	Status() (details interface{}, err error)
	// Stop the status of the check
	Stop() (err error)
}
