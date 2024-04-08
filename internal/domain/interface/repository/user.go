package repository

import (
	"context"

	"rainbow-love-memory/internal/domain/entity"
)

type User interface {
	CreateOne(ctx context.Context, user *entity.UserToCreate) (entity.ID, error)
	ReadOne(ctx context.Context, id entity.ID) (*entity.User, error)
	UpdateOne(ctx context.Context, id entity.ID, update *entity.UserToUpdate) error
	ReadOneByUsername(ctx context.Context, username string) (*entity.User, error)
}
