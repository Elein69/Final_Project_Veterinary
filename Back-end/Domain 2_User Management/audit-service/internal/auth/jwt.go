package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secretKey")

func ValidateToken(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return false, errors.New("invalid token")
	}
	return true, nil
}
