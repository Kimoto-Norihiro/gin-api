package line_bot

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewBotClient() (*linebot.Client, error) {
	client, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
