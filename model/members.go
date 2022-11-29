package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          uint
	User            User
	MemberTypeID    uint
	MemberType      MemberType
	ExpiredAt       time.Time
	Duration        uint
	Status          StatusType `gorm:"type:enum('PENDING', 'WAITING', 'ACTIVE', 'INACTIVE', 'REJECT', 'DONE', 'CENCEL');column:status"`
	ProofPayment    string     `gorm:"type:varchar(50)"`
	PaymentMethodId uint
	PaymentMethod   PaymentMethod
	Total           uint
	Code            uuid.UUID
}
