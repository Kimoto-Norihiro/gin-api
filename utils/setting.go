package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Setting struct {
	Port               string
	DatabaseUrl        string
	ChannelSecret      string
	ChannelAccessToken string
}

func LoadSetting() (*Setting, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, err
	}

	return &Setting{
		Port:               os.Getenv("PORT"),
		DatabaseUrl:        os.Getenv("DATABASE_URL"),
		ChannelSecret:      os.Getenv("CHANNEL_SECRET"),
		ChannelAccessToken: os.Getenv("CHANNEL_ACCESS_TOKEN"),
	}, nil
}
