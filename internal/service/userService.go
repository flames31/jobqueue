package service

import (
	"fmt"

	"github.com/flames31/jobqueue/internal/auth"
	"github.com/flames31/jobqueue/internal/model"
	"gorm.io/gorm"
)

type UserService struct {
	userDB *gorm.DB
}

func (us *UserService) CreateUser(user *model.User) (*model.User, error) {
	err := us.userDB.Create(user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (us *UserService) VerifyCredentials(email, password string) (uint, error) {
	var user model.User
	us.userDB.Where(&model.User{Email: email}).First(&user)

	if err := auth.CheckPasswordHash(user.PasswordHash, password); err != nil {
		return 0, err
	}

	return user.ID, nil
}
