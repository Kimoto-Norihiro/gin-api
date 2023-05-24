package models

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  UserID string `json:"user_id"`
  Todos  []Todo `json:"todos"`
}

type Todo struct {
  gorm.Model
  Title   string `json:"title"`
  UserID  string `json:"user_id"`
}