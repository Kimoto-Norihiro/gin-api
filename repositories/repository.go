package repositories

import (
	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	CreateUser(c *gin.Context, m *models.User) error
}