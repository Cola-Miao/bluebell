package utils

import (
	"bluebell/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

const (
	JWTExpiresTime = time.Hour * 24 * 3
	jwtSecret      = "no secret"
)

func GenerateJWT(u *model.User) (string, error) {
	now := time.Now()
	claims := model.BBClaims{
		UUID: u.UUID,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "bluebell",
			Subject:   "auth_token",
			Audience:  jwt.ClaimStrings{u.Username},
			ExpiresAt: jwt.NewNumericDate(now.Add(JWTExpiresTime)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			// TODO: Need ID Generator?
			ID: "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	ss, err := token.SignedString([]byte(jwtSecret))
	return ss, err
}

func ParseJWT(tokenString string) (user *model.User, newJWT string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.BBClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, "", fmt.Errorf("parse token failed: %w", err)
	}
	claim, ok := token.Claims.(*model.BBClaims)
	if !ok {
		return nil, "", fmt.Errorf("wrong claim type: %w", err)
	}
	newJWT, err = refreshJWT(claim)
	if err != nil {
		slog.Warn("refresh JWT failed", "error", err.Error())
	}
	return &model.User{UUID: claim.UUID, Username: claim.Audience[0]}, newJWT, nil
}

func refreshJWT(claims *model.BBClaims) (string, error) {
	if time.Now().Add(JWTExpiresTime / 2).Before(claims.ExpiresAt.Time) {
		return "", nil
	}
	fmt.Println("refresh")
	return GenerateJWT(&model.User{UUID: claims.UUID, Username: claims.Audience[0]})
}

func SetJWT(c *gin.Context, token string) {
	c.SetCookie("jwt", token, int(JWTExpiresTime.Seconds()), "/", "/", false, false)
}
