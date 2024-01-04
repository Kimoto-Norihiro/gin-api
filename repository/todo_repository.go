package repository

import (
	"context"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"gorm.io/gorm"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) Create(ctx context.Context, tx *gorm.DB, line_user_id string, name string) error {
	todo := &models.Todo{
		LineUserID: line_user_id,
		Name:       name,
	}
	return tx.Create(todo).Error
}

func (r *TodoRepository) List(ctx context.Context, db *gorm.DB, line_user_id string) ([]models.Todo, error) {
	var todos []models.Todo
	if err := db.Where("line_user_id = ?", line_user_id).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) Delete(ctx context.Context, tx *gorm.DB, line_user_id string, title string) error {
	return tx.Where("line_user_id = ? AND title = ?", line_user_id, title).Delete(&models.Todo{}).Error
}
