package main

import (
	"fmt"
	"log"
	"net/http"

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
  client := &http.Client{}

  // line-bot
  bot, _ := linebot.New(
    "1f0cc8bdd3631e80428d94f001061e78",
    "8zuf8b9zJkAQVX+oM7jK2aY095e7G7E7CYJE2WthY6+bRhphEv++j6lwaZO4/psmlIImaqXg0tnvIzUdT+GOA995UAbOhCyycTZJCCucx11Iar6fpdSp9bMr8whf/nm55bYQJGchDU4Iv3Z8bIDT4AdB04t89/1O/w1cDnyilFU=",
    linebot.WithHTTPClient(client),
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

  // line-bot
  router.GET("/bot", func(c *gin.Context) {
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
      if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
        case *linebot.TextMessage:
          log.Println(message.Text)
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

  fmt.Println("Server running on port 8080")
  router.Run(":8080")
}