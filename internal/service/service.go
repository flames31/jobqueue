package service

import "gorm.io/gorm"

type Service struct {
	JobService  *JobService
	UserService *UserService
}

func NewService(dbConn *gorm.DB) *Service {
	return &Service{
		JobService:  &JobService{jobDB: dbConn},
		UserService: &UserService{userDB: dbConn},
	}
}
