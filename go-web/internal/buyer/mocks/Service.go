// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	buyer "mercado-frescos-time-7/go-web/internal/buyer"

	mock "github.com/stretchr/testify/mock"

	models "mercado-frescos-time-7/go-web/internal/models"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: card_number_id, first_name, last_name
func (_m *Service) Create(card_number_id string, first_name string, last_name string) (models.Buyer, error) {
	ret := _m.Called(card_number_id, first_name, last_name)

	var r0 models.Buyer
	if rf, ok := ret.Get(0).(func(string, string, string) models.Buyer); ok {
		r0 = rf(card_number_id, first_name, last_name)
	} else {
		r0 = ret.Get(0).(models.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(card_number_id, first_name, last_name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id int) error {
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
func (_m *Service) GetAll() (models.Buyers, error) {
	ret := _m.Called()

	var r0 models.Buyers
	if rf, ok := ret.Get(0).(func() models.Buyers); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Buyers)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetId provides a mock function with given fields: id
func (_m *Service) GetId(id int) (models.Buyer, error) {
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
func (_m *Service) Update(id int, body buyer.RequestPatch) (models.Buyer, error) {
	ret := _m.Called(id, body)

	var r0 models.Buyer
	if rf, ok := ret.Get(0).(func(int, buyer.RequestPatch) models.Buyer); ok {
		r0 = rf(id, body)
	} else {
		r0 = ret.Get(0).(models.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, buyer.RequestPatch) error); ok {
		r1 = rf(id, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
