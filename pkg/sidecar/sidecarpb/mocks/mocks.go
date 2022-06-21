package mocks

import mock "github.com/stretchr/testify/mock"

type GwSidecarProvider struct {
	mock.Mock
}

// CheckIfVppIntfPresent provides a mock function with given fields:
func (_m *GwSidecarProvider) CheckIfVppIntfPresent() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewGwSidecarProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewGwSidecarProvider creates a new instance of GwSidecarProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGwSidecarProvider(t mockConstructorTestingTNewGwSidecarProvider) *GwSidecarProvider {
	mock := &GwSidecarProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
