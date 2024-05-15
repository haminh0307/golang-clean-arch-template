package usecase_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/haminh0307/golang-clean-arch-template/internal/domain"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/interface/usecase"
	ucimpl "github.com/haminh0307/golang-clean-arch-template/internal/usecase"
	mocksinfra "github.com/haminh0307/golang-clean-arch-template/mocks/infra"
	mocksrepository "github.com/haminh0307/golang-clean-arch-template/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationTestSuite struct {
	suite.Suite

	mockRepo  *mocksrepository.User
	mockJwt   *mocksinfra.JwtProvider
	jwtExpiry time.Duration
	usecase   usecase.Authentication
}

func (suite *AuthenticationTestSuite) SetupSuite() {
	suite.mockRepo = mocksrepository.NewUser(suite.T())
	suite.mockJwt = mocksinfra.NewJwtProvider(suite.T())
	suite.jwtExpiry = 5 * time.Minute
	suite.usecase = ucimpl.NewAuthentication(suite.mockRepo, suite.mockJwt, suite.jwtExpiry)
}

func (suite *AuthenticationTestSuite) TestSignUp() {
	// testcases
	testcases := []struct {
		Name          string
		MockInput     []any
		MockOutput    []any
		User          *entity.UserToCreate
		ExpectedID    entity.ID
		ExpectedError error
	}{
		{
			Name: "OK",
			MockInput: []any{
				mock.MatchedBy(func(_ context.Context) bool { return true }),
				mock.MatchedBy(func(user *entity.UserToCreate) bool {
					err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(user.Password))
					return err == nil
				}),
			},
			MockOutput: []any{
				entity.ID("0"),
				nil,
			},
			User:          &entity.UserToCreate{},
			ExpectedID:    entity.ID("0"),
			ExpectedError: nil,
		},
		{
			Name: "PasswordTooLong",
			User: &entity.UserToCreate{
				Password: strings.Repeat("*", 73),
			},
			ExpectedID:    entity.NilID,
			ExpectedError: bcrypt.ErrPasswordTooLong,
		},
		{
			Name: "Error",
			MockInput: []any{
				mock.MatchedBy(func(_ context.Context) bool { return true }),
				mock.MatchedBy(func(user *entity.UserToCreate) bool {
					err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(user.Password))
					return err == nil
				}),
			},
			MockOutput: []any{
				entity.NilID,
				errors.New("unexpected"),
			},
			User:          &entity.UserToCreate{},
			ExpectedID:    entity.NilID,
			ExpectedError: errors.New("unexpected"),
		},
	}

	for _, testcase := range testcases {
		suite.Run(testcase.Name, func() {
			// setup expected calls
			if len(testcase.MockInput) > 0 {
				suite.mockRepo.On("CreateOne", testcase.MockInput...).Return(testcase.MockOutput...).Once()
			}

			id, err := suite.usecase.SignUp(context.Background(), testcase.User)

			// check return values
			suite.Equal(testcase.ExpectedID, id)
			suite.Equal(testcase.ExpectedError, err)
		})
	}
}

