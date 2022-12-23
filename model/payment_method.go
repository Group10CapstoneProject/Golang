package model

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID            *uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string         `gorm:"type:varchar(255)"`
	PaymentNumber string         `gorm:"type:varchar(255)"`
	Description   string
	Picture       string
	Member        []Member
	OnlineClass   []OnlineClassBooking
	OfflineClass  []OfflineClassBooking
	Trainer       []TrainerBooking
}

type PaymentRequest struct {
	ID       uint           `validate:"required"`
	UserID   uint           `validate:"required"`
	FileName string         `validate:"required,image"`
	File     multipart.File `validate:"required,file"`
}
