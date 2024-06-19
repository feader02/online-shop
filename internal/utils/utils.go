package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJWT(username string) (string, error) {
	myAppKey, _ := os.LookupEnv("MY_APP_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(myAppKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
