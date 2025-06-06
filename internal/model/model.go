package model

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Description string
	Status      string
	Result      string
}
