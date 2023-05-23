package todo

import (
	"log"

	"gorm.io/gorm"
	"github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/Kimoto-Norihiro/gin-api/models"
	"github.com/Kimoto-Norihiro/gin-api/utils/database"
	"github.com/Kimoto-Norihiro/gin-api/utils/bot"
)

func create(event *linebot.Event, message *linebot.TextMessage) {
	var user models.User
	result := database.Db.Where("user_id").First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		// recode not exist
		user = models.User{
			UserID: event.Source.UserID,
			Todos: []models.Todo{
				{Title: message.Text},
			},
		}
		database.Db.Create(&user)
	} else {
		// recode exist 
		todo := models.Todo{
			UserID: user.UserID,
			Title: message.Text,
		}
		database.Db.Create(&todo)
	}

	var err error
	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+"を登録したよ")).Do(); err != nil {
		log.Print(err)
	}
}

func index() {
	// get all todos
}