// Code generated by mockery v2.49.0. DO NOT EDIT.

package core

import (
	context "context"

	core "github.com/kotopesp/sos-kotopes/internal/core"
	mock "github.com/stretchr/testify/mock"
)

// MockAnimalStore is an autogenerated mock type for the AnimalStore type
type MockAnimalStore struct {
	mock.Mock
}

type MockAnimalStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAnimalStore) EXPECT() *MockAnimalStore_Expecter {
	return &MockAnimalStore_Expecter{mock: &_m.Mock}
}

// CreateAnimal provides a mock function with given fields: ctx, animal
func (_m *MockAnimalStore) CreateAnimal(ctx context.Context, animal core.Animal) (core.Animal, error) {
	ret := _m.Called(ctx, animal)

	if len(ret) == 0 {
		panic("no return value specified for CreateAnimal")
	}

	var r0 core.Animal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, core.Animal) (core.Animal, error)); ok {
		return rf(ctx, animal)
	}
	if rf, ok := ret.Get(0).(func(context.Context, core.Animal) core.Animal); ok {
		r0 = rf(ctx, animal)
	} else {
		r0 = ret.Get(0).(core.Animal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, core.Animal) error); ok {
		r1 = rf(ctx, animal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAnimalStore_CreateAnimal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAnimal'
type MockAnimalStore_CreateAnimal_Call struct {
	*mock.Call
}

// CreateAnimal is a helper method to define mock.On call
//   - ctx context.Context
//   - animal core.Animal
func (_e *MockAnimalStore_Expecter) CreateAnimal(ctx interface{}, animal interface{}) *MockAnimalStore_CreateAnimal_Call {
	return &MockAnimalStore_CreateAnimal_Call{Call: _e.mock.On("CreateAnimal", ctx, animal)}
}

func (_c *MockAnimalStore_CreateAnimal_Call) Run(run func(ctx context.Context, animal core.Animal)) *MockAnimalStore_CreateAnimal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(core.Animal))
	})
	return _c
}

func (_c *MockAnimalStore_CreateAnimal_Call) Return(_a0 core.Animal, _a1 error) *MockAnimalStore_CreateAnimal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAnimalStore_CreateAnimal_Call) RunAndReturn(run func(context.Context, core.Animal) (core.Animal, error)) *MockAnimalStore_CreateAnimal_Call {
	_c.Call.Return(run)
	return _c
}

// GetAnimalByID provides a mock function with given fields: ctx, id
func (_m *MockAnimalStore) GetAnimalByID(ctx context.Context, id int) (core.Animal, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetAnimalByID")
	}

	var r0 core.Animal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (core.Animal, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) core.Animal); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(core.Animal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAnimalStore_GetAnimalByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAnimalByID'
type MockAnimalStore_GetAnimalByID_Call struct {
	*mock.Call
}

// GetAnimalByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockAnimalStore_Expecter) GetAnimalByID(ctx interface{}, id interface{}) *MockAnimalStore_GetAnimalByID_Call {
	return &MockAnimalStore_GetAnimalByID_Call{Call: _e.mock.On("GetAnimalByID", ctx, id)}
}

func (_c *MockAnimalStore_GetAnimalByID_Call) Run(run func(ctx context.Context, id int)) *MockAnimalStore_GetAnimalByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockAnimalStore_GetAnimalByID_Call) Return(_a0 core.Animal, _a1 error) *MockAnimalStore_GetAnimalByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAnimalStore_GetAnimalByID_Call) RunAndReturn(run func(context.Context, int) (core.Animal, error)) *MockAnimalStore_GetAnimalByID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAnimal provides a mock function with given fields: ctx, animal
func (_m *MockAnimalStore) UpdateAnimal(ctx context.Context, animal core.Animal) (core.Animal, error) {
	ret := _m.Called(ctx, animal)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAnimal")
	}

	var r0 core.Animal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, core.Animal) (core.Animal, error)); ok {
		return rf(ctx, animal)
	}
	if rf, ok := ret.Get(0).(func(context.Context, core.Animal) core.Animal); ok {
		r0 = rf(ctx, animal)
	} else {
		r0 = ret.Get(0).(core.Animal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, core.Animal) error); ok {
		r1 = rf(ctx, animal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAnimalStore_UpdateAnimal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAnimal'
type MockAnimalStore_UpdateAnimal_Call struct {
	*mock.Call
}

// UpdateAnimal is a helper method to define mock.On call
//   - ctx context.Context
//   - animal core.Animal
func (_e *MockAnimalStore_Expecter) UpdateAnimal(ctx interface{}, animal interface{}) *MockAnimalStore_UpdateAnimal_Call {
	return &MockAnimalStore_UpdateAnimal_Call{Call: _e.mock.On("UpdateAnimal", ctx, animal)}
}

func (_c *MockAnimalStore_UpdateAnimal_Call) Run(run func(ctx context.Context, animal core.Animal)) *MockAnimalStore_UpdateAnimal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(core.Animal))
	})
	return _c
}

func (_c *MockAnimalStore_UpdateAnimal_Call) Return(_a0 core.Animal, _a1 error) *MockAnimalStore_UpdateAnimal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAnimalStore_UpdateAnimal_Call) RunAndReturn(run func(context.Context, core.Animal) (core.Animal, error)) *MockAnimalStore_UpdateAnimal_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAnimalStore creates a new instance of MockAnimalStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAnimalStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAnimalStore {
	mock := &MockAnimalStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
