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
	"github.com/haminh0307/golang-clean-arch-template/internal/domain"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
	mocksusecase "github.com/haminh0307/golang-clean-arch-template/mocks/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite

	w   *httptest.ResponseRecorder
	ctx *gin.Context

	mockUC  *mocksusecase.User
	handler *handler.User
}

func (suite *UserTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	suite.mockUC = mocksusecase.NewUser(suite.T())
	suite.handler = handler.NewUser(suite.mockUC)
}

func (suite *UserTestSuite) SetupSubTest() {
	suite.w = httptest.NewRecorder()

	suite.ctx, _ = gin.CreateTestContext(suite.w)
	suite.ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
}

func (suite *UserTestSuite) TestReadOne() {
	// testcases
	testcases := []struct {
		Name             string
		MockInput        []any
		MockOutput       []any
		UserID           string
		ExpectedCode     int
		ExpectedResponse any
	}{
		{
			Name: "OK",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				entity.ID("0"),
			},
			MockOutput: []any{
				&entity.User{Base: entity.Base{ID: "0"}},
				nil,
			},
			UserID:       "0",
			ExpectedCode: http.StatusOK,
			ExpectedResponse: restapi.Response{
				Data: map[string]any{"user": &entity.User{Base: entity.Base{ID: "0"}}},
			},
		},
		{
			Name: "NotFound",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				entity.ID("1"),
			},
			MockOutput: []any{
				nil,
				domain.ErrNotFound,
			},
			UserID:           "1",
			ExpectedCode:     http.StatusNotFound,
			ExpectedResponse: nil,
		},
		{
			Name: "InternalServerError",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				entity.ID("2"),
			},
			MockOutput: []any{
				nil,
				errors.New("unexpected"),
			},
			UserID:       "2",
			ExpectedCode: http.StatusInternalServerError,
			ExpectedResponse: restapi.Response{
				Error: "unexpected",
			},
		},
	}

	for _, testcase := range testcases {
		suite.Run(testcase.Name, func() {
			// setup expected calls
			if len(testcase.MockInput) > 0 {
				suite.mockUC.On("ReadOne", testcase.MockInput...).Return(testcase.MockOutput...).Once()
			}

			// setup gin context
			suite.ctx.Request.Header.Set("Content-Type", "application/json")
			suite.ctx.Params = append(suite.ctx.Params, gin.Param{
				Key:   "userID",
				Value: testcase.UserID,
			})

			// handle gin context
			suite.handler.ReadOne(suite.ctx)

			// workaround for the problem that ctx.Status() does not send the status code until .WriteHeaderNow() is called
			// see https://github.com/gin-gonic/gin/issues/3443
			suite.ctx.Writer.WriteHeaderNow()

			// check http code
			suite.EqualValues(testcase.ExpectedCode, suite.w.Code)

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

func (suite *UserTestSuite) TestUpdateOne() {
	// testcases
	testcases := []struct {
		Name             string
		MockInput        []any
		MockOutput       []any
		UserID           string
		RequestBody      entity.UserToUpdate
		ExpectedCode     int
		ExpectedResponse any
	}{
		{
			Name: "OK",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				entity.ID("0"),
				&entity.UserToUpdate{
					Username: "username0",
				},
			},
			MockOutput: []any{
				nil,
			},
			UserID: "0",
			RequestBody: entity.UserToUpdate{
				Username: "username0",
			},
			ExpectedCode:     http.StatusNoContent,
			ExpectedResponse: nil,
		},
		{
			Name:         "BadRequest",
			UserID:       "1",
			RequestBody:  entity.UserToUpdate{},
			ExpectedCode: http.StatusBadRequest,
			ExpectedResponse: restapi.Response{
				Error: `Key: 'UserToUpdate.Username' Error:Field validation for 'Username' failed on the 'required_without' tag
Key: 'UserToUpdate.FullName' Error:Field validation for 'FullName' failed on the 'required_without' tag`,
			},
		},
		{
			Name: "NotFound",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				entity.ID("2"),
				&entity.UserToUpdate{
					Username: "username2",
				},
			},
			MockOutput: []any{
				domain.ErrNotFound,
			},
			UserID: "2",
			RequestBody: entity.UserToUpdate{
				Username: "username2",
			},
			ExpectedCode:     http.StatusNotFound,
			ExpectedResponse: nil,
		},
		{
			Name: "InternalServerError",
			MockInput: []any{
				mock.AnythingOfType("*gin.Context"),
				entity.ID("3"),
				&entity.UserToUpdate{
					Username: "username3",
				},
			},
			MockOutput: []any{
				errors.New("unexpected"),
			},
			UserID: "3",
			RequestBody: entity.UserToUpdate{
				Username: "username3",
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedResponse: restapi.Response{
				Error: "unexpected",
			},
		},
	}

	for _, testcase := range testcases {
		suite.Run(testcase.Name, func() {
			// setup expected calls
			if len(testcase.MockInput) > 0 {
				suite.mockUC.On("UpdateOne", testcase.MockInput...).Return(testcase.MockOutput...).Once()
			}

			// setup gin context
			suite.ctx.Request.Header.Set("Content-Type", "application/json")
			suite.ctx.Params = append(suite.ctx.Params, gin.Param{
				Key:   "userID",
				Value: testcase.UserID,
			})

			jsonbytes, err := json.Marshal(testcase.RequestBody)
			suite.NoError(err)
			suite.ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

			// handle gin context
			suite.handler.UpdateOne(suite.ctx)

			// workaround for the problem that ctx.Status() does not send the status code until .WriteHeaderNow() is called
			// see https://github.com/gin-gonic/gin/issues/3443
			suite.ctx.Writer.WriteHeaderNow()

			// check http code
			suite.EqualValues(testcase.ExpectedCode, suite.w.Code)

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

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
