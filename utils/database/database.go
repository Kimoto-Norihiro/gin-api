package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

var Db *gorm.DB

func Init() error {
	dsn := os.Getenv("DATABASE_URL")
	log.Println(dsn)

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			return err
	}
	log.Println(Db)

	// ログ出力有効化
	Db.Debug()

	// テーブル自動マイグレーション
	err = Db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
			return err
	}
	log.Println(Db)

	return nil
}
