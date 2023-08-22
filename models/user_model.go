package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" validate:"required"`
	Email     string         `json:"email" validate:"required"`
	Password  string         `json:"-" validate:"required"`
	Address   string         `json:"address" validate:"required"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index, column:"deleted_at""`
}
