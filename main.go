package main

import (
	"fmt"
	"os"

  "github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"

  "github.com/Kimoto-Norihiro/gin-line-bot/utils/bot"
)

func main() {
  router := gin.Default()

  router.POST("/bot", func(c *gin.Context) {
    events, err := Bot.ParseRequest(c.Request)
    if err != nil {
      if err == linebot.ErrInvalidSignature {
        c.Writer.WriteHeader(400)
      } else {
        c.Writer.WriteHeader(500)
      }
      return
    }

    for _, event := range events {
      if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
        case *linebot.TextMessage:
          switch {
          case strings.HasPrefix(message.Text, "登録"):
            var user models.User
            result := db.Where("user_id").First(&user)

            if result.Error == gorm.ErrRecordNotFound {
              // recode not exist
              user = models.User{
                UserID: event.Source.UserID,
                Todos: []models.Todo{
                  {Title: message.Text},
                },
              }
              db.Create(&user)
            } else {
              // recode exist 
              todo := models.Todo{
                UserID: user.UserID,
                Title: message.Text,
              }
              db.Create(&todo)
            }

            if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+"を登録したよ")).Do(); err != nil {
              log.Print(err)
            }
          case strings.HasPrefix(message.Text, "削除"):
          case strings.HasPrefix(message.Text, "一覧"):
          }
        }
      }
    }
  })

  fmt.Println("Server running on port "+os.Getenv("PORT"))
  router.Run(":"+os.Getenv("PORT"))
}