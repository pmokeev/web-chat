package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pmokeev/web-chat/internal/models"
)

type AuthStorage struct {
	dbConnection *gorm.DB
}

func NewAuthStorage() *AuthStorage {
	return &AuthStorage{}
}

func (authStorage *AuthStorage) InitDBConnection(config models.DBConfig) error {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.SSLMode,
	)
	dbConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return err
	}

	authStorage.dbConnection = dbConnection
	err = authStorage.dbConnection.AutoMigrate(&models.RegisterForm{})

	return err
}

func (authStorage *AuthStorage) AddNewUser(registerForm models.RegisterForm) error {
	result := authStorage.dbConnection.Create(&registerForm)
	return result.Error
}

func (authStorage *AuthStorage) GetUserPassword(loginForm models.LoginForm) (string, error) {
	var registerForm models.RegisterForm
	result := authStorage.dbConnection.Table("register_forms").Find(&registerForm, "email = ?", loginForm.Email)
	if result.Error != nil {
		return "", result.Error
	}

	return registerForm.PasswordHash, nil
}
