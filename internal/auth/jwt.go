package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJWT(userID uint, tokenSecret string) (string, error) {
	claimes := jwt.MapClaims{
		"user_id": userID,
		"expiry":  jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claimes)

	return tokenString.SignedString([]byte(tokenSecret))
}
