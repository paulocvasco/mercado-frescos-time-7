// Code generated by mockery v2.13.1. DO NOT EDIT.

package mockRepository

import (
	models "github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
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
func (_m *Repository) GetAll() (models.Products, error) {
	ret := _m.Called()

	var r0 models.Products
	if rf, ok := ret.Get(0).(func() models.Products); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Products)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id int) (models.Product, error) {
	ret := _m.Called(id)

	var r0 models.Product
	if rf, ok := ret.Get(0).(func(int) models.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: product
func (_m *Repository) Insert(product models.Product) (models.Product, error) {
	ret := _m.Called(product)

	var r0 models.Product
	if rf, ok := ret.Get(0).(func(models.Product) models.Product); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Product) error); ok {
		r1 = rf(product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: product
func (_m *Repository) Update(product models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(product)
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
