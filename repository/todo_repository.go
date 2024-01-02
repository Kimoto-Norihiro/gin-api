package repository

import (
	"errors"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Create(line_user_id string, name string) error {
	err := r.db.Create(&models.Todo{
		LineUserID: line_user_id,
		Name:       name,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) List(line_user_id string) ([]models.Todo, error) {
	return nil, errors.New("not implemented")
}

func (r *TodoRepository) Delete(line_user_id string, title string) error {
	return errors.New("not implemented")
}
