package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Trainer struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Name           string         `gorm:"type:varchar(255)"`
	Email          string         `gorm:"type:varchar(255)"`
	Phone          string         `gorm:"type:varchar(20)"`
	Dob            time.Time
	Gender         string `gorm:"type:varchar(5)"`
	Price          uint
	DailySlot      uint
	Description    string
	TrainerSkill   []TrainerSkill
	Picture        string
	TrainerBooking []TrainerBooking
}

type Skill struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `gorm:"type:varchar(255);uniqueIndex"`
	Description  string
	TrainerSkill []TrainerSkill
}

type TrainerSkill struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	TrainerID uint
	Trainer   Trainer
	SkillID   uint
	Skill     Skill
}

type TrainerBooking struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          uint
	User            User
	TrainerID       uint
	Trainer         Trainer
	ExpiredAt       time.Time
	ActivedAt       time.Time `gorm:"default:null"`
	Time            time.Time
	Status          StatusType `gorm:"type:enum('PENDING', 'WAITING', 'ACTIVE', 'INACTIVE', 'REJECT', 'DONE', 'CENCEL');column:status"`
	ProofPayment    string
	PaymentMethodID *uint
	PaymentMethod   PaymentMethod
	Code            uuid.UUID
	Total           uint
}
