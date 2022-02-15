// monitor.go
//
// Avesha LLC
// Sept 2020
//
// Copyright (c) Avesha LLC. 2020
//
// Module: Avesha Sidecar - Status Module's Monitor

package status

import (
	"sync"
	"time"

	"github.com/pkg/errors"

	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/exec"
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/logger"
)

// Types defines possible status types.
type Types uint

// Constant to define the StatusCheck Type
const (
	TunnelCheck Types = 1 << iota // Tunnel Status Check
)

var typeNames = []string{
	"TunnelCheck",
}

func (f Types) String() string {
	s := ""
	for i, name := range typeNames {
		if f&(1<<uint(i)) != 0 {
			if s != "" {
				s += "|"
			}
			s += name
		}
	}
	if s == "" {
		s = "0"
	}
	return s
}

// IMonitor is the API for registering / deregistering status checks.
type IMonitor interface {
	// RegisterCheck registers a status  check according to the given configuration.
	// Once RegisterCheck() is called, the check is scheduled to run in it's own goroutine.
	RegisterCheck(cfg *Config) error
	// Deregister removes a health check from this instance, and stops it's next executions.
	// If the check is running while Deregister() is called, the check may complete it's current execution.
	// Once a check is removed, it's results are no longer returned.
	Deregister(name string) error
	// DeregisterAll Deregister removes all health checks from this instance, and stops their next executions.
	// It is equivalent of calling Deregister() for each currently registered check.
	DeregisterAll()
	Checks() map[string]Check
}

// Monitor shall hold the state of status monitor service
type Monitor struct {
	log          *logger.Logger
	regChecks    map[string]Check
	regExeModule map[string]*exec.Module
	lock sync.RWMutex
}

// Config defines a health Check and it's scheduling timing requirements.
type Config struct {
	// Name of the status check
	Name string
	// Check is the status Check to be scheduled for execution.
	Checker Check
	// Interval is the time between successive executions.
	Interval time.Duration
	// InitialDelay is the time to delay first execution; defaults to zero.
	InitialDelay time.Duration
}


// NewMonitor creates a new Status Monitor.
func NewMonitor(log *logger.Logger) *Monitor {
	return &Monitor{
		log:          log,
		regChecks:    make(map[string]Check),
		regExeModule: make(map[string]*exec.Module),
	}
}

// RegisterCheck registers the status check
func (m *Monitor) RegisterCheck(cfg *Config) (*exec.Module, error) {
	if cfg.Checker == nil || cfg.Name == "" {
		return nil, errors.Errorf("Invalid status check %v", cfg.Checker)
	}

	m.lock.Lock()
	defer m.lock.Unlock()
	_, ok := m.regChecks[cfg.Name]
	if ok {
		return nil, errors.New("Check already exists: " + cfg.Name)
	}
	m.regChecks[cfg.Name] = cfg.Checker
	m.regExeModule[cfg.Name] = exec.NewModule(m.log, cfg.Interval, cfg.Checker.Execute, nil, cfg.Checker.MessageHandler)
	m.regExeModule[cfg.Name].Start()
	return m.regExeModule[cfg.Name], nil
}

// Checks provides the available status checks.
func (m *Monitor) Checks() map[string]Check {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.regChecks
}

// Deregister unregisters the health check.
func (m *Monitor) Deregister(name string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.log.Infof("Deregister %v Module", name)

	if m.regExeModule[name] != nil {
		m.regExeModule[name].Stop()
	}
	return nil
}

// DeregisterAll deregisters all the status checks.
func (m *Monitor) DeregisterAll() {
	for k := range m.regChecks {
		_ = m.Deregister(k)
	}
}

// Local function for the status check callback.
func (*Monitor) statusChecker(args interface{}) error {
	cfg, _ := args.(*Config)

	err := cfg.Checker.Execute(nil)
	if err != nil {
		return err
	}
	return nil
}
