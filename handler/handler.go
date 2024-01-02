package handler

import (
	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"github.com/gin-gonic/gin"
)

type IBotHandler interface {
	Response(c *gin.Context)
}

type ITodoHandler interface {
	CreateTodo(line_user_id string, name string) (*models.Todo, error)
	ListTodos(line_user_id string) ([]models.Todo, error)
	DeleteTodo(line_user_id string, name string) error
}
