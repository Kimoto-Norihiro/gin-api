package handlers

import (
	"errors"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	// "gorm.io/gorm"

	"github.com/Kimoto-Norihiro/gin-line-bot/models"
)

func resolveUser(event *linebot.Event) (*models.User, error) {
	// user, err := user_handler.Show(event.Source.UserID)
	// if err == gorm.ErrRecordNotFound {
	// 	user, err = user_handler.Create(event.Source.UserID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// } else if err != nil {
	// 	return nil, err
	// }
	return nil, errors.New("not implemented")
}

func CreateHandler(event *linebot.Event, title string) (*models.Todo, error) {
	// user, err := resolveUser(event)
	// if err != nil {
	// 	return models.Todo{}, err
	// }
	// todo := models.Todo{
	// 	UserID: user.ID,
	// 	Title:  title,
	// }
	// result := database.Db.Create(&todo)
	// if result.Error != nil {
	// 	return todo, result.Error
	// }
	// return todo, nil
	return nil, errors.New("not implemented")
}

func Index(event *linebot.Event) ([]models.Todo, error) {
	// user, err := resolveUser(event)
	// if err != nil {
	// 	return nil, err
	// }

	// var todos []models.Todo
	// result := database.Db.Where("user_id = ?", user.ID).Find(&todos)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return todos, nil
	return nil, errors.New("not implemented")
}

func Delete(event *linebot.Event, title string) error {
	// user, err := resolveUser(event)
	// if err != nil {
	// 	return err
	// }
	// result := database.Db.Where("user_id = ? AND title = ?", user.ID, title).Delete(&models.Todo{})
	// if result.Error != nil {
	// 	return result.Error
	// }
	// return nil
	return errors.New("not implemented")
}
