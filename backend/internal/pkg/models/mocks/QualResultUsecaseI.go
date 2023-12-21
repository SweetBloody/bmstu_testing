// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	models "github.com/SweetBloody/bmstu_testing/backend/internal/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// QualResultUsecaseI is an autogenerated mock type for the QualResultUsecaseI type
type QualResultUsecaseI struct {
	mock.Mock
}

// Create provides a mock function with given fields: result
func (_m *QualResultUsecaseI) Create(result *models.QualResult) (int, error) {
	ret := _m.Called(result)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.QualResult) (int, error)); ok {
		return rf(result)
	}
	if rf, ok := ret.Get(0).(func(*models.QualResult) int); ok {
		r0 = rf(result)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*models.QualResult) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *QualResultUsecaseI) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetQualResultById provides a mock function with given fields: id
func (_m *QualResultUsecaseI) GetQualResultById(id int) (*models.QualResultView, error) {
	ret := _m.Called(id)

	var r0 *models.QualResultView
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.QualResultView, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.QualResultView); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.QualResultView)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQualResultsOfGP provides a mock function with given fields: gp_id
func (_m *QualResultUsecaseI) GetQualResultsOfGP(gp_id int) ([]*models.QualResultView, error) {
	ret := _m.Called(gp_id)

	var r0 []*models.QualResultView
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*models.QualResultView, error)); ok {
		return rf(gp_id)
	}
	if rf, ok := ret.Get(0).(func(int) []*models.QualResultView); ok {
		r0 = rf(gp_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.QualResultView)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(gp_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, newResult
func (_m *QualResultUsecaseI) Update(id int, newResult *models.QualResult) error {
	ret := _m.Called(id, newResult)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.QualResult) error); ok {
		r0 = rf(id, newResult)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQualResultUsecaseI creates a new instance of QualResultUsecaseI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQualResultUsecaseI(t interface {
	mock.TestingT
	Cleanup(func())
}) *QualResultUsecaseI {
	mock := &QualResultUsecaseI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
