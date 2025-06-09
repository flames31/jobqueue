package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/flames31/jobqueue/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")

	if db_password == "" {
		return nil, errors.New("db password not found in .env")
	}

	if db_user == "" {
		return nil, errors.New("db username not found in .env")
	}

	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=jobqueue port=5432 sslmode=disable TimeZone=Asia/Shanghai", db_user, db_password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Job{}, &model.User{})

	return db, nil
}
