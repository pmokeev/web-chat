package storage

import "gorm.io/gorm"

type AuthStorage struct {
	dbConnection *gorm.DB
}

func NewAuthStorage() *AuthStorage {
	return &AuthStorage{}
}
