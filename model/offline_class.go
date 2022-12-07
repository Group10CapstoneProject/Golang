package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OfflineClass struct {
	ID                     uint `gorm:"primarykey"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              gorm.DeletedAt `gorm:"index"`
	Title                  string         `gorm:"type:varchar(50)"`
	Time                   time.Time
	Duration               uint
	Slot                   uint
	SlotBooked             uint
	Price                  uint
	Picture                string
	TrainerID              uint
	Description            string `gorm:"type:varchar(255)"`
	Location               string `gorm:"type:varchar(255)"`
	OfflineClassCategoryID uint
	OfflineClassCategory   OfflineClassCategory
	OfflineClassBooking    []OfflineClassBooking
}

type OfflineClassCategory struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `gorm:"type:varchar(255);uniqueIndex"`
	Description  string         `gorm:"type:varchar(255)"`
	Picture      string
	OfflineClass []OfflineClass
}

type OfflineClassBooking struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          uint
	User            User
	OfflineClassID  uint
	OfflineClass    OfflineClass
	ExpiredAt       time.Time
	ActivedAt       time.Time  `gorm:"default:null"`
	Status          StatusType `gorm:"type:enum('PENDING', 'WAITING', 'ACTIVE', 'INACTIVE', 'REJECT', 'DONE', 'CENCEL');column:status"`
	ProofPayment    string
	PaymentMethodId uint
	PaymentMethod   PaymentMethod
	Code            uuid.UUID
	Total           uint
}
