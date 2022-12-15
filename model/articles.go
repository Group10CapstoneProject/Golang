package model

import (
	"time"

	"gorm.io/gorm"
)

type Articles struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string         `gorm:"type:varchar(255)"`
	Description string         `gorm:"type:varchar(255)"`
	Picture     string         `gorm:"type:varchar(255)"`
	Content     string         `gorm:"type:varchar(255)"`
	Pinned      *bool
}
