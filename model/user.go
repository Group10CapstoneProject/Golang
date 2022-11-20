package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password  string         `gorm:"type:varchar(255);not null"`
	Role      int            `gorm:"not null"`
	MemberID  string
	SessionID uuid.UUID `gorm:"type:varchar(50)"`
}
