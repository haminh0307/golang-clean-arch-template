package jwtprovider

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type jwtProvider struct {
	signingMethod jwt.SigningMethod
	signingKey    []byte
}

func NewJwtProvider(alg string, key []byte) *jwtProvider {
	return &jwtProvider{
		jwt.GetSigningMethod(alg),
		key,
	}
}

func (jp *jwtProvider) Issue(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jp.signingMethod, claims)
	return token.SignedString(jp.signingKey)
}

func (jp *jwtProvider) ParseWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jp.signingMethod.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jp.signingKey, nil
	})
}
