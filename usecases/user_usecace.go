package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-api/repositories"
	"github.com/Kimoto-Norihiro/gin-api/models"
)

type UserUsecase struct {
	repository repositories.Repository
}

func NewUserUsecase(repository repositories.Repository) *UserUsecase {
	return &UserUsecase{repository: repository}
}

// CreateUser is a function to create a user.
func (uu *UserUsecase) CreateUser(c *gin.Context, m *models.User) error {
	return uu.repository.CreateUser(c, m)
}