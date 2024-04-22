package usecase

import (
	"context"
	"errors"
	"golang-clean-arch-template/internal/domain"
	"golang-clean-arch-template/internal/domain/entity"
	"golang-clean-arch-template/internal/domain/interface/infra"
	"golang-clean-arch-template/internal/domain/interface/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authentication struct {
	userRepo    repository.User
	jwtProvider infra.JwtProvider
	jwtExpiry   time.Duration
}

func NewAuthentication(ur repository.User, jp infra.JwtProvider, e time.Duration) *authentication {
	return &authentication{ur, jp, e}
}

func (uc *authentication) SignUp(ctx context.Context, user *entity.UserToCreate) (entity.ID, error) {
	var err error
	user.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.NilID, err
	}

	id, err := uc.userRepo.CreateOne(ctx, user)
	if err != nil {
		return entity.NilID, err
	}

	return id, nil
}

func (uc *authentication) SignIn(ctx context.Context, username string, password string) (string, error) {
	user, err := uc.userRepo.ReadOneByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", domain.ErrWrongCredentials
		}

		return "", err
	}

	claims := entity.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.jwtExpiry)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return uc.jwtProvider.Issue(claims)
}
