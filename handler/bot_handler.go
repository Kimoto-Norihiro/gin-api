package handler

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"gorm.io/gorm"
)

type BotHandler struct {
	db      *gorm.DB
	botClient *linebot.Client
}

func NewBotHandler(db *gorm.DB, botClient *linebot.Client) *BotHandler {
	return &BotHandler{
		db: db, 
		botClient: botClient,
	}
}

func (h *BotHandler) Response(c *gin.Context) {
	events, err := h.botClient.ParseRequest(c.Request)
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
					// title := strings.TrimPrefix(message.Text, "登録　")
					// todo, err := todo_handler.Create(event, title)
					// if err != nil {
					// 	line_bot.BotReplyMessage(event, "登録に失敗しました。")
					// }
					// line_bot.BotReplyMessage(event, todo.Title+"を登録しました。")

				case strings.HasPrefix(message.Text, "削除"):
					// title := strings.TrimPrefix(message.Text, "削除　")
					// err := todo_handler.Delete(event, title)
					// if err != nil {
					// 	line_bot.BotReplyMessage(event, "削除に失敗しました。")
					// }
					// line_bot.BotReplyMessage(event, title+"を削除しました。")

				case strings.HasPrefix(message.Text, "一覧"):
					// todos, err := todo_handler.Index(event)
					// if err != nil {
					// 	line_bot.BotReplyMessage(event, "一覧の取得に失敗しました。")
					// }
					// var titles []string
					// for _, todo := range todos {
					// 	titles = append(titles, todo.Title)
					// }
					// log.Println(strings.Join(titles, "\n"))
					// line_bot.BotReplyMessage(event, "一覧\n"+strings.Join(titles, "\n"))

				default:
					log.Println("default")
					// res, err := h.lineBot.ReplyMessage(event, "登録、削除、一覧のいずれかを入力してください。")
					// if err != nil {
					// 	log.Println(err)
					// }
					// log.Println(res)
				}
			}
		}
	}
}
