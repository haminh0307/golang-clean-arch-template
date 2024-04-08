package mongorepo

import (
	"context"
	"errors"
	"rainbow-love-memory/internal/domain"
	"rainbow-love-memory/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const users = "users"

type user struct {
	mongo *mongo.Database
}

func NewUser(mongo *mongo.Database) *user {
	return &user{mongo}
}

func (r *user) CreateOne(ctx context.Context, user *entity.UserToCreate) (entity.ID, error) {
	res, err := r.mongo.Collection(users).InsertOne(ctx, user)
	if err != nil {
		return entity.NilID, err
	}

	return entity.ID(res.InsertedID.(primitive.ObjectID).Hex()), nil
}

func (r *user) ReadOne(ctx context.Context, id entity.ID) (*entity.User, error) {
	var res entity.User
	err := r.mongo.Collection(users).FindOne(ctx, bson.M{"_id": id, "isDeleted": false}).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return &res, nil
}

func (r *user) UpdateOne(ctx context.Context, id entity.ID, update *entity.UserToUpdate) error {
	res, err := r.mongo.Collection(users).UpdateOne(ctx, bson.M{"_id": id, "isDeleted": false}, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *user) ReadOneByUsername(ctx context.Context, username string) (*entity.User, error) {
	filter := bson.M{"username": username, "isDeleted": false}

	var res entity.User
	if err := r.mongo.Collection(users).FindOne(ctx, filter).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return &res, nil
}
