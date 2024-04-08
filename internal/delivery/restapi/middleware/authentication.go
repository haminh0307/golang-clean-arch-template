package middleware

import (
	"errors"
	"net/http"
	"rainbow-love-memory/internal/delivery/restapi"
	"rainbow-love-memory/internal/domain/entity"
	"rainbow-love-memory/internal/domain/interface/infra"
	"strings"

	"github.com/gin-gonic/gin"
)

var ErrInvalidAuthentication = errors.New("invalid authentication")

type Authentication struct {
	jwtProvider infra.JwtProvider
}

func NewAuthentication(jp infra.JwtProvider) *Authentication {
	return &Authentication{jp}
}

func (a *Authentication) Authenticate(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	scheme, tokenString, found := strings.Cut(authHeader, " ")
	if !found || scheme != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, restapi.Response{Error: ErrInvalidAuthentication.Error()})
		return
	}

	var claims entity.Claims
	_, err := a.jwtProvider.ParseWithClaims(tokenString, &claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, restapi.Response{Error: ErrInvalidAuthentication.Error()})
		return
	}

	if ctx.Param("userID") != claims.Subject {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.Next()
}
