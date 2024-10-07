// Code generated by mockery v2.46.2. DO NOT EDIT.

package core

import (
	context "context"

	chat "github.com/kotopesp/sos-kotopes/internal/controller/http/model/chat"

	core "github.com/kotopesp/sos-kotopes/internal/core"

	mock "github.com/stretchr/testify/mock"
)

// MockMessageStore is an autogenerated mock type for the MessageStore type
type MockMessageStore struct {
	mock.Mock
}

type MockMessageStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMessageStore) EXPECT() *MockMessageStore_Expecter {
	return &MockMessageStore_Expecter{mock: &_m.Mock}
}

// CreateMessage provides a mock function with given fields: ctx, data
func (_m *MockMessageStore) CreateMessage(ctx context.Context, data chat.Message) (chat.Message, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for CreateMessage")
	}

	var r0 chat.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chat.Message) (chat.Message, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chat.Message) chat.Message); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(chat.Message)
	}

	if rf, ok := ret.Get(1).(func(context.Context, chat.Message) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMessageStore_CreateMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMessage'
type MockMessageStore_CreateMessage_Call struct {
	*mock.Call
}

// CreateMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - data chat.Message
func (_e *MockMessageStore_Expecter) CreateMessage(ctx interface{}, data interface{}) *MockMessageStore_CreateMessage_Call {
	return &MockMessageStore_CreateMessage_Call{Call: _e.mock.On("CreateMessage", ctx, data)}
}

func (_c *MockMessageStore_CreateMessage_Call) Run(run func(ctx context.Context, data chat.Message)) *MockMessageStore_CreateMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(chat.Message))
	})
	return _c
}

func (_c *MockMessageStore_CreateMessage_Call) Return(message chat.Message, err error) *MockMessageStore_CreateMessage_Call {
	_c.Call.Return(message, err)
	return _c
}

func (_c *MockMessageStore_CreateMessage_Call) RunAndReturn(run func(context.Context, chat.Message) (chat.Message, error)) *MockMessageStore_CreateMessage_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteMessage provides a mock function with given fields: ctx, chatID, userID, messageID
func (_m *MockMessageStore) DeleteMessage(ctx context.Context, chatID int, userID int, messageID int) error {
	ret := _m.Called(ctx, chatID, userID, messageID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) error); ok {
		r0 = rf(ctx, chatID, userID, messageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageStore_DeleteMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteMessage'
type MockMessageStore_DeleteMessage_Call struct {
	*mock.Call
}

// DeleteMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - chatID int
//   - userID int
//   - messageID int
func (_e *MockMessageStore_Expecter) DeleteMessage(ctx interface{}, chatID interface{}, userID interface{}, messageID interface{}) *MockMessageStore_DeleteMessage_Call {
	return &MockMessageStore_DeleteMessage_Call{Call: _e.mock.On("DeleteMessage", ctx, chatID, userID, messageID)}
}

func (_c *MockMessageStore_DeleteMessage_Call) Run(run func(ctx context.Context, chatID int, userID int, messageID int)) *MockMessageStore_DeleteMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *MockMessageStore_DeleteMessage_Call) Return(err error) *MockMessageStore_DeleteMessage_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockMessageStore_DeleteMessage_Call) RunAndReturn(run func(context.Context, int, int, int) error) *MockMessageStore_DeleteMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllMessages provides a mock function with given fields: ctx, chatID, userID, sortType, searchText
func (_m *MockMessageStore) GetAllMessages(ctx context.Context, chatID int, userID int, sortType string, searchText string) ([]chat.Message, error) {
	ret := _m.Called(ctx, chatID, userID, sortType, searchText)

	if len(ret) == 0 {
		panic("no return value specified for GetAllMessages")
	}

	var r0 []chat.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string, string) ([]chat.Message, error)); ok {
		return rf(ctx, chatID, userID, sortType, searchText)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string, string) []chat.Message); ok {
		r0 = rf(ctx, chatID, userID, sortType, searchText)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chat.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, string, string) error); ok {
		r1 = rf(ctx, chatID, userID, sortType, searchText)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMessageStore_GetAllMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllMessages'
type MockMessageStore_GetAllMessages_Call struct {
	*mock.Call
}

// GetAllMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - chatID int
//   - userID int
//   - sortType string
//   - searchText string
func (_e *MockMessageStore_Expecter) GetAllMessages(ctx interface{}, chatID interface{}, userID interface{}, sortType interface{}, searchText interface{}) *MockMessageStore_GetAllMessages_Call {
	return &MockMessageStore_GetAllMessages_Call{Call: _e.mock.On("GetAllMessages", ctx, chatID, userID, sortType, searchText)}
}

func (_c *MockMessageStore_GetAllMessages_Call) Run(run func(ctx context.Context, chatID int, userID int, sortType string, searchText string)) *MockMessageStore_GetAllMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockMessageStore_GetAllMessages_Call) Return(messages []chat.Message, err error) *MockMessageStore_GetAllMessages_Call {
	_c.Call.Return(messages, err)
	return _c
}

