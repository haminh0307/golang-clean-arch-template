// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocksinfra

import (
	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"
)

// JwtProvider is an autogenerated mock type for the JwtProvider type
type JwtProvider struct {
	mock.Mock
}

type JwtProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *JwtProvider) EXPECT() *JwtProvider_Expecter {
	return &JwtProvider_Expecter{mock: &_m.Mock}
}

// Issue provides a mock function with given fields: claims
func (_m *JwtProvider) Issue(claims jwt.Claims) (string, error) {
	ret := _m.Called(claims)

	if len(ret) == 0 {
		panic("no return value specified for Issue")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(jwt.Claims) (string, error)); ok {
		return rf(claims)
	}
	if rf, ok := ret.Get(0).(func(jwt.Claims) string); ok {
		r0 = rf(claims)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(jwt.Claims) error); ok {
		r1 = rf(claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JwtProvider_Issue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Issue'
type JwtProvider_Issue_Call struct {
	*mock.Call
}

// Issue is a helper method to define mock.On call
//   - claims jwt.Claims
func (_e *JwtProvider_Expecter) Issue(claims interface{}) *JwtProvider_Issue_Call {
	return &JwtProvider_Issue_Call{Call: _e.mock.On("Issue", claims)}
}

func (_c *JwtProvider_Issue_Call) Run(run func(claims jwt.Claims)) *JwtProvider_Issue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(jwt.Claims))
	})
	return _c
}

func (_c *JwtProvider_Issue_Call) Return(_a0 string, _a1 error) *JwtProvider_Issue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *JwtProvider_Issue_Call) RunAndReturn(run func(jwt.Claims) (string, error)) *JwtProvider_Issue_Call {
	_c.Call.Return(run)
	return _c
}

// ParseWithClaims provides a mock function with given fields: tokenString, claims
func (_m *JwtProvider) ParseWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	ret := _m.Called(tokenString, claims)

	if len(ret) == 0 {
		panic("no return value specified for ParseWithClaims")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string, jwt.Claims) (*jwt.Token, error)); ok {
		return rf(tokenString, claims)
	}
	if rf, ok := ret.Get(0).(func(string, jwt.Claims) *jwt.Token); ok {
		r0 = rf(tokenString, claims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string, jwt.Claims) error); ok {
		r1 = rf(tokenString, claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JwtProvider_ParseWithClaims_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseWithClaims'
type JwtProvider_ParseWithClaims_Call struct {
	*mock.Call
}

// ParseWithClaims is a helper method to define mock.On call
//   - tokenString string
//   - claims jwt.Claims
func (_e *JwtProvider_Expecter) ParseWithClaims(tokenString interface{}, claims interface{}) *JwtProvider_ParseWithClaims_Call {
	return &JwtProvider_ParseWithClaims_Call{Call: _e.mock.On("ParseWithClaims", tokenString, claims)}
}

func (_c *JwtProvider_ParseWithClaims_Call) Run(run func(tokenString string, claims jwt.Claims)) *JwtProvider_ParseWithClaims_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(jwt.Claims))
	})
	return _c
}

func (_c *JwtProvider_ParseWithClaims_Call) Return(_a0 *jwt.Token, _a1 error) *JwtProvider_ParseWithClaims_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *JwtProvider_ParseWithClaims_Call) RunAndReturn(run func(string, jwt.Claims) (*jwt.Token, error)) *JwtProvider_ParseWithClaims_Call {
	_c.Call.Return(run)
	return _c
}

// NewJwtProvider creates a new instance of JwtProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJwtProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *JwtProvider {
	mock := &JwtProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
