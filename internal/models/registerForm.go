package models

type RegisterForm struct {
	ID           int    `json:"-" gorm:"primaryKey"`
	Name         string `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Email        string `json:"email" binding:"required" gorm:"type:varchar(100);unique;not null"`
	PasswordHash string `json:"password" binding:"required" gorm:"type:varchar(100);not null"`
}
