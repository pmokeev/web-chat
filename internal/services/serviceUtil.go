package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashPassword), err
}

func CompareHashPasswords(correctPassword, requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(correctPassword), []byte(requestPassword))
	return err
}
