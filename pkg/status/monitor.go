package status

import (
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/exec"
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/logger"
	"github.com/pkg/errors"
	"sync"
)

// Monitor shall hold the state of status monitor service
type Monitor struct {
	log          *logger.Logger
	regChecks    map[string]Check
	regExeModule map[string]*exec.Module
	lock         sync.RWMutex
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
	delete(m.regChecks, name)
	return nil
}

// DeregisterAll deregisters all the status checks.
func (m *Monitor) DeregisterAll() {
	for k := range m.regChecks {
		_ = m.Deregister(k)
	}
}
