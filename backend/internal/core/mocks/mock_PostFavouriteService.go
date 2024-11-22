// Code generated by mockery v2.49.0. DO NOT EDIT.

package core

import (
	context "context"

	core "github.com/kotopesp/sos-kotopes/internal/core"
	mock "github.com/stretchr/testify/mock"
)

// MockPostFavouriteService is an autogenerated mock type for the PostFavouriteService type
type MockPostFavouriteService struct {
	mock.Mock
}

type MockPostFavouriteService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPostFavouriteService) EXPECT() *MockPostFavouriteService_Expecter {
	return &MockPostFavouriteService_Expecter{mock: &_m.Mock}
}

// AddToFavourites provides a mock function with given fields: ctx, postFavourite
func (_m *MockPostFavouriteService) AddToFavourites(ctx context.Context, postFavourite core.PostFavourite) error {
	ret := _m.Called(ctx, postFavourite)

	if len(ret) == 0 {
		panic("no return value specified for AddToFavourites")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.PostFavourite) error); ok {
		r0 = rf(ctx, postFavourite)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostFavouriteService_AddToFavourites_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddToFavourites'
type MockPostFavouriteService_AddToFavourites_Call struct {
	*mock.Call
}

// AddToFavourites is a helper method to define mock.On call
//   - ctx context.Context
//   - postFavourite core.PostFavourite
func (_e *MockPostFavouriteService_Expecter) AddToFavourites(ctx interface{}, postFavourite interface{}) *MockPostFavouriteService_AddToFavourites_Call {
	return &MockPostFavouriteService_AddToFavourites_Call{Call: _e.mock.On("AddToFavourites", ctx, postFavourite)}
}

func (_c *MockPostFavouriteService_AddToFavourites_Call) Run(run func(ctx context.Context, postFavourite core.PostFavourite)) *MockPostFavouriteService_AddToFavourites_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(core.PostFavourite))
	})
	return _c
}

func (_c *MockPostFavouriteService_AddToFavourites_Call) Return(_a0 error) *MockPostFavouriteService_AddToFavourites_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostFavouriteService_AddToFavourites_Call) RunAndReturn(run func(context.Context, core.PostFavourite) error) *MockPostFavouriteService_AddToFavourites_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteFromFavourites provides a mock function with given fields: ctx, postID, userID
func (_m *MockPostFavouriteService) DeleteFromFavourites(ctx context.Context, postID int, userID int) error {
	ret := _m.Called(ctx, postID, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFromFavourites")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, postID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostFavouriteService_DeleteFromFavourites_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFromFavourites'
type MockPostFavouriteService_DeleteFromFavourites_Call struct {
	*mock.Call
}

// DeleteFromFavourites is a helper method to define mock.On call
//   - ctx context.Context
//   - postID int
//   - userID int
func (_e *MockPostFavouriteService_Expecter) DeleteFromFavourites(ctx interface{}, postID interface{}, userID interface{}) *MockPostFavouriteService_DeleteFromFavourites_Call {
	return &MockPostFavouriteService_DeleteFromFavourites_Call{Call: _e.mock.On("DeleteFromFavourites", ctx, postID, userID)}
}

func (_c *MockPostFavouriteService_DeleteFromFavourites_Call) Run(run func(ctx context.Context, postID int, userID int)) *MockPostFavouriteService_DeleteFromFavourites_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockPostFavouriteService_DeleteFromFavourites_Call) Return(_a0 error) *MockPostFavouriteService_DeleteFromFavourites_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostFavouriteService_DeleteFromFavourites_Call) RunAndReturn(run func(context.Context, int, int) error) *MockPostFavouriteService_DeleteFromFavourites_Call {
	_c.Call.Return(run)
	return _c
}

// GetFavouritePosts provides a mock function with given fields: ctx, userID
func (_m *MockPostFavouriteService) GetFavouritePosts(ctx context.Context, userID int) ([]core.PostDetails, int, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetFavouritePosts")
	}

	var r0 []core.PostDetails
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]core.PostDetails, int, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []core.PostDetails); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.PostDetails)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) int); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int) error); ok {
		r2 = rf(ctx, userID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockPostFavouriteService_GetFavouritePosts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFavouritePosts'
type MockPostFavouriteService_GetFavouritePosts_Call struct {
	*mock.Call
}

// GetFavouritePosts is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int
func (_e *MockPostFavouriteService_Expecter) GetFavouritePosts(ctx interface{}, userID interface{}) *MockPostFavouriteService_GetFavouritePosts_Call {
	return &MockPostFavouriteService_GetFavouritePosts_Call{Call: _e.mock.On("GetFavouritePosts", ctx, userID)}
}

func (_c *MockPostFavouriteService_GetFavouritePosts_Call) Run(run func(ctx context.Context, userID int)) *MockPostFavouriteService_GetFavouritePosts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockPostFavouriteService_GetFavouritePosts_Call) Return(_a0 []core.PostDetails, _a1 int, _a2 error) *MockPostFavouriteService_GetFavouritePosts_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockPostFavouriteService_GetFavouritePosts_Call) RunAndReturn(run func(context.Context, int) ([]core.PostDetails, int, error)) *MockPostFavouriteService_GetFavouritePosts_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPostFavouriteService creates a new instance of MockPostFavouriteService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPostFavouriteService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPostFavouriteService {
	mock := &MockPostFavouriteService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
