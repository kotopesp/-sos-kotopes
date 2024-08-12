// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/kotopesp/sos-kotopes/internal/core"
	mock "github.com/stretchr/testify/mock"
)

// KeeperService is an autogenerated mock type for the KeeperService type
type KeeperService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, keeper
func (_m *KeeperService) Create(ctx context.Context, keeper core.Keepers) error {
	ret := _m.Called(ctx, keeper)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.Keepers) error); ok {
		r0 = rf(ctx, keeper)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateReview provides a mock function with given fields: ctx, keeperReview
func (_m *KeeperService) CreateReview(ctx context.Context, keeperReview core.KeeperReviews) error {
	ret := _m.Called(ctx, keeperReview)

	if len(ret) == 0 {
		panic("no return value specified for CreateReview")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.KeeperReviews) error); ok {
		r0 = rf(ctx, keeperReview)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *KeeperService) DeleteByID(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReviewByID provides a mock function with given fields: ctx, id
func (_m *KeeperService) DeleteReviewByID(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteReviewByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, params
func (_m *KeeperService) GetAll(ctx context.Context, params core.GetAllKeepersParams) ([]core.Keepers, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []core.Keepers
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, core.GetAllKeepersParams) ([]core.Keepers, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, core.GetAllKeepersParams) []core.Keepers); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.Keepers)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, core.GetAllKeepersParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReviews provides a mock function with given fields: ctx, params
func (_m *KeeperService) GetAllReviews(ctx context.Context, params core.GetAllKeeperReviewsParams) ([]core.KeeperReviews, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetAllReviews")
	}

	var r0 []core.KeeperReviews
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, core.GetAllKeeperReviewsParams) ([]core.KeeperReviews, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, core.GetAllKeeperReviewsParams) []core.KeeperReviews); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.KeeperReviews)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, core.GetAllKeeperReviewsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *KeeperService) GetByID(ctx context.Context, id int) (core.Keepers, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 core.Keepers
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (core.Keepers, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) core.Keepers); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(core.Keepers)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SoftDeleteByID provides a mock function with given fields: ctx, id
func (_m *KeeperService) SoftDeleteByID(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for SoftDeleteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SoftDeleteReviewByID provides a mock function with given fields: ctx, id
func (_m *KeeperService) SoftDeleteReviewByID(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for SoftDeleteReviewByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateByID provides a mock function with given fields: ctx, keeper
func (_m *KeeperService) UpdateByID(ctx context.Context, keeper core.Keepers) (core.Keepers, error) {
	ret := _m.Called(ctx, keeper)

	if len(ret) == 0 {
		panic("no return value specified for UpdateByID")
	}

	var r0 core.Keepers
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, core.Keepers) (core.Keepers, error)); ok {
		return rf(ctx, keeper)
	}
	if rf, ok := ret.Get(0).(func(context.Context, core.Keepers) core.Keepers); ok {
		r0 = rf(ctx, keeper)
	} else {
		r0 = ret.Get(0).(core.Keepers)
	}

	if rf, ok := ret.Get(1).(func(context.Context, core.Keepers) error); ok {
		r1 = rf(ctx, keeper)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReviewByID provides a mock function with given fields: ctx, keeperReview
func (_m *KeeperService) UpdateReviewByID(ctx context.Context, keeperReview core.KeeperReviews) error {
	ret := _m.Called(ctx, keeperReview)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReviewByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.KeeperReviews) error); ok {
		r0 = rf(ctx, keeperReview)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKeeperService creates a new instance of KeeperService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKeeperService(t interface {
	mock.TestingT
	Cleanup(func())
}) *KeeperService {
	mock := &KeeperService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
