// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	models "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// RaceResultRepositoryI is an autogenerated mock type for the RaceResultRepositoryI type
type RaceResultRepositoryI struct {
	mock.Mock
}

// Create provides a mock function with given fields: result
func (_m *RaceResultRepositoryI) Create(result *models.RaceResult) (int, error) {
	ret := _m.Called(result)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.RaceResult) (int, error)); ok {
		return rf(result)
	}
	if rf, ok := ret.Get(0).(func(*models.RaceResult) int); ok {
		r0 = rf(result)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*models.RaceResult) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *RaceResultRepositoryI) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRaceResultById provides a mock function with given fields: id
func (_m *RaceResultRepositoryI) GetRaceResultById(id int) (*models.RaceResultView, error) {
	ret := _m.Called(id)

	var r0 *models.RaceResultView
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.RaceResultView, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.RaceResultView); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RaceResultView)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRaceResultsOfGP provides a mock function with given fields: gp_id
func (_m *RaceResultRepositoryI) GetRaceResultsOfGP(gp_id int) ([]*models.RaceResultView, error) {
	ret := _m.Called(gp_id)

	var r0 []*models.RaceResultView
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*models.RaceResultView, error)); ok {
		return rf(gp_id)
	}
	if rf, ok := ret.Get(0).(func(int) []*models.RaceResultView); ok {
		r0 = rf(gp_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.RaceResultView)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(gp_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: newResult
func (_m *RaceResultRepositoryI) Update(newResult *models.RaceResult) error {
	ret := _m.Called(newResult)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.RaceResult) error); ok {
		r0 = rf(newResult)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRaceResultRepositoryI creates a new instance of RaceResultRepositoryI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRaceResultRepositoryI(t interface {
	mock.TestingT
	Cleanup(func())
}) *RaceResultRepositoryI {
	mock := &RaceResultRepositoryI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
