package user_handler

import (
  "github.com/Kimoto-Norihiro/gin-line-bot/models"
  "github.com/Kimoto-Norihiro/gin-line-bot/utils/database"
)

func Create(UserID string) (*models.User, error) {
  user := &models.User{
    UserID: UserID,
  }
  result := database.Db.Create(user)
  if result.Error != nil {
    return nil, result.Error
  }
  return user, nil
}

func Show(UserID string) (*models.User, error) {
  var user models.User
  result := database.Db.Where("user_id = ?", UserID).First(&user)
  if result.Error != nil {
    return nil, result.Error
  }
  return &user, nil
}