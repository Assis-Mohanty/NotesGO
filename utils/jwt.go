package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, isRefresh bool) (string, error) {
	secret := os.Getenv("JWT_ACCESS_SECRET")
	expireTime := time.Minute * 15
	tokenType := "access"

	if isRefresh {
		secret = os.Getenv("JWT_REFRESH_SECRET")
		expireTime = time.Hour * 24 * 7
		tokenType = "refresh"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expireTime).Unix(),
		"type":    tokenType,
	})

	return token.SignedString([]byte(secret))
}

func VerifyJWT(tokenStr string, isRefresh bool) (*jwt.Token, error) {
	secret := os.Getenv("JWT_ACCESS_SECRET")
	if isRefresh {
		secret = os.Getenv("JWT_REFRESH_SECRET")
	}
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(secret), nil
	})
}
