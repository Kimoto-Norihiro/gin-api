package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-line-bot/handlers/bot_handler"
	"github.com/Kimoto-Norihiro/gin-line-bot/utils/database"
	"github.com/Kimoto-Norihiro/gin-line-bot/utils/line_bot"
)

func init() {
  database.Init()
  line_bot.Init()
}

func main() {
  router := gin.Default()

  router.POST("/bot", func(c *gin.Context) {
    bot_handler.MainHandler(c)
  })

  log.Println("Server running on port "+os.Getenv("PORT"))
  router.Run(":"+os.Getenv("PORT"))
}