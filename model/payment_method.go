package model

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string         `gorm:"type:varchar(255);uniqueIndex"`
	PaymentNumber string         `gorm:"type:varchar(255)"`
	Description   string
	Picture       string
}

type PaymentRequest struct {
	ID       uint           `validate:"required"`
	UserID   uint           `validate:"required"`
	FileName string         `validate:"required,image"`
	File     multipart.File `validate:"required,file"`
}
