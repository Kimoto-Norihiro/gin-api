package handler

import (
	"github.com/Kimoto-Norihiro/gin-line-bot/models"
	"github.com/Kimoto-Norihiro/gin-line-bot/utils/database"
)

func CreateUser(LineUserID string) (*models.User, error) {
	user := &models.User{
		LineUserID: LineUserID,
	}
	result := database.Db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func ShowUser(LineUserID string) (*models.User, error) {
	var user models.User
	result := database.Db.Where("line_user_id = ?", LineUserID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
