package storage

import (
	"gorm.io/gorm"
	"pmokeev/web-chat/internal/models"
)

type AuthStorage struct {
	dbConnection *gorm.DB
}

func NewAuthStorage(dbConnection *gorm.DB) *AuthStorage {
	return &AuthStorage{dbConnection: dbConnection}
}

func (authStorage *AuthStorage) MigrateTable() error {
	err := authStorage.dbConnection.AutoMigrate(&models.RegisterForm{})
	return err
}

func (authStorage *AuthStorage) CreateUser(registerForm models.RegisterForm) error {
	result := authStorage.dbConnection.Create(&registerForm)
	return result.Error
}

func (authStorage *AuthStorage) GetUser(loginForm models.LoginForm) (models.RegisterForm, error) {
	var registerForm models.RegisterForm
	result := authStorage.dbConnection.Table("users").Find(&registerForm, "email = ?", loginForm.Email)
	if result.Error != nil {
		return models.RegisterForm{}, result.Error
	}

	return registerForm, nil
}

func (authStorage *AuthStorage) ChangeUserPassword(changePasswordForm models.ChangePassword) error {
	result := authStorage.dbConnection.Table("users").Where("email = ?", changePasswordForm.Email).Update("password_hash", changePasswordForm.NewPassword)

	return result.Error
}
