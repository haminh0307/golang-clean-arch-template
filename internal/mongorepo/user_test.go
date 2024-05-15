package mongorepo_test

import (
	"context"
	"testing"

	"github.com/haminh0307/golang-clean-arch-template/internal/domain"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
	"github.com/haminh0307/golang-clean-arch-template/internal/mongorepo"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type UserTestSuite struct {
	suite.Suite

	mt *mtest.T
}

func (suite *UserTestSuite) SetupSuite() {
	suite.mt = mtest.New(suite.T(), mtest.NewOptions().ClientType(mtest.Mock))
}

func (suite *UserTestSuite) TestCreateOne() {
	testcases := []struct {
		Name          string
		MockResponses []primitive.D
		ErrorCheck    func(err error) bool
	}{
		{
			Name:          "OK",
			MockResponses: []primitive.D{mtest.CreateSuccessResponse()},
			ErrorCheck: func(err error) bool {
				return err == nil
			},
		},
		{
			Name: "DuplicateKey",
			MockResponses: []primitive.D{mtest.CreateWriteErrorsResponse(mtest.WriteError{
				Index:   1,
				Code:    11000,
				Message: "DuplicateKey",
			})},
			ErrorCheck: func(err error) bool {
				return mongo.IsDuplicateKeyError(err)
			},
		},
		{
			Name:          "Error",
			MockResponses: []primitive.D{bson.D{{Key: "ok", Value: 0}}},
			ErrorCheck: func(err error) bool {
				return err != nil
			},
		},
	}

	for _, testcase := range testcases {
		suite.mt.Run(testcase.Name, func(mt *mtest.T) {
			repo := mongorepo.NewUser(mt.DB)
			mt.AddMockResponses(testcase.MockResponses...)

			id, err := repo.CreateOne(context.Background(), &entity.UserToCreate{})
			if err != nil {
				suite.Equal(entity.NilID, id)
			}
			suite.True(testcase.ErrorCheck(err))
		})
	}
}

func (suite *UserTestSuite) TestReadOne() {
	testcases := []struct {
		Name          string
		MockResponses []primitive.D
		ExpectedUser  *entity.User
		ExpectedError error
	}{
		{
			Name: "OK",
			MockResponses: []primitive.D{
				mtest.CreateCursorResponse(1, ".users", mtest.FirstBatch, bson.D{
					{Key: "_id", Value: primitive.ObjectID([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})},
					{Key: "username", Value: "username"},
					{Key: "fullName", Value: "fullname"},
				}),
			},
			ExpectedUser: &entity.User{
				Base:     entity.Base{ID: entity.ID("000000000000000000000001")},
				Username: "username",
				FullName: "fullname",
			},
			ExpectedError: nil,
		},
		{
			Name:          "NotFound",
			MockResponses: []primitive.D{mtest.CreateCursorResponse(0, ".users", mtest.FirstBatch)},
			ExpectedUser:  nil,
			ExpectedError: domain.ErrNotFound,
		},
		{
			Name:          "Error",
			MockResponses: []primitive.D{bson.D{{Key: "ok", Value: 0}}},
			ExpectedUser:  nil,
			ExpectedError: mongo.CommandError{
				Message: "command failed",
				Raw:     bson.Raw{0xd, 0x0, 0x0, 0x0, 0x10, 0x6f, 0x6b, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
		},
	}

	for _, testcase := range testcases {
		suite.mt.Run(testcase.Name, func(mt *mtest.T) {
			repo := mongorepo.NewUser(mt.DB)
			mt.AddMockResponses(testcase.MockResponses...)

			user, err := repo.ReadOne(context.Background(), entity.ID(""))
			suite.Equal(testcase.ExpectedUser, user)
			suite.Equal(testcase.ExpectedError, err)
		})
	}
}

func (suite *UserTestSuite) TestUpdateOne() {
	testcases := []struct {
		Name          string
		MockResponses []primitive.D
		ExpectedError error
	}{
		{
			Name: "OK",
			MockResponses: []primitive.D{
				bson.D{
					{Key: "ok", Value: 1},
					{Key: "acknowledged", Value: true},
					{Key: "n", Value: 1},
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "NotFound",
			MockResponses: []primitive.D{
				bson.D{
					{Key: "ok", Value: 1},
					{Key: "acknowledged", Value: false},
					{Key: "n", Value: 0},
				},
			},
			ExpectedError: domain.ErrNotFound,
		},
		{
			Name:          "Error",
			MockResponses: []primitive.D{bson.D{{Key: "ok", Value: 0}}},
			ExpectedError: mongo.CommandError{
				Message: "command failed",
				Raw:     bson.Raw{0xd, 0x0, 0x0, 0x0, 0x10, 0x6f, 0x6b, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
		},
	}

	for _, testcase := range testcases {
		suite.mt.Run(testcase.Name, func(mt *mtest.T) {
			repo := mongorepo.NewUser(mt.DB)
			mt.AddMockResponses(testcase.MockResponses...)

			err := repo.UpdateOne(context.Background(), entity.ID(""), &entity.UserToUpdate{})
			suite.Equal(testcase.ExpectedError, err)
		})
	}
}

func (suite *UserTestSuite) TestReadOneByUsername() {
	testcases := []struct {
		Name          string
		MockResponses []primitive.D
		ExpectedUser  *entity.User
		ExpectedError error
	}{
		{
			Name: "OK",
			MockResponses: []primitive.D{
				mtest.CreateCursorResponse(1, ".users", mtest.FirstBatch, bson.D{
					{Key: "_id", Value: primitive.ObjectID([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})},
					{Key: "username", Value: "username"},
					{Key: "fullName", Value: "fullname"},
				}),
			},
			ExpectedUser: &entity.User{
				Base:     entity.Base{ID: entity.ID("000000000000000000000001")},
				Username: "username",
				FullName: "fullname",
			},
			ExpectedError: nil,
		},
		{
			Name:          "NotFound",
			MockResponses: []primitive.D{mtest.CreateCursorResponse(0, ".users", mtest.FirstBatch)},
			ExpectedUser:  nil,
			ExpectedError: domain.ErrNotFound,
		},
		{
			Name:          "Error",
			MockResponses: []primitive.D{bson.D{{Key: "ok", Value: 0}}},
			ExpectedUser:  nil,
			ExpectedError: mongo.CommandError{
				Message: "command failed",
				Raw:     bson.Raw{0xd, 0x0, 0x0, 0x0, 0x10, 0x6f, 0x6b, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
		},
	}

	for _, testcase := range testcases {
		suite.mt.Run(testcase.Name, func(mt *mtest.T) {
			repo := mongorepo.NewUser(mt.DB)
			mt.AddMockResponses(testcase.MockResponses...)

			user, err := repo.ReadOneByUsername(context.Background(), "")
			suite.Equal(testcase.ExpectedUser, user)
			suite.Equal(testcase.ExpectedError, err)
		})
	}
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
