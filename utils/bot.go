package utils

import (
  "os"
  "github.com/line/line-bot-sdk-go/v7/linebot"
)

var Bot *linebot.Client

func init() {
  var err error
  Bot, err = linebot.New(
    os.Getenv("CHANNEL_SECRET"),
    os.Getenv("CHANNEL_ACCESS_TOKEN"),
  )
  if err != nil {
    panic(err)
  }
}

func BotReplyMessage(event *linebot.Event, message string) {
  _, err := Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
  if err != nil {
    panic(err)
  }
}
