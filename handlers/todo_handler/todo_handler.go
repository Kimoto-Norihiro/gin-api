package todo_handler

import (
  "github.com/line/line-bot-sdk-go/v7/linebot"
  "gorm.io/gorm"

  "github.com/Kimoto-Norihiro/gin-line-bot/models"
  "github.com/Kimoto-Norihiro/gin-line-bot/utils"
  "github.com/Kimoto-Norihiro/gin-line-bot/handlers/user_handler"
)

func resolveUser(event *linebot.Event) (*models.User, error) {
  user, err := user_handler.Show(event.Source.UserID)
  if err == gorm.ErrRecordNotFound {
    user, err = user_handler.Create(event.Source.UserID)
    if err != nil {
        return nil, err
    }
  } else if err != nil {
    return nil, err
  }
  return user, nil
}

func Create(event *linebot.Event, title string) (models.Todo, error) {
  user, err := resolveUser(event)
  if err != nil {
    return models.Todo{}, err
  }
  todo := models.Todo{
    UserID: user.UserID,
    Title: title,
  }
  result := utils.Db.Create(&todo)
  if result.Error != nil {
    return todo, result.Error
  }
  return todo, nil
}

func Index(event *linebot.Event) ([]models.Todo, error) {
  user, err := resolveUser(event)
  if err != nil {
    return nil, err
  }
  todos := user.Todos
  return todos, nil
}

func Delete(event *linebot.Event, title string) error {
  user, err := resolveUser(event)
  if err != nil {
    return err
  }
  result := utils.Db.Where("user_id = ? AND title = ?", user.ID, title).Delete(&models.Todo{})
  if result.Error != nil {
    return result.Error
  }
  return nil
}
