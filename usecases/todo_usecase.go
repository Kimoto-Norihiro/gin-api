package usecases

import (
	"errors"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"github.com/Kimoto-Norihiro/gin-line-bot/repository"
	"gorm.io/gorm"
)

// user, err := user_handler.Show(event.Source.UserID)

type TodoUsecase struct {
	db   *gorm.DB
	repo repository.ITodoRepository
}

func NewTodoUsecase(repo repository.ITodoRepository) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (h *TodoUsecase) CreateTodo(line_user_id string, title string) (*models.Todo, error) {
	return nil, errors.New("not implemented")
}

func (h *TodoUsecase) ListTodos(line_user_id string) (*[]models.Todo, error) {
	return nil, errors.New("not implemented")
}

func (h *TodoUsecase) DeleteTodo(line_user_id string, title string) error {
	return errors.New("not implemented")
}
