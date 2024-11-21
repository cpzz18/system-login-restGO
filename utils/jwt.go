package utils

import (
	"time"

	"github.com/golang-jwt/jwt"

)

const SecretKey = "secret" // kunci jwt
// var jwtKey = os.Getenv("secret") kunci jwt dari environment variabl


func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
