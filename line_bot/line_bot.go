package line_bot

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewBotClient(secret string, access_token string) (*linebot.Client, error) {
	client, err := linebot.New(secret, access_token)
	if err != nil {
		return nil, fmt.Errorf("failed to create linebot client: %w", err)
	}
	return client, nil
}
