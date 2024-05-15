package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi"
	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi/handler"
	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi/request"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
	mocksusecase "github.com/haminh0307/golang-clean-arch-template/mocks/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthenticationTestSuite struct {
	suite.Suite

	w   *httptest.ResponseRecorder
	ctx *gin.Context

	mockUC  *mocksusecase.Authentication
	handler *handler.Authentication
}

func (suite *AuthenticationTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	suite.mockUC = mocksusecase.NewAuthentication(suite.T())
	suite.handler = handler.NewAuthentication(suite.mockUC)
}

func (suite *AuthenticationTestSuite) SetupSubTest() {
	suite.w = httptest.NewRecorder()

	suite.ctx, _ = gin.CreateTestContext(suite.w)
	suite.ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
}

func (suite *AuthenticationTestSuite) TestSignUp() {
	// testcases
	testcases := []struct {
		Name             string
		MockInput        []any
		MockOutput       []any
		RequestBody      any
		ExpectedCode     int
		ExpectedHeader   http.Header
		ExpectedResponse any
	}{
		{
			Name: "OK",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				&entity.UserToCreate{
					Username: "username0",
					Password: "password0",
					FullName: "fullname0",
				},
			},
			MockOutput: []any{
				entity.ID("0"),
				nil,
			},
			RequestBody: entity.UserToCreate{
				Username: "username0",
				Password: "password0",
				FullName: "fullname0",
			},
			ExpectedCode: http.StatusCreated,
			ExpectedHeader: http.Header{
				"Location": []string{"/users/0"},
			},
			ExpectedResponse: nil,
		},
		{
			Name:         "BadRequest",
			RequestBody:  entity.UserToCreate{},
			ExpectedCode: http.StatusBadRequest,
			ExpectedHeader: http.Header{
				"Content-Type": []string{"application/json; charset=utf-8"},
			},
			ExpectedResponse: restapi.Response{
				Error: `Key: 'UserToCreate.Username' Error:Field validation for 'Username' failed on the 'required' tag
Key: 'UserToCreate.Password' Error:Field validation for 'Password' failed on the 'required' tag
Key: 'UserToCreate.FullName' Error:Field validation for 'FullName' failed on the 'required' tag`,
			},
		},
		{
			Name: "InternalServerError",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				&entity.UserToCreate{
					Username: "username2",
					Password: "password2",
					FullName: "fullname2",
				},
			},
			MockOutput: []any{
				entity.NilID,
				errors.New("unexpected"),
			},
			RequestBody: entity.UserToCreate{
				Username: "username2",
				Password: "password2",
				FullName: "fullname2",
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedHeader: http.Header{
				"Content-Type": []string{"application/json; charset=utf-8"},
			},
			ExpectedResponse: restapi.Response{
				Error: "unexpected",
			},
		},
	}

	for _, testcase := range testcases {
		suite.Run(testcase.Name, func() {
			// setup expected calls
			if len(testcase.MockInput) > 0 {
				suite.mockUC.On("SignUp", testcase.MockInput...).Return(testcase.MockOutput...).Once()
			}

			// setup gin context
			suite.ctx.Request.Header.Set("Content-Type", "application/json")

			jsonbytes, err := json.Marshal(testcase.RequestBody)
			suite.NoError(err)
			suite.ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

			// handle gin context
			suite.handler.SignUp(suite.ctx)

			// workaround for the problem that ctx.Status() does not send the status code until .WriteHeaderNow() is called
			// see https://github.com/gin-gonic/gin/issues/3443
			suite.ctx.Writer.WriteHeaderNow()

			// check http code
			suite.Equal(testcase.ExpectedCode, suite.w.Code)

			// check http header
			suite.Equal(testcase.ExpectedHeader, suite.w.Header())

			// check http response
			if testcase.ExpectedResponse == nil {
				suite.Equal(0, suite.w.Body.Len())
			} else {
				expectedBody, err := json.Marshal(testcase.ExpectedResponse)
				suite.NoError(err)

				body, err := io.ReadAll(suite.w.Body)
				suite.NoError(err)

				suite.Equal(expectedBody, body)
			}
		})
	}
}

func (suite *AuthenticationTestSuite) TestSignIn() {
	// new mock object
	mockUC := mocksusecase.NewAuthentication(suite.T())

	// handler
	h := handler.NewAuthentication(mockUC)

	// testcases
	testcases := []struct {
		Name             string
		MockInput        []any
		MockOutput       []any
		RequestBody      request.SignIn
		ExpectedCode     int
		ExpectedResponse any
	}{
		{
			Name: "OK",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				"username0",
				"password0",
			},
			MockOutput: []any{
				"token",
				nil,
			},
			RequestBody: request.SignIn{
				Username: "username0",
				Password: "password0",
			},
			ExpectedCode:     http.StatusOK,
			ExpectedResponse: restapi.Response{Data: map[string]any{"token": "token"}},
		},
		{
			Name:         "BadRequest",
			RequestBody:  request.SignIn{},
			ExpectedCode: http.StatusBadRequest,
			ExpectedResponse: restapi.Response{
				Error: `Key: 'SignIn.Username' Error:Field validation for 'Username' failed on the 'required' tag
Key: 'SignIn.Password' Error:Field validation for 'Password' failed on the 'required' tag`,
			},
		},
		{
			Name: "NotFound",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				"username2",
				"password2",
			},
			MockOutput: []any{
				"",
				domain.ErrNotFound,
			},
			RequestBody: request.SignIn{
				Username: "username2",
				Password: "password2",
			},
			ExpectedCode:     http.StatusNotFound,
			ExpectedResponse: nil,
		},
		{
			Name: "WrongCredentials",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				"username3",
				"password3",
			},
			MockOutput: []any{
				"",
				domain.ErrWrongCredentials,
			},
			RequestBody: request.SignIn{
				Username: "username3",
				Password: "password3",
			},
			ExpectedCode:     http.StatusUnauthorized,
			ExpectedResponse: restapi.Response{Error: domain.ErrWrongCredentials.Error()},
		},
		{
			Name: "InternalServerError",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				"username4",
				"password4",
			},
			MockOutput: []any{
				"",
				errors.New("unexpected"),
			},
			RequestBody: request.SignIn{
				Username: "username4",
				Password: "password4",
			},
			ExpectedCode:     http.StatusInternalServerError,
			ExpectedResponse: restapi.Response{Error: "unexpected"},
		},
	}

	for _, testcase := range testcases {
		suite.Run(testcase.Name, func() {
			// setup expected calls
			if len(testcase.MockInput) > 0 {
				mockUC.On("SignIn", testcase.MockInput...).Return(testcase.MockOutput...).Once()
			}

			// setup gin context
			suite.ctx.Request.Header.Set("Content-Type", "application/json")

			jsonbytes, err := json.Marshal(testcase.RequestBody)
			suite.NoError(err)
			suite.ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

			// handle gin context
			h.SignIn(suite.ctx)

			// workaround for the problem that ctx.Status() does not send the status code until .WriteHeaderNow() is called
			// see https://github.com/gin-gonic/gin/issues/3443
			suite.ctx.Writer.WriteHeaderNow()

			// check http code
			suite.Equal(testcase.ExpectedCode, suite.w.Code)

			// check http response
			if testcase.ExpectedResponse == nil {
				suite.Equal(0, suite.w.Body.Len())
			} else {
				expectedBody, err := json.Marshal(testcase.ExpectedResponse)
				suite.NoError(err)

				body, err := io.ReadAll(suite.w.Body)
				suite.NoError(err)

				suite.Equal(expectedBody, body)
			}
		})
	}
}

func TestAuthentication(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}
