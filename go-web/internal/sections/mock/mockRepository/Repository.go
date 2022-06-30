// Code generated by mockery v2.13.1. DO NOT EDIT.

package mockRepository

import (
	models "mercado-frescos-time-7/go-web/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *Repository) Delete(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() (models.Sections, error) {
	ret := _m.Called()

	var r0 models.Sections
	if rf, ok := ret.Get(0).(func() models.Sections); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Sections)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: _a0
func (_m *Repository) GetById(_a0 int) (models.Section, error) {
	ret := _m.Called(_a0)

	var r0 models.Section
	if rf, ok := ret.Get(0).(func(int) models.Section); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Section)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0
func (_m *Repository) Store(_a0 models.Section) (models.Section, error) {
	ret := _m.Called(_a0)

	var r0 models.Section
	if rf, ok := ret.Get(0).(func(models.Section) models.Section); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Section)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Section) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *Repository) Update(_a0 int, _a1 models.Section) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, models.Section) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateID provides a mock function with given fields: _a0
func (_m *Repository) ValidateID(_a0 int) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// VerifySectionNumber provides a mock function with given fields: _a0
func (_m *Repository) VerifySectionNumber(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
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
