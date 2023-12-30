package models

import (
	"gorm.io/gorm"
)

// User モデル
type User struct {
	gorm.Model
	LineUserID string `json:"line_user_id"`
	Todos      []Todo `json:"todos"`
}

// Todo モデル
type Todo struct {
	ID int 
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user" gorm:"foreignkey:UserID"` // User への外部キー
}
