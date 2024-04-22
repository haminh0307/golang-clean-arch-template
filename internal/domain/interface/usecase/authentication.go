package usecase

import (
	"context"

	"github.com/haminh0307/golang-clean-arch-template/internal/domain/entity"
)

type Authentication interface {
	SignUp(ctx context.Context, user *entity.UserToCreate) (entity.ID, error)
	SignIn(ctx context.Context, username string, password string) (string, error)
}
