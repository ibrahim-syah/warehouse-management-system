// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "warehouse-management-system/entity"

	mock "github.com/stretchr/testify/mock"
)

// AuthUsecase is an autogenerated mock type for the AuthUsecase type
type AuthUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, req
func (_m *AuthUsecase) Login(ctx context.Context, req *entity.EmailPassword) (*entity.LoginToken, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *entity.LoginToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.EmailPassword) (*entity.LoginToken, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.EmailPassword) *entity.LoginToken); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.LoginToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.EmailPassword) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, req
func (_m *AuthUsecase) Register(ctx context.Context, req *entity.InsertUser) (int, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.InsertUser) (int, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.InsertUser) int); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.InsertUser) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthUsecase creates a new instance of AuthUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthUsecase {
	mock := &AuthUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
