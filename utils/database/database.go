package database

import (
	"os"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	// "github.com/joho/godotenv"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

var Db *gorm.DB

func Init() {
	// var err error
	// err = godotenv.Load()
  // if err != nil {
  //   log.Println("Error loading .env file")
  // }

	dsn := os.Getenv("DATABASE_URL")

	log.Println(dsn)
	
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
		log.Println("DB接続エラー")
		log.Println(err)
  }
	Db = db

  err = Db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Println("DBマイグレーションエラー")
		log.Fatal(err)
	}
}