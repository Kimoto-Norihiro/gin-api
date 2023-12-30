package line_bot

import (
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineBot struct {
	client *linebot.Client
}

func NewLineBot() (*LineBot, error) {
	client, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		return nil, err
	}
	return &LineBot{client: client}, nil
}

func (b *LineBot) ReplyMessage(replyToken string, message linebot.SendingMessage) (*linebot.BasicResponse, error) {
	return b.client.ReplyMessage(replyToken, message).Do()
}

func (b *LineBot) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return b.client.ParseRequest(r)
}
