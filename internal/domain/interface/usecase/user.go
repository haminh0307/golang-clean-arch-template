package usecase

import (
	"context"
	"golang-clean-arch-template/internal/domain/entity"
)

type User interface {
	ReadOne(ctx context.Context, id entity.ID) (*entity.User, error)
	UpdateOne(ctx context.Context, id entity.ID, update *entity.UserToUpdate) error
}
