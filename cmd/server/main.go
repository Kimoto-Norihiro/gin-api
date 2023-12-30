package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-line-bot/database"
	"github.com/Kimoto-Norihiro/gin-line-bot/handler"
)

func init() {
	database.Init()
}

func main() {
	r := gin.Default()

	r.POST("/bot", handler.Handler)

	log.Println("Server running on port " + os.Getenv("PORT"))
	r.Run(":" + os.Getenv("PORT"))
}