func (suite *AuthenticationTestSuite) TestSignIn() {
	// testcases
	testcases := []struct {
		Name           string
		MockRepoInput  []any
		MockRepoOutput []any
		MockJwtInput   []any
		MockJwtOutput  []any
		Username       string
		Password       string
		ExpectedToken  string
		ExpectedError  error
	}{
		{
			Name: "OK",
			MockRepoInput: []any{
				mock.MatchedBy(func(_ context.Context) bool { return true }),
				"username",
			},
			MockRepoOutput: []any{
				&entity.User{
					Base:     entity.Base{ID: "0"},
					Username: "username",
					// password hash generated from "password"
					PasswordHash: []byte("$2a$10$pLGvEXpQiEfrFzyFjrOpzOvRWjfmQM2Mhvg8MxddFsVSHr1PalNtu"),
				},
				nil,
			},
			MockJwtInput: []any{
				mock.MatchedBy(func(claims entity.Claims) bool {
					return claims.Subject == "0" &&
						claims.ExpiresAt.Sub(claims.IssuedAt.Time) == suite.jwtExpiry &&
						claims.NotBefore.Time == claims.IssuedAt.Time
				}),
			},
			MockJwtOutput: []any{
				"token",
				nil,
			},
			Username:      "username",
			Password:      "password",
			ExpectedToken: "token",
			ExpectedError: nil,
		},
		{
			Name: "NotFound",
			MockRepoInput: []any{
				mock.MatchedBy(func(_ context.Context) bool { return true }),
				"username1",
			},
			MockRepoOutput: []any{
				nil,
				domain.ErrNotFound,
			},
			Username:      "username1",
			Password:      "password1",
			ExpectedToken: "",
			ExpectedError: domain.ErrNotFound,
		},
		{
			Name: "WrongPassword",
			MockRepoInput: []any{
				mock.MatchedBy(func(_ context.Context) bool { return true }),
				"username2",
			},
			MockRepoOutput: []any{
				&entity.User{
					Base:     entity.Base{ID: "2"},
					Username: "username2",
					// password hash generated from "password"
					PasswordHash: []byte("$2a$10$pLGvEXpQiEfrFzyFjrOpzOvRWjfmQM2Mhvg8MxddFsVSHr1PalNtu"),
				},
				nil,
			},
			Username:      "username2",
			Password:      "password2",
			ExpectedToken: "",
			ExpectedError: domain.ErrWrongCredentials,
		},
		{
			Name: "HashTooShort",
			MockRepoInput: []any{
				mock.MatchedBy(func(_ context.Context) bool { return true }),
				"username3",
			},
			MockRepoOutput: []any{
				&entity.User{
					Base:     entity.Base{ID: "3"},
					Username: "username3",
					// hash too short
					PasswordHash: []byte{},
				},
				nil,
			},
			Username:      "username3",
			Password:      strings.Repeat("password3", 10),
			ExpectedToken: "",
			ExpectedError: bcrypt.ErrHashTooShort,
		},
		{
			Name: "JwtError",
			MockRepoInput: []any{
				mock.MatchedBy(func(_ context.Context) bool { return true }),
				"username4",
			},
			MockRepoOutput: []any{
				&entity.User{
					Base:     entity.Base{ID: "4"},
					Username: "username4",
					// password hash generated from "password"
					PasswordHash: []byte("$2a$10$pLGvEXpQiEfrFzyFjrOpzOvRWjfmQM2Mhvg8MxddFsVSHr1PalNtu"),
				},
				nil,
			},
			MockJwtInput: []any{
				mock.MatchedBy(func(claims entity.Claims) bool {
					return claims.Subject == "4" &&
						claims.ExpiresAt.Sub(claims.IssuedAt.Time) == suite.jwtExpiry &&
						claims.NotBefore.Time == claims.IssuedAt.Time
				}),
			},
			MockJwtOutput: []any{
				"",
				errors.New("unexpected"),
			},
			Username:      "username4",
			Password:      "password",
			ExpectedToken: "",
			ExpectedError: errors.New("unexpected"),
		},
	}

	for _, testcase := range testcases {
		suite.Run(testcase.Name, func() {
			// setup expected calls
			suite.mockRepo.On("ReadOneByUsername", testcase.MockRepoInput...).Return(testcase.MockRepoOutput...).Once()
			if len(testcase.MockJwtInput) > 0 {
				suite.mockJwt.On("Issue", testcase.MockJwtInput...).Return(testcase.MockJwtOutput...).Once()
			}

			token, err := suite.usecase.SignIn(context.Background(), testcase.Username, testcase.Password)

			// check return values
			suite.Equal(testcase.ExpectedToken, token)
			suite.Equal(testcase.ExpectedError, err)
		})
	}
}

func TestAuthentication(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}
