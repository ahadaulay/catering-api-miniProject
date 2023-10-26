package middlewaresadmin

import (
	"catering-api/helpers"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var whitelistAdmin []string = make([]string, 5)

type JwtUserClaims struct {
	ID uint64 `json:"id"`
	jwt.StandardClaims
}

func GenerateTokenAdmin(adminID uint64) (string, error) {
	
	claims := JwtUserClaims{
		adminID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 2).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(helpers.GetConfig("ADMIN_TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	whitelistAdmin = append(whitelistAdmin, token)

	return token, nil
}

func GetUserCustomer(c echo.Context) *JwtUserClaims {
	user := c.Get("user").(*jwt.Token)

	isListed := CheckTokenAdmin(user.Raw)

	if !isListed {
		return nil
	}

	// jika ada code yang panic
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	claims := user.Claims.(*JwtUserClaims)
	return claims
}

func CheckTokenAdmin(token string) bool {
	for _, tkn := range whitelistAdmin {
		if tkn == token {
			return true
		}
	}

	return false
}

func LogoutAdmin(token string) bool {
	for idx, tkn := range whitelistAdmin {
		if tkn == token {
			whitelistAdmin = append(whitelistAdmin[:idx], whitelistAdmin[idx+1:]...)
		}
	}

	return true
}
