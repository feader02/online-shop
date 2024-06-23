package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GetJwtKey() []byte {
	key, exists := os.LookupEnv("LOGIN_USER_COOKIE_SECRET_KEY")
	if !exists {
		key = "avion_secret_key_1233211"
	}
	return []byte(key)
}

func CreateJWT(username string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)), // 168 hours = 1 week
		Issuer:    "Avion",
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJwtKey())
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return GetJwtKey(), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
