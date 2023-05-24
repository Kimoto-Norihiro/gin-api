package line_bot

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
  // "github.com/joho/godotenv"
)

var Bot *linebot.Client

func Init() {
  // var err error
	// err = godotenv.Load()
  // if err != nil {
  //   log.Println("Error loading .env file")
  // }

  bot, err := linebot.New(
    os.Getenv("CHANNEL_SECRET"),
    os.Getenv("CHANNEL_ACCESS_TOKEN"),
  )
  Bot = bot
  if err != nil {
    log.Println("LINE BOT接続エラー")
    log.Fatal(err)
  }
}

func BotReplyMessage(event *linebot.Event, message string) {
  _, err := Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
  if err != nil {
    panic(err)
  }
}
