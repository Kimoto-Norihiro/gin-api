package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-line-bot/database"
	"github.com/Kimoto-Norihiro/gin-line-bot/handler"
	"github.com/Kimoto-Norihiro/gin-line-bot/line_bot"
	"github.com/Kimoto-Norihiro/gin-line-bot/utils"
)

func main() {
	setting, err := utils.LoadSetting()
	if err != nil {
		fmt.Println(err)
	}
	botClient, err := line_bot.NewBotClient(setting.ChannelSecret, setting.ChannelAccessToken)
	if err != nil {
		fmt.Println(err)
	}
	db, err := database.ConnDB(setting.DatabaseUrl)
	if err != nil {
		fmt.Println(err)
	}

	bh := handler.NewBotHandler(db, botClient)

	r := gin.Default()

	r.POST("/bot", bh.Response)
	r.Run(":" + setting.Port)
}
