package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string         `gorm:"type:varchar(255);unique"`
	Description   string
	PaymentNumber string `gorm:"type:varchar(255)"`
}
