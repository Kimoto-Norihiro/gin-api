package repository

import "github.com/Kimoto-Norihiro/gin-line-bot/models"

type ITodoRepository interface {
	Create(line_user_id string, name string) (*models.Todo, error)
	List(line_user_id string) ([]models.Todo, error)
	Delete(line_user_id string, name string) error
}
