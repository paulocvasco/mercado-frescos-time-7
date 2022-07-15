// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	models "mercado-frescos-time-7/go-web/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryFile is an autogenerated mock type for the RepositoryFile type
type RepositoryFile struct {
	mock.Mock
}

// Create provides a mock function with given fields: CardNumberID, FirstName, LastName
func (_m *RepositoryFile) Create(CardNumberID string, FirstName string, LastName string) (models.Buyer, error) {
	ret := _m.Called(CardNumberID, FirstName, LastName)

	var r0 models.Buyer
	if rf, ok := ret.Get(0).(func(string, string, string) models.Buyer); ok {
		r0 = rf(CardNumberID, FirstName, LastName)
	} else {
		r0 = ret.Get(0).(models.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(CardNumberID, FirstName, LastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *RepositoryFile) Delete(id int) error {
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
func (_m *RepositoryFile) GetAll() ([]models.Buyer, error) {
	ret := _m.Called()

	var r0 []models.Buyer
	if rf, ok := ret.Get(0).(func() []models.Buyer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Buyer)
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

// GetCardNumberId provides a mock function with given fields: id
func (_m *RepositoryFile) GetCardNumberId(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetId provides a mock function with given fields: id
func (_m *RepositoryFile) GetId(id int) (models.Buyer, error) {
	ret := _m.Called(id)

	var r0 models.Buyer
	if rf, ok := ret.Get(0).(func(int) models.Buyer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, body
func (_m *RepositoryFile) Update(id int, body models.Buyer) (models.Buyer, error) {
	ret := _m.Called(id, body)

	var r0 models.Buyer
	if rf, ok := ret.Get(0).(func(int, models.Buyer) models.Buyer); ok {
		r0 = rf(id, body)
	} else {
		r0 = ret.Get(0).(models.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, models.Buyer) error); ok {
		r1 = rf(id, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositoryFile interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryFile creates a new instance of RepositoryFile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryFile(t mockConstructorTestingTNewRepositoryFile) *RepositoryFile {
	mock := &RepositoryFile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
