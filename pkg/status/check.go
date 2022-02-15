// check.go
//
// Avesha LLC
// Sept 2020
//
// Copyright (c) Avesha LLC. 2020
//
// Module: Avesha Sidecar - Status Module Check interface

package status

// Check is the interface for defining Status checks.
// A valid check has a non empty Execute() and a check Status() function.
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
