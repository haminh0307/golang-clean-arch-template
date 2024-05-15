// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocksrepository

import (
	context "context"

	entity "github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

type User_Expecter struct {
	mock *mock.Mock
}

func (_m *User) EXPECT() *User_Expecter {
	return &User_Expecter{mock: &_m.Mock}
}

// CreateOne provides a mock function with given fields: ctx, user
func (_m *User) CreateOne(ctx context.Context, user *entity.UserToCreate) (entity.ID, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateOne")
	}

	var r0 entity.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserToCreate) (entity.ID, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserToCreate) entity.ID); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(entity.ID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.UserToCreate) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_CreateOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOne'
type User_CreateOne_Call struct {
	*mock.Call
}

// CreateOne is a helper method to define mock.On call
//   - ctx context.Context
//   - user *entity.UserToCreate
func (_e *User_Expecter) CreateOne(ctx interface{}, user interface{}) *User_CreateOne_Call {
	return &User_CreateOne_Call{Call: _e.mock.On("CreateOne", ctx, user)}
}

func (_c *User_CreateOne_Call) Run(run func(ctx context.Context, user *entity.UserToCreate)) *User_CreateOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.UserToCreate))
	})
	return _c
}

func (_c *User_CreateOne_Call) Return(_a0 entity.ID, _a1 error) *User_CreateOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_CreateOne_Call) RunAndReturn(run func(context.Context, *entity.UserToCreate) (entity.ID, error)) *User_CreateOne_Call {
	_c.Call.Return(run)
	return _c
}

// ReadOne provides a mock function with given fields: ctx, id
func (_m *User) ReadOne(ctx context.Context, id entity.ID) (*entity.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ReadOne")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ID) (*entity.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.ID) *entity.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.ID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_ReadOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadOne'
type User_ReadOne_Call struct {
	*mock.Call
}

// ReadOne is a helper method to define mock.On call
//   - ctx context.Context
//   - id entity.ID
func (_e *User_Expecter) ReadOne(ctx interface{}, id interface{}) *User_ReadOne_Call {
	return &User_ReadOne_Call{Call: _e.mock.On("ReadOne", ctx, id)}
}

func (_c *User_ReadOne_Call) Run(run func(ctx context.Context, id entity.ID)) *User_ReadOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.ID))
	})
	return _c
}

func (_c *User_ReadOne_Call) Return(_a0 *entity.User, _a1 error) *User_ReadOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_ReadOne_Call) RunAndReturn(run func(context.Context, entity.ID) (*entity.User, error)) *User_ReadOne_Call {
	_c.Call.Return(run)
	return _c
}

// ReadOneByUsername provides a mock function with given fields: ctx, username
func (_m *User) ReadOneByUsername(ctx context.Context, username string) (*entity.User, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for ReadOneByUsername")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_ReadOneByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadOneByUsername'
type User_ReadOneByUsername_Call struct {
	*mock.Call
}

// ReadOneByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *User_Expecter) ReadOneByUsername(ctx interface{}, username interface{}) *User_ReadOneByUsername_Call {
	return &User_ReadOneByUsername_Call{Call: _e.mock.On("ReadOneByUsername", ctx, username)}
}

func (_c *User_ReadOneByUsername_Call) Run(run func(ctx context.Context, username string)) *User_ReadOneByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *User_ReadOneByUsername_Call) Return(_a0 *entity.User, _a1 error) *User_ReadOneByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_ReadOneByUsername_Call) RunAndReturn(run func(context.Context, string) (*entity.User, error)) *User_ReadOneByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOne provides a mock function with given fields: ctx, id, update
func (_m *User) UpdateOne(ctx context.Context, id entity.ID, update *entity.UserToUpdate) error {
	ret := _m.Called(ctx, id, update)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ID, *entity.UserToUpdate) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// User_UpdateOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOne'
type User_UpdateOne_Call struct {
	*mock.Call
}

// UpdateOne is a helper method to define mock.On call
//   - ctx context.Context
//   - id entity.ID
//   - update *entity.UserToUpdate
func (_e *User_Expecter) UpdateOne(ctx interface{}, id interface{}, update interface{}) *User_UpdateOne_Call {
	return &User_UpdateOne_Call{Call: _e.mock.On("UpdateOne", ctx, id, update)}
}

func (_c *User_UpdateOne_Call) Run(run func(ctx context.Context, id entity.ID, update *entity.UserToUpdate)) *User_UpdateOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.ID), args[2].(*entity.UserToUpdate))
	})
	return _c
}

func (_c *User_UpdateOne_Call) Return(_a0 error) *User_UpdateOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *User_UpdateOne_Call) RunAndReturn(run func(context.Context, entity.ID, *entity.UserToUpdate) error) *User_UpdateOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}