package entity

import "github.com/golang-jwt/jwt/v5"

// Claims is the JWT custom claims.
type Claims struct {
	jwt.RegisteredClaims
}
