package repositories

import (
	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(c *gin.Context, m *models.User) error {
	return r.db.Create(m).Error
}