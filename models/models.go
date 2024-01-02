package models

type Todo struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	LineUserID string `json:"line_user_id"`
}
