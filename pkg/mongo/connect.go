package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/mgocompat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// New returns mongo database instance.
func New(ctx context.Context, connURI string, dbName string) (*mongo.Database, error) {
	client, err := mongo.Connect(
		ctx,
		options.Client().SetRegistry(mgocompat.Registry),
		options.Client().ApplyURI(connURI),
	)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return db, nil
}
