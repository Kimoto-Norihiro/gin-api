package line_bot

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var Bot *linebot.Client

func Init() {
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
