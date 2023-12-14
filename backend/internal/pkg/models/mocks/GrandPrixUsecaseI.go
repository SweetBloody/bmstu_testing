// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	models "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// GrandPrixUsecaseI is an autogenerated mock type for the GrandPrixUsecaseI type
type GrandPrixUsecaseI struct {
	mock.Mock
}

// Create provides a mock function with given fields: grandPrix
func (_m *GrandPrixUsecaseI) Create(grandPrix *models.GrandPrix) (int, error) {
	ret := _m.Called(grandPrix)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.GrandPrix) (int, error)); ok {
		return rf(grandPrix)
	}
	if rf, ok := ret.Get(0).(func(*models.GrandPrix) int); ok {
		r0 = rf(grandPrix)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*models.GrandPrix) error); ok {
		r1 = rf(grandPrix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *GrandPrixUsecaseI) Delete(id int) error {
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
func (_m *GrandPrixUsecaseI) GetAll() ([]*models.GrandPrix, error) {
	ret := _m.Called()

	var r0 []*models.GrandPrix
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.GrandPrix, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.GrandPrix); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.GrandPrix)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGPById provides a mock function with given fields: id
func (_m *GrandPrixUsecaseI) GetGPById(id int) (*models.GrandPrix, error) {
	ret := _m.Called(id)

	var r0 *models.GrandPrix
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.GrandPrix, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.GrandPrix); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GrandPrix)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, newGrandPrix
func (_m *GrandPrixUsecaseI) Update(id int, newGrandPrix *models.GrandPrix) error {
	ret := _m.Called(id, newGrandPrix)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.GrandPrix) error); ok {
		r0 = rf(id, newGrandPrix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewGrandPrixUsecaseI creates a new instance of GrandPrixUsecaseI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGrandPrixUsecaseI(t interface {
	mock.TestingT
	Cleanup(func())
}) *GrandPrixUsecaseI {
	mock := &GrandPrixUsecaseI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
