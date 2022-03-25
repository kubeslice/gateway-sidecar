// exec.go
//
// Avesha LLC
// Sept 2020
//
// Copyright (c) Avesha LLC. 2020
//
// Module: Avesha Sidecar - Execution Module
package exec

import (
	"time"

	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/logger"
)

// RunFunc of function type
type RunFunc func(interface{}) error

// MsgHandler of function type
type MsgHandler func(interface{}) error

// Module defines the datastructure for a Module implementation.
type Module struct {
	log      *logger.Logger
	interval time.Duration
	exec     RunFunc
	exArgs   interface{}
	msgHdlr  MsgHandler
	rxCh     chan interface{}
	ticker   *time.Ticker
}

// NewModule is Constructor for Module Object.
func NewModule(log *logger.Logger, interval time.Duration, exec RunFunc, exArgs interface{}, msgHdlr MsgHandler) *Module {
	ticker := time.NewTicker(interval)
	return &Module{log, interval, exec, exArgs, msgHdlr, make(chan interface{}), ticker}
}

// The Module run local method
func (m *Module) run() {
	for {
		select {
		// case <-time.After(m.interval):
		case <-m.ticker.C:
			m.log.Debugf("Timer Expired  ", m.exArgs)
			_ = m.exec(m.exArgs)
		case msg, ok := <-m.rxCh:
			m.log.Debugf("msg Rxd ", msg)
			m.msgHdlr(msg)
			if !ok {
				m.log.Infof("Closing the Channel and break")
				return // channel was closed and drained
			}
		}
	}
}

// Start the Module
func (m *Module) Start() {
	go m.run()
}

// Stop the Module
func (m *Module) Stop() {
	m.log.Infof("Stop is called")
	m.ticker.Stop()
	close(m.rxCh)
}

// SendMsg the Module
func (m *Module) SendMsg(msg interface{}) {
	m.log.Debugf("SendMsg : %v", msg)
	m.rxCh <- msg
}
