package usecases

import (
	"context"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

type ITodoHandler interface {
	CreateTodo(ctx context.Context, line_user_id string, name string) error
	ListTodos(ctx context.Context, line_user_id string) (*[]models.Todo, error)
	DeleteTodo(ctx context.Context, line_user_id string, name string) error
}
