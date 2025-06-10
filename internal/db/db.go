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

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	if db_password == "" {
		return nil, errors.New("db password not found in .env")
	}

	if db_host == "" {
		return nil, errors.New("db host not found in .env")
	}

	if db_user == "" {
		return nil, errors.New("db username not found in .env")
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC", db_host, db_user, db_password, db_name, db_port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Job{})

	return db, nil
}
