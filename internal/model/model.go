package model

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Description string
	Status      string
	Result      string
	UserID      uint
}

type User struct {
	gorm.Model
	Email        string `gorm:"unique"`
	PasswordHash string
	Jobs         []Job `gorm:"constraint:OnDelete:CASCADE;"`
}
