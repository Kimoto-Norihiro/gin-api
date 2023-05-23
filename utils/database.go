package utils

import (
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"github.com/Kimoto-Norihiro/gin-api/models"
)

func init() {
	var Db *gorm.DB
	dsn := os.Getenv("CLEARDB_DATABASE_URL")
	
  Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  Db.AutoMigrate(&models.User{}, &models.Todo{})
}