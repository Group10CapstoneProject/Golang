package model

import (
	"time"

	"gorm.io/gorm"
)

type OnlineClass struct {
	ID                    uint `gorm:"primarykey"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	Title                 string         `gorm:"type:varchar(255)"`
	Link                  string         `gorm:"type:varchar(255)"`
	Path                  string         `gorm:"type:varchar(255)"`
	Price                 uint
	Description           string
	OnlineClassCategoryID uint
	OnlineClassCategory   OnlineClassCategory
	Tools                 string `gorm:"type:varchar(255)"`
	TargetArea            string `gorm:"type:varchar(255)"`
	Duration              uint
	TrainerID             uint
	Trainer               Trainer
	Level                 string `gorm:"type:varchar(255)"`
	Picture               string
}

type OnlineClassCategory struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"type:varchar(255);uniqueIndex"`
	Description string
	Picture     string
	OnlineClass []OnlineClass
}

type OnlineClassBooking struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          uint
	User            User
	OnlineClassID   uint
	OnlineClass     OnlineClass
	ExpiredAt       time.Time
	ActivedAt       time.Time `gorm:"default:null"`
	Duration        uint
	Status          StatusType `gorm:"type:enum('PENDING', 'WAITING', 'ACTIVE', 'INACTIVE', 'REJECT', 'DONE', 'CENCEL');column:status"`
	ProofPayment    string
	PaymentMethodId uint
	PaymentMethod   PaymentMethod
	Total           uint
}
