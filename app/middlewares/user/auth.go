package middlewares

import (
	"catering-api/helpers"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var whitelistUser []string = make([]string, 5)

type JwtCostumerClaims struct {
	ID uint64 `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func GenerateTokenUser(userID uint64) (string, error) {
	role := "user"
	claims := JwtCostumerClaims{
		userID,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 2).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(helpers.GetConfig("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	whitelistUser = append(whitelistUser, token)

	return token, nil
}

func GetUserCustomer(c echo.Context) *JwtCostumerClaims {
	user := c.Get("user").(*jwt.Token)

	isListed := CheckTokenUser(user.Raw)

	if !isListed {
		return nil
	}

	// jika ada code yang panic
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	claims := user.Claims.(*JwtCostumerClaims)
	return claims
}

func CheckTokenUser(token string) bool {
	for _, tkn := range whitelistUser {
		if tkn == token {
			return true
		}
	}

	return false
}

func LogoutUser(token string) bool {
	for idx, tkn := range whitelistUser {
		if tkn == token {
			whitelistUser = append(whitelistUser[:idx], whitelistUser[idx+1:]...)
		}
	}

	return true
}
