package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(255);not null;uniqueIndex"`
	Email     string         `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password  string         `gorm:"type:varchar(255);not null"`
	MemberID  string
}
