package handler

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

// func resolveUser(event *linebot.Event) (*models.User, error) {
// user, err := user_handler.Show(event.Source.UserID)
// if err == gorm.ErrRecordNotFound {
// 	user, err = user_handler.Create(event.Source.UserID)
// 	if err != nil {
// 		return nil, err
// 	}
// } else if err != nil {
// 	return nil, err
// }
// 	return nil, errors.New("not implemented")
// }

type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{
		db: db,
	}
}

func (h *TodoHandler) CreateTodo(line_user_id string, title string) (*models.Todo, error) {
	return nil, errors.New("not implemented")
}

func (h *TodoHandler) ListTodos(line_user_id string) (*[]models.Todo, error) {
	return nil, errors.New("not implemented")
}

func (h *TodoHandler) DeleteTodo(line_user_id string, title string) error {
	return errors.New("not implemented")
}
