package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

var db *gorm.DB

func Init() error {
	dsn := os.Getenv("DATABASE_URL")
	log.Println(dsn)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			return err
	}
	log.Println(db)

	// ログ出力有効化
	db.Debug()

	// テーブル自動マイグレーション
	err = db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
			return err
	}
	log.Println(db)

	return nil
}

func GetDB() *gorm.DB {
  return db
}
