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

func CompareHashPasswords(correctPassword, requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(correctPassword), []byte(requestPassword))
	return err
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

func (authService *AuthService) SignIn(loginForm models.LoginForm) error {
	correctPassword, err := authService.authStorage.GetUserPassword(loginForm)
	if err != nil {
		return err
	}

	return CompareHashPasswords(correctPassword, loginForm.Password)
}
