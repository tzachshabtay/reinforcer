// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	generator "github.com/csueiras/reinforcer/internal/generator"
	executor "github.com/csueiras/reinforcer/internal/generator/executor"

	mock "github.com/stretchr/testify/mock"
)

// Executor is an autogenerated mock type for the Executor type
type Executor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: settings
func (_m *Executor) Execute(settings *executor.Parameters) (*generator.Generated, error) {
	ret := _m.Called(settings)

	var r0 *generator.Generated
	if rf, ok := ret.Get(0).(func(*executor.Parameters) *generator.Generated); ok {
		r0 = rf(settings)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*generator.Generated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*executor.Parameters) error); ok {
		r1 = rf(settings)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
