// Code generated by mockery v2.13.1. DO NOT EDIT.

package mockRepository

import (
	context "context"
	domain "mercado-frescos-time-7/go-web/internal/sections/domain"

	mock "github.com/stretchr/testify/mock"
)

// SectionRepository is an autogenerated mock type for the SectionRepository type
type SectionRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *SectionRepository) Delete(_a0 context.Context, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: _a0
func (_m *SectionRepository) GetAll(_a0 context.Context) (*domain.Sections, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Sections
	if rf, ok := ret.Get(0).(func(context.Context) *domain.Sections); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Sections)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: _a0, _a1
func (_m *SectionRepository) GetById(_a0 context.Context, _a1 int) (*domain.Section, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Section
	if rf, ok := ret.Get(0).(func(context.Context, int) *domain.Section); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportProducts provides a mock function with given fields: _a0, _a1
func (_m *SectionRepository) GetReportProducts(_a0 context.Context, _a1 int) (*domain.ProductReports, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.ProductReports
	if rf, ok := ret.Get(0).(func(context.Context, int) *domain.ProductReports); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ProductReports)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *SectionRepository) Store(_a0 context.Context, _a1 *domain.Section) (*domain.Section, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Section
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Section) *domain.Section); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Section) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *SectionRepository) Update(_a0 context.Context, _a1 *domain.Section) (*domain.Section, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Section
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Section) *domain.Section); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Section) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSectionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewSectionRepository creates a new instance of SectionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSectionRepository(t mockConstructorTestingTNewSectionRepository) *SectionRepository {
	mock := &SectionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
