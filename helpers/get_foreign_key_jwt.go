package helpers

import (
	"github.com/golang-jwt/jwt"
)

func GetUserIDJWT(token string) (uint64,error) {
	// Pisahkan token dari "Bearer "
	token = token[len("Bearer "):]

	// Parse token dengan kunci rahasia
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Anda perlu menggantikan "your-secret-key" dengan kunci rahasia yang sesuai dengan token Anda.
		return []byte(GetConfig("TOKEN_SECRET")), nil
	})

	if err != nil {
		return 0 , err
	}

	// Dapatkan user_id dari claims
	user_id, ok := claims["user_id"].(uint64)
	if !ok {
		return 0 ,err
	}

	return user_id , nil
}
