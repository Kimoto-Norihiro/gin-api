package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

type Usecase interface {
	CreateUser(c *gin.Context, m *models.User) error
}
