package usecase

import (
	"context"

	"github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/interface/repository"
)

type user struct {
	userRepo repository.User
}

func NewUser(ur repository.User) *user {
	return &user{ur}
}

func (uc *user) ReadOne(ctx context.Context, id entity.ID) (*entity.User, error) {
	return uc.userRepo.ReadOne(ctx, id)
}

func (uc *user) UpdateOne(ctx context.Context, id entity.ID, update *entity.UserToUpdate) error {
	return uc.userRepo.UpdateOne(ctx, id, update)
}
