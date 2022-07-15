// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	models "mercado-frescos-time-7/go-web/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
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
func (_m *Service) GetAll() ([]models.Seller, error) {
	ret := _m.Called()

	var r0 []models.Seller
	if rf, ok := ret.Get(0).(func() []models.Seller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Seller)
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

// GetId provides a mock function with given fields: indice
func (_m *Service) GetId(indice int) (models.Seller, error) {
	ret := _m.Called(indice)

	var r0 models.Seller
	if rf, ok := ret.Get(0).(func(int) models.Seller); ok {
		r0 = rf(indice)
	} else {
		r0 = ret.Get(0).(models.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(indice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: sel
func (_m *Service) Store(sel models.Seller) (models.Seller, error) {
	ret := _m.Called(sel)

	var r0 models.Seller
	if rf, ok := ret.Get(0).(func(models.Seller) models.Seller); ok {
		r0 = rf(sel)
	} else {
		r0 = ret.Get(0).(models.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Seller) error); ok {
		r1 = rf(sel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: s, id
func (_m *Service) Update(s []byte, id int) (models.Seller, error) {
	ret := _m.Called(s, id)

	var r0 models.Seller
	if rf, ok := ret.Get(0).(func([]byte, int) models.Seller); ok {
		r0 = rf(s, id)
	} else {
		r0 = ret.Get(0).(models.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, int) error); ok {
		r1 = rf(s, id)
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
