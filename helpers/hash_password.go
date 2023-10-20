package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string , error){
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	
	if err != nil {
		return string(hashPassword) , err
	}

	return string(hashPassword) , nil
}