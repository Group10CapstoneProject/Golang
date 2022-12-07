package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OfflineClass struct {
	ID           uuid.UUID `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	title        string         `gorm:"type:varchar(50);not null"`
	time         time.Time
	duration     time.Timer
	slot         uint
	slot_ready   uint
	price        uint32 `gorm:"type:uint;not null"`
	picture_path string
	trainer_id   uuid.UUID
	description  string `gorm:"type:varchar(255)"`
}
