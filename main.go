package main

import (
	"fmt"
	"log"
	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

type User struct {
  ID        uint   `json:"id"`
  Name      string `json:"name"`
}

func main() {
  router := gin.Default()

  // line-bot
  bot, _ := linebot.New(
    os.Getenv("CHANNEL_SECRET"),
    os.Getenv("CHANNEL_ACCESS_TOKEN"),
  )

  // db
  // dsn := "root:password@tcp(localhost:3306)/gin_api?charset=utf8mb4&parseTime=True&loc=Local"
  // db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  // if err != nil {
  //   panic(err)
  // }

  // db.AutoMigrate(&User{})

  // user := User{Name: "Genie"}
  // db.Create(&user)

  router.POST("/bot", func(c *gin.Context) {
    events, err := bot.ParseRequest(c.Request)
    if err != nil {
      if err == linebot.ErrInvalidSignature {
        c.Writer.WriteHeader(400)
      } else {
        c.Writer.WriteHeader(500)
      }
      return
    }

    for _, event := range events {
      log.Println(event)
      log.Println(event.Type)
      log.Println(event.Message)
      log.Println(event.ReplyToken)
      log.Println(event.Source.UserID)
      if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
        case *linebot.TextMessage:
          log.Println(message.Text)
          if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
            log.Print(err)
          }
        }
      }
    }
  })

  // user設定
  // router.GET("/users", func(c *gin.Context) {
  //   var users []User
  //   db.Find(&users)
  //   c.JSON(200, gin.H{
  //     "users": users,
  //   })
  // })

  // router.POST("/users", func(c *gin.Context) {
  //   var user User
  //   c.BindJSON(&user)
  //   db.Create(&user)
  //   c.JSON(200, user)
  // })

  fmt.Println("Server running on port "+os.Getenv("PORT"))
  router.Run(":"+os.Getenv("PORT"))
}