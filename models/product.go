package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id             uint             `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	DeletedAt      gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
	Name           string           `json:"name"`
	Price          uint             `json:"price"`
	Model          string           `json:"model"`
	Description    string           `json:"description"`
	Sizes          Size             `json:"size"`
	Characteristic []Characteristic `json:"characteristic"`
	Color          []Color          `json:"color"`
}

type Color struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ProductId uint           `json:"product_id"`
	Name      string         `json:"name"`
	Hex       string         `json:"hex"`
	Image     string         `json:"image"`
}

type Characteristic struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ProductId uint           `json:"product_id"`
	Name      string         `json:"name"`
	Value     string         `json:"value"`
}

type Size struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ProductId uint           `json:"product_id"`
	Width     uint           `json:"width"`
	Height    uint           `json:"height"`
	Wide      uint           `json:"wide"`
}

type Category struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string
	Image     string
}
