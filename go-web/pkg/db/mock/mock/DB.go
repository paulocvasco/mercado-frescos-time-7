// Code generated by mockery v2.13.1. DO NOT EDIT.

package mockDB

import mock "github.com/stretchr/testify/mock"

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// Load provides a mock function with given fields: _a0
func (_m *DB) Load(_a0 interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: _a0
func (_m *DB) Save(_a0 interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDB(t mockConstructorTestingTNewDB) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
