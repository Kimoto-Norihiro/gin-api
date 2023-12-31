package database

import (
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

func ConnDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Debug()

	return db, nil
}

func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		return err
	}
	return nil
}
