// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CategoryCodesRepo is an autogenerated mock type for the categoryCodesRepo type
type CategoryCodesRepo struct {
	mock.Mock
}

// GetCategoriesIdByCodes provides a mock function with given fields: _a0, _a1
func (_m *CategoryCodesRepo) GetCategoriesIdByCodes(_a0 context.Context, _a1 []string) ([]int, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCategoriesIdByCodes")
	}

	var r0 []int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]int, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCategoryCodesRepo creates a new instance of CategoryCodesRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategoryCodesRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *CategoryCodesRepo {
	mock := &CategoryCodesRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
