package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	
  "github.com/Kimoto-Norihiro/gin-line-bot/handlers/bot_handler"
)

func main() {
  router := gin.Default()

  router.POST("/bot", func(c *gin.Context) {
    bot_handler.MainHandler(c)
  })

  fmt.Println("Server running on port "+os.Getenv("PORT"))
  router.Run(":"+os.Getenv("PORT"))
}