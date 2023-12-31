package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-line-bot/database"
	"github.com/Kimoto-Norihiro/gin-line-bot/handlers"
	"github.com/Kimoto-Norihiro/gin-line-bot/line_bot"
)

func main() {
	botClient, err := line_bot.NewBotClient()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal(err)
	}
	bh := handlers.NewBotHandler(db, botClient)
	
	r := gin.Default()

	r.POST("/bot", bh.Response)

	r.Run(":" + os.Getenv("PORT"))
}
