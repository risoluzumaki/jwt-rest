package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/risoluzumaki/jwt/go/pkg/config"
)

// "github.com/golang-jwt/jwt/v5"

var secret = []byte("supersecret")

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, email string) (string, error) {
	config.Load()
	claims := &Claims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, keyFunc)

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, err
	}
	return claims, nil
}

func keyFunc(token *jwt.Token) (any, error) {
	return secret, nil
}
