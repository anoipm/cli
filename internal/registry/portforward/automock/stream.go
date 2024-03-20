// Code generated by mockery v2.33.3. DO NOT EDIT.

package automock

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Stream is an autogenerated mock type for the Stream type
type Stream struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Stream) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Headers provides a mock function with given fields:
func (_m *Stream) Headers() http.Header {
	ret := _m.Called()

	var r0 http.Header
	if rf, ok := ret.Get(0).(func() http.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Header)
		}
	}

	return r0
}

// Identifier provides a mock function with given fields:
func (_m *Stream) Identifier() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// Read provides a mock function with given fields: p
func (_m *Stream) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reset provides a mock function with given fields:
func (_m *Stream) Reset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: p
func (_m *Stream) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStream creates a new instance of Stream. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStream(t interface {
	mock.TestingT
	Cleanup(func())
}) *Stream {
	mock := &Stream{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
