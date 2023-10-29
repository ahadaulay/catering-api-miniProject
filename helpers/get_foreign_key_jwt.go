package helpers

import (
	"github.com/golang-jwt/jwt"
	"strconv"
)

type JwtClaims struct {
	UserID             string `json:"user_id"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	jwt.StandardClaims        // Embed StandardClaims untuk mendapatkan waktu kedaluwarsa token
}

func GetUserIDJWT(token string) (uint64, error) {

	// Mengekstrak klaim dari token JWT
	claims := &JwtClaims{} // Sesuaikan dengan struktur klaim Anda
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil // Ganti dengan kunci rahasia JWT Anda
	})

	if err != nil {
		// Handle kesalahan JWT
		return 0, err
	}

	// Gunakan user_id dari klaim
	user_idStr := claims.UserID // Sesuaikan dengan nama klaim yang sesuai di token JWT Anda

	user_id, err := strconv.ParseUint(user_idStr, 10, 64)

	return user_id, nil

}
