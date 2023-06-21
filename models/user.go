package models

import (
	"github.com/Kimoto-Norihiro/gin-api/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
}

func (u *User) CreateUser() error {
	db := database.GetDB()
	if db == nil {
		panic("failed to connect to the database.")
	}

	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func (u *User) GetUserById(id uint) error {
// 	result := database.GetDB().First(&u, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func (u *User) GetUser() error {
// 	result := database.GetDB().First(&u)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func (u *User) DeleteUser() error {
// 	result := database.GetDB().Delete(&u)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func (u *User) UpdateUser() error {
// 	result := database.GetDB().Save(&u)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }