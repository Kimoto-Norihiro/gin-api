package utils

import (
	"os"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

var Db *gorm.DB

func init() {
	dsn := os.Getenv("CLEARDB_DATABASE_URL")
	var err error
	
  Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
		log.Println("DB接続エラー")
		log.Fatal(err)
  }

  err = Db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Println("DBマイグレーションエラー")
		log.Fatal(err)
	}
}