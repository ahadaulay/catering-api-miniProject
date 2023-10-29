package middleware

import (
	"catering-api/helpers"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateTokenUser(UserID uint64) (string, error) {
	jwtSecret := []byte(helpers.GetConfig("USER_TOKEN_SECRET"))

	claims := jwt.MapClaims{
		"id":   UserID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"iat":  time.Now().Unix(),
		"role": "User",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenAdmin(AdminID uint64) (string, error) {
	jwtSecret := []byte(helpers.GetConfig("ADMIN_TOKEN_SECRET"))

	claims := jwt.MapClaims{
		"id":   AdminID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"iat":  time.Now().Unix(),
		"role": "Admin",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
