// Code generated by mockery v2.46.3. DO NOT EDIT.

package core

import (
	context "context"

	core "github.com/kotopesp/sos-kotopes/internal/core"
	mock "github.com/stretchr/testify/mock"
)

// MockUserFavouriteService is an autogenerated mock type for the UserFavouriteService type
type MockUserFavouriteService struct {
	mock.Mock
}

type MockUserFavouriteService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserFavouriteService) EXPECT() *MockUserFavouriteService_Expecter {
	return &MockUserFavouriteService_Expecter{mock: &_m.Mock}
}

// AddUserToFavourite provides a mock function with given fields: ctx, personID, userID
func (_m *MockUserFavouriteService) AddUserToFavourite(ctx context.Context, personID int, userID int) error {
	ret := _m.Called(ctx, personID, userID)

	if len(ret) == 0 {
		panic("no return value specified for AddUserToFavourite")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, personID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserFavouriteService_AddUserToFavourite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUserToFavourite'
type MockUserFavouriteService_AddUserToFavourite_Call struct {
	*mock.Call
}

// AddUserToFavourite is a helper method to define mock.On call
//   - ctx context.Context
//   - personID int
//   - userID int
func (_e *MockUserFavouriteService_Expecter) AddUserToFavourite(ctx interface{}, personID interface{}, userID interface{}) *MockUserFavouriteService_AddUserToFavourite_Call {
	return &MockUserFavouriteService_AddUserToFavourite_Call{Call: _e.mock.On("AddUserToFavourite", ctx, personID, userID)}
}

func (_c *MockUserFavouriteService_AddUserToFavourite_Call) Run(run func(ctx context.Context, personID int, userID int)) *MockUserFavouriteService_AddUserToFavourite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockUserFavouriteService_AddUserToFavourite_Call) Return(err error) *MockUserFavouriteService_AddUserToFavourite_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUserFavouriteService_AddUserToFavourite_Call) RunAndReturn(run func(context.Context, int, int) error) *MockUserFavouriteService_AddUserToFavourite_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUserFromFavourite provides a mock function with given fields: ctx, personID, userID
func (_m *MockUserFavouriteService) DeleteUserFromFavourite(ctx context.Context, personID int, userID int) error {
	ret := _m.Called(ctx, personID, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserFromFavourite")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, personID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserFavouriteService_DeleteUserFromFavourite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserFromFavourite'
type MockUserFavouriteService_DeleteUserFromFavourite_Call struct {
	*mock.Call
}

// DeleteUserFromFavourite is a helper method to define mock.On call
//   - ctx context.Context
//   - personID int
//   - userID int
func (_e *MockUserFavouriteService_Expecter) DeleteUserFromFavourite(ctx interface{}, personID interface{}, userID interface{}) *MockUserFavouriteService_DeleteUserFromFavourite_Call {
	return &MockUserFavouriteService_DeleteUserFromFavourite_Call{Call: _e.mock.On("DeleteUserFromFavourite", ctx, personID, userID)}
}

func (_c *MockUserFavouriteService_DeleteUserFromFavourite_Call) Run(run func(ctx context.Context, personID int, userID int)) *MockUserFavouriteService_DeleteUserFromFavourite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockUserFavouriteService_DeleteUserFromFavourite_Call) Return(err error) *MockUserFavouriteService_DeleteUserFromFavourite_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUserFavouriteService_DeleteUserFromFavourite_Call) RunAndReturn(run func(context.Context, int, int) error) *MockUserFavouriteService_DeleteUserFromFavourite_Call {
	_c.Call.Return(run)
	return _c
}

// GetFavouriteUsers provides a mock function with given fields: ctx, userID, params
func (_m *MockUserFavouriteService) GetFavouriteUsers(ctx context.Context, userID int, params core.GetFavourites) ([]core.User, error) {
	ret := _m.Called(ctx, userID, params)

	if len(ret) == 0 {
		panic("no return value specified for GetFavouriteUsers")
	}

	var r0 []core.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, core.GetFavourites) ([]core.User, error)); ok {
		return rf(ctx, userID, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, core.GetFavourites) []core.User); ok {
		r0 = rf(ctx, userID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, core.GetFavourites) error); ok {
		r1 = rf(ctx, userID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserFavouriteService_GetFavouriteUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFavouriteUsers'
type MockUserFavouriteService_GetFavouriteUsers_Call struct {
	*mock.Call
}

// GetFavouriteUsers is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int
//   - params core.GetFavourites
func (_e *MockUserFavouriteService_Expecter) GetFavouriteUsers(ctx interface{}, userID interface{}, params interface{}) *MockUserFavouriteService_GetFavouriteUsers_Call {
	return &MockUserFavouriteService_GetFavouriteUsers_Call{Call: _e.mock.On("GetFavouriteUsers", ctx, userID, params)}
}

func (_c *MockUserFavouriteService_GetFavouriteUsers_Call) Run(run func(ctx context.Context, userID int, params core.GetFavourites)) *MockUserFavouriteService_GetFavouriteUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(core.GetFavourites))
	})
	return _c
}

func (_c *MockUserFavouriteService_GetFavouriteUsers_Call) Return(favouriteUsers []core.User, err error) *MockUserFavouriteService_GetFavouriteUsers_Call {
	_c.Call.Return(favouriteUsers, err)
	return _c
}

func (_c *MockUserFavouriteService_GetFavouriteUsers_Call) RunAndReturn(run func(context.Context, int, core.GetFavourites) ([]core.User, error)) *MockUserFavouriteService_GetFavouriteUsers_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserFavouriteService creates a new instance of MockUserFavouriteService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserFavouriteService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserFavouriteService {
	mock := &MockUserFavouriteService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
