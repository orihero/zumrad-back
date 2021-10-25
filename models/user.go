package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	FirsName  string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Number    string         `json:"number"`
}
