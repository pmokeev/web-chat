package services

import (
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/storage"
)

type AuthorizationService interface {
	SignUP(registerForm models.RegisterForm) error
	SignIn(loginForm models.LoginForm) (string, error)
	JWTVerify(JWTTokenString string) (bool, error)
	GetUserInformation(JWTTokenString string) (map[string]string, error, bool)
	ChangeUserPassword(changePasswordForm models.ChangePassword) error
}

type Service struct {
	AuthorizationService
}

func NewService(storage *storage.Storage) *Service {
	return &Service{AuthorizationService: NewAuthService(storage.AuthorizationStorage)}
}
