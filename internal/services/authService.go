package services

import (
	"pmokeev/web-chat/internal/storage"
)

type AuthService struct {
	authStorage *storage.AuthStorage
}

func NewAuthService(authStorage *storage.AuthStorage) *AuthService {
	return &AuthService{authStorage: authStorage}
}
