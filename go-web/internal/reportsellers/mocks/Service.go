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

// ReportSellers provides a mock function with given fields: id
func (_m *Service) ReportSellers(id int) ([]models.ReportSeller, error) {
	ret := _m.Called(id)

	var r0 []models.ReportSeller
	if rf, ok := ret.Get(0).(func(int) []models.ReportSeller); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ReportSeller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
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
