// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	models "mercado-frescos-time-7/go-web/internal/models"
	repository "mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryMysql is an autogenerated mock type for the RepositoryMysql type
type RepositoryMysql struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *RepositoryMysql) Create(data models.PurchaseOrders) (repository.ResultPost, error) {
	ret := _m.Called(data)

	var r0 repository.ResultPost
	if rf, ok := ret.Get(0).(func(models.PurchaseOrders) repository.ResultPost); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(repository.ResultPost)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.PurchaseOrders) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositoryMysql interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryMysql creates a new instance of RepositoryMysql. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryMysql(t mockConstructorTestingTNewRepositoryMysql) *RepositoryMysql {
	mock := &RepositoryMysql{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}