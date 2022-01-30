package storage

import (
	"gorm.io/gorm"
	"pmokeev/web-chat/internal/models"
)

type AuthorizationStorage interface {
	MigrateTable() error
	CreateUser(registerForm models.RegisterForm) error
	GetUser(loginForm models.LoginForm) (models.RegisterForm, error)
	ChangeUserPassword(changePasswordForm models.ChangePassword) error
}

type Storage struct {
	AuthorizationStorage
}

func NewStorage(dbConnection *gorm.DB) *Storage {
	return &Storage{
		AuthorizationStorage: NewAuthStorage(dbConnection)}
}
