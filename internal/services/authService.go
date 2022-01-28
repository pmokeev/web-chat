package services

import (
	"golang.org/x/crypto/bcrypt"
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/storage"
)

type AuthService struct {
	authStorage *storage.AuthStorage
}

func NewAuthService(authStorage *storage.AuthStorage) *AuthService {
	return &AuthService{authStorage: authStorage}
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashPassword), err
}

func (authService *AuthService) SignUP(registerForm models.RegisterForm) error {
	hashPassword, err := HashPassword(registerForm.PasswordHash)
	if err != nil {
		return err
	}

	registerForm.PasswordHash = hashPassword
	err = authService.authStorage.AddNewUser(registerForm)

	return err
}
