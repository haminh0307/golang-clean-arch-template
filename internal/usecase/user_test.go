package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/interface/usecase"
	ucimpl "github.com/haminh0307/golang-clean-arch-template/internal/usecase"
	mocksrepository "github.com/haminh0307/golang-clean-arch-template/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite

	mockRepo *mocksrepository.User
	usecase  usecase.User
}

func (suite *UserTestSuite) SetupSuite() {
	suite.mockRepo = mocksrepository.NewUser(suite.T())
	suite.usecase = ucimpl.NewUser(suite.mockRepo)
}

func (suite *UserTestSuite) TestReadOne() {
	user := entity.User{
		Base: entity.Base{
			ID: entity.ID("0"),
		},
	}

	// testcases
	testcases := []struct {
		Name          string
		MockInput     []any
		MockOutput    []any
		UserID        entity.ID
		ExpectedUser  *entity.User
		ExpectedError error
	}{
		{
			Name: "OK",
			MockInput: []any{
				mock.Anything,
				user.ID,
			},
			MockOutput: []any{
				&user,
				nil,
			},
			UserID:        user.ID,
			ExpectedUser:  &user,
			ExpectedError: nil,
		},
		{
			Name: "Error",
			MockInput: []any{
				mock.Anything,
				entity.ID("1"),
			},
			MockOutput: []any{
				nil,
				errors.New("unexpected"),
			},
			UserID:        entity.ID("1"),
			ExpectedUser:  nil,
			ExpectedError: errors.New("unexpected"),
		},
	}

	for _, testcase := range testcases {
		suite.Run(testcase.Name, func() {
			// setup expected calls
			suite.mockRepo.On("ReadOne", testcase.MockInput...).Return(testcase.MockOutput...).Once()

			user, err := suite.usecase.ReadOne(context.Background(), testcase.UserID)

			// check return values
			suite.Equal(testcase.ExpectedUser, user)
			suite.Equal(testcase.ExpectedError, err)
		})
	}
}

func (suite *UserTestSuite) TestUpdateOne() {
	// testcases
	testcases := []struct {
		Name          string
		MockInput     []any
		MockOutput    []any
		UserID        entity.ID
		Update        *entity.UserToUpdate
		ExpectedError error
	}{
		{
			Name: "OK",
			MockInput: []any{
				mock.Anything,
				entity.ID("0"),
				&entity.UserToUpdate{},
			},
			MockOutput: []any{
				nil,
			},
			UserID:        entity.ID("0"),
			Update:        &entity.UserToUpdate{},
			ExpectedError: nil,
		},
		{
			Name: "Error",
			MockInput: []any{
				mock.Anything,
				entity.ID("1"),
				&entity.UserToUpdate{},
			},
			MockOutput: []any{
				errors.New("unexpected"),
			},
			UserID:        entity.ID("1"),
			Update:        &entity.UserToUpdate{},
			ExpectedError: errors.New("unexpected"),
		},
	}

	for _, testcase := range testcases {
		suite.Run(testcase.Name, func() {
			// setup expected calls
			suite.mockRepo.On("UpdateOne", testcase.MockInput...).Return(testcase.MockOutput...).Once()

			err := suite.usecase.UpdateOne(context.Background(), testcase.UserID, testcase.Update)

			// check return values
			suite.Equal(testcase.ExpectedError, err)
		})
	}
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
