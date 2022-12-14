package model

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          uint
	User            User
	TransactionID   uint
	TransactionType string `gorm:"type:varchar(50)"`
	Title           string `gorm:"type:varchar(20)"`
}
