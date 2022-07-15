// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	employees "mercado-frescos-time-7/go-web/internal/employees"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: id, card_number_id, first_name, last_name, warehouse_id
func (_m *Repository) Create(id int, card_number_id string, first_name string, last_name string, warehouse_id int) (employees.Employee, error) {
	ret := _m.Called(id, card_number_id, first_name, last_name, warehouse_id)

	var r0 employees.Employee
	if rf, ok := ret.Get(0).(func(int, string, string, string, int) employees.Employee); ok {
		r0 = rf(id, card_number_id, first_name, last_name, warehouse_id)
	} else {
		r0 = ret.Get(0).(employees.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string, int) error); ok {
		r1 = rf(id, card_number_id, first_name, last_name, warehouse_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() ([]employees.Employee, error) {
	ret := _m.Called()

	var r0 []employees.Employee
	if rf, ok := ret.Get(0).(func() []employees.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]employees.Employee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id int) (employees.Employee, error) {
	ret := _m.Called(id)

	var r0 employees.Employee
	if rf, ok := ret.Get(0).(func(int) employees.Employee); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(employees.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastID provides a mock function with given fields:
func (_m *Repository) LastID() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: e, id
func (_m *Repository) Update(e employees.Employee, id int) (employees.Employee, error) {
	ret := _m.Called(e, id)

	var r0 employees.Employee
	if rf, ok := ret.Get(0).(func(employees.Employee, int) employees.Employee); ok {
		r0 = rf(e, id)
	} else {
		r0 = ret.Get(0).(employees.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(employees.Employee, int) error); ok {
		r1 = rf(e, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidationCardNumberID provides a mock function with given fields: card_number_id
func (_m *Repository) ValidationCardNumberID(card_number_id string) error {
	ret := _m.Called(card_number_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(card_number_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}