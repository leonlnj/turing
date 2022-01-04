// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	json "encoding/json"

	manager "github.com/gojek/turing/engines/experiment/manager"
	mock "github.com/stretchr/testify/mock"
)

// ExperimentsService is an autogenerated mock type for the ExperimentsService type
type ExperimentsService struct {
	mock.Mock
}

// GetExperimentRunnerConfig provides a mock function with given fields: engine, cfg
func (_m *ExperimentsService) GetExperimentRunnerConfig(engine string, cfg json.RawMessage) (json.RawMessage, error) {
	ret := _m.Called(engine, cfg)

	var r0 json.RawMessage
	if rf, ok := ret.Get(0).(func(string, json.RawMessage) json.RawMessage); ok {
		r0 = rf(engine, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(json.RawMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, json.RawMessage) error); ok {
		r1 = rf(engine, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsStandardExperimentManager provides a mock function with given fields: engine
func (_m *ExperimentsService) IsStandardExperimentManager(engine string) bool {
	ret := _m.Called(engine)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(engine)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListClients provides a mock function with given fields: engine
func (_m *ExperimentsService) ListClients(engine string) ([]manager.Client, error) {
	ret := _m.Called(engine)

	var r0 []manager.Client
	if rf, ok := ret.Get(0).(func(string) []manager.Client); ok {
		r0 = rf(engine)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]manager.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(engine)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEngines provides a mock function with given fields:
func (_m *ExperimentsService) ListEngines() []manager.Engine {
	ret := _m.Called()

	var r0 []manager.Engine
	if rf, ok := ret.Get(0).(func() []manager.Engine); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]manager.Engine)
		}
	}

	return r0
}

// ListExperiments provides a mock function with given fields: engine, clientID
func (_m *ExperimentsService) ListExperiments(engine string, clientID string) ([]manager.Experiment, error) {
	ret := _m.Called(engine, clientID)

	var r0 []manager.Experiment
	if rf, ok := ret.Get(0).(func(string, string) []manager.Experiment); ok {
		r0 = rf(engine, clientID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]manager.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(engine, clientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVariables provides a mock function with given fields: engine, clientID, experimentIDs
func (_m *ExperimentsService) ListVariables(engine string, clientID string, experimentIDs []string) (manager.Variables, error) {
	ret := _m.Called(engine, clientID, experimentIDs)

	var r0 manager.Variables
	if rf, ok := ret.Get(0).(func(string, string, []string) manager.Variables); ok {
		r0 = rf(engine, clientID, experimentIDs)
	} else {
		r0 = ret.Get(0).(manager.Variables)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []string) error); ok {
		r1 = rf(engine, clientID, experimentIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateExperimentConfig provides a mock function with given fields: engine, cfg
func (_m *ExperimentsService) ValidateExperimentConfig(engine string, cfg json.RawMessage) error {
	ret := _m.Called(engine, cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, json.RawMessage) error); ok {
		r0 = rf(engine, cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