func (_c *MockMessageStore_GetAllMessages_Call) RunAndReturn(run func(context.Context, int, int, string, string) ([]chat.Message, error)) *MockMessageStore_GetAllMessages_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnreadMessageCount provides a mock function with given fields: ctx, chatID, userID
func (_m *MockMessageStore) GetUnreadMessageCount(ctx context.Context, chatID int, userID int) (int64, error) {
	ret := _m.Called(ctx, chatID, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUnreadMessageCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (int64, error)); ok {
		return rf(ctx, chatID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) int64); ok {
		r0 = rf(ctx, chatID, userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, chatID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMessageStore_GetUnreadMessageCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnreadMessageCount'
type MockMessageStore_GetUnreadMessageCount_Call struct {
	*mock.Call
}

// GetUnreadMessageCount is a helper method to define mock.On call
//   - ctx context.Context
//   - chatID int
//   - userID int
func (_e *MockMessageStore_Expecter) GetUnreadMessageCount(ctx interface{}, chatID interface{}, userID interface{}) *MockMessageStore_GetUnreadMessageCount_Call {
	return &MockMessageStore_GetUnreadMessageCount_Call{Call: _e.mock.On("GetUnreadMessageCount", ctx, chatID, userID)}
}

func (_c *MockMessageStore_GetUnreadMessageCount_Call) Run(run func(ctx context.Context, chatID int, userID int)) *MockMessageStore_GetUnreadMessageCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockMessageStore_GetUnreadMessageCount_Call) Return(_a0 int64, _a1 error) *MockMessageStore_GetUnreadMessageCount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMessageStore_GetUnreadMessageCount_Call) RunAndReturn(run func(context.Context, int, int) (int64, error)) *MockMessageStore_GetUnreadMessageCount_Call {
	_c.Call.Return(run)
	return _c
}

// MarkMessagesAsRead provides a mock function with given fields: ctx, chatID, userID
func (_m *MockMessageStore) MarkMessagesAsRead(ctx context.Context, chatID int, userID int) error {
	ret := _m.Called(ctx, chatID, userID)

	if len(ret) == 0 {
		panic("no return value specified for MarkMessagesAsRead")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, chatID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageStore_MarkMessagesAsRead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkMessagesAsRead'
type MockMessageStore_MarkMessagesAsRead_Call struct {
	*mock.Call
}

// MarkMessagesAsRead is a helper method to define mock.On call
//   - ctx context.Context
//   - chatID int
//   - userID int
func (_e *MockMessageStore_Expecter) MarkMessagesAsRead(ctx interface{}, chatID interface{}, userID interface{}) *MockMessageStore_MarkMessagesAsRead_Call {
	return &MockMessageStore_MarkMessagesAsRead_Call{Call: _e.mock.On("MarkMessagesAsRead", ctx, chatID, userID)}
}

func (_c *MockMessageStore_MarkMessagesAsRead_Call) Run(run func(ctx context.Context, chatID int, userID int)) *MockMessageStore_MarkMessagesAsRead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockMessageStore_MarkMessagesAsRead_Call) Return(_a0 error) *MockMessageStore_MarkMessagesAsRead_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageStore_MarkMessagesAsRead_Call) RunAndReturn(run func(context.Context, int, int) error) *MockMessageStore_MarkMessagesAsRead_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMessage provides a mock function with given fields: ctx, chatID, userID, messageID, data
func (_m *MockMessageStore) UpdateMessage(ctx context.Context, chatID int, userID int, messageID int, data string) (core.Message, error) {
	ret := _m.Called(ctx, chatID, userID, messageID, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMessage")
	}

	var r0 core.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int, string) (core.Message, error)); ok {
		return rf(ctx, chatID, userID, messageID, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int, string) core.Message); ok {
		r0 = rf(ctx, chatID, userID, messageID, data)
	} else {
		r0 = ret.Get(0).(core.Message)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, int, string) error); ok {
		r1 = rf(ctx, chatID, userID, messageID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMessageStore_UpdateMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMessage'
type MockMessageStore_UpdateMessage_Call struct {
	*mock.Call
}

// UpdateMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - chatID int
//   - userID int
//   - messageID int
//   - data string
func (_e *MockMessageStore_Expecter) UpdateMessage(ctx interface{}, chatID interface{}, userID interface{}, messageID interface{}, data interface{}) *MockMessageStore_UpdateMessage_Call {
	return &MockMessageStore_UpdateMessage_Call{Call: _e.mock.On("UpdateMessage", ctx, chatID, userID, messageID, data)}
}

func (_c *MockMessageStore_UpdateMessage_Call) Run(run func(ctx context.Context, chatID int, userID int, messageID int, data string)) *MockMessageStore_UpdateMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int), args[3].(int), args[4].(string))
	})
	return _c
}

func (_c *MockMessageStore_UpdateMessage_Call) Return(message core.Message, err error) *MockMessageStore_UpdateMessage_Call {
	_c.Call.Return(message, err)
	return _c
}

func (_c *MockMessageStore_UpdateMessage_Call) RunAndReturn(run func(context.Context, int, int, int, string) (core.Message, error)) *MockMessageStore_UpdateMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMessageStore creates a new instance of MockMessageStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessageStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessageStore {
	mock := &MockMessageStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
