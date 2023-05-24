package utils

import (
	"os"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	// "github.com/joho/godotenv"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

var Db *gorm.DB

func init() {
	var err error
	// err = godotenv.Load()
  // if err != nil {
  //   log.Println("Error loading .env file")
  // }

	dsn := os.Getenv("CLEARDB_DATABASE_URL")
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
		log.Println(err)
  }
	Db = db

  err = Db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal(err)
	}
}