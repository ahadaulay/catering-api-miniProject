package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"os"
)

func HashPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password+os.Getenv("SALT")), bcrypt.DefaultCost)
	return string(result)
}

func ComparePassword(hash, password string) error {
	passwordWithSalt := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordWithSalt))
	if err != nil {
		return err
	}

	return nil
}
