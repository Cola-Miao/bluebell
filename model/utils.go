package model

import "github.com/golang-jwt/jwt/v5"

type BBClaims struct {
	UUID int64
	*jwt.RegisteredClaims
}
