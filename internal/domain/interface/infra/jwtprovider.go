package infra

import "github.com/golang-jwt/jwt/v5"

type JwtProvider interface {
	Issue(claims jwt.Claims) (string, error)
	ParseWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error)
}
