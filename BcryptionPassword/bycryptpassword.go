package bcryptionpassword

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashpassword), err
}

func CheckPassword(password, hashpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashpassword))
	return err == nil
}
