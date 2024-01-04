package usecases

import (
	"context"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"github.com/Kimoto-Norihiro/gin-line-bot/repository"
	"gorm.io/gorm"
)

// user, err := user_handler.Show(event.Source.UserID)

type TodoUsecase struct {
	db   *gorm.DB
	repo repository.ITodoRepository
}

func NewTodoUsecase(db *gorm.DB, repo repository.ITodoRepository) *TodoUsecase {
	return &TodoUsecase{
		db:   db,
		repo: repo,
	}
}

func (h *TodoUsecase) CreateTodo(ctx context.Context, line_user_id string, title string) error {
	err := h.db.Transaction(func(tx *gorm.DB) error {
		return h.repo.Create(ctx, tx, line_user_id, title)
	})
	if err != nil {
		return err
	}

	return nil
}

func (h *TodoUsecase) ListTodos(ctx context.Context, line_user_id string) (*[]models.Todo, error) {
	return h.repo.List(ctx, h.db, line_user_id)
}

func (h *TodoUsecase) DeleteTodo(ctx context.Context, line_user_id string, title string) error {
	err := h.db.Transaction(func(tx *gorm.DB) error {
		return h.repo.Delete(ctx, tx, line_user_id, title)
	})
	if err != nil {
		return err
	}
	return nil
}
