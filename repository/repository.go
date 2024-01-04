package repository

import (
	"context"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(ctx context.Context, db *gorm.DB, line_user_id string, name string) error
	List(ctx context.Context, db *gorm.DB, line_user_id string) (*[]models.Todo, error)
	Delete(ctx context.Context, tx *gorm.DB, line_user_id string, name string) error
}
