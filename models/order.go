package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Client    User           `json:"client"`
	Quantity  uint           `json:"quantity"`
	Note      string         `json:"note"`
	Status    string         `json:"stasus"`
}

const (
	STATUS_PROCESSING  = "processing"
	STATUS_PREPARING   = "preparing"
	STATUS_ONTHEWAY    = "on_the_way"
	STATUS_DELIVERD    = "delivered"
	STATUS_COMMENTERED = "commented"
)

const (
	PAYMENT_TYPE_CARD        = "card"
	PAYMENT_TYPE_CASH        = "cash"
	PAYMENT_TYPE_TRANSACTION = "transaction"
)
