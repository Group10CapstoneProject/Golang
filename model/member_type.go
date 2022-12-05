package model

import (
	"time"

	"gorm.io/gorm"
)

type MemberType struct {
	ID                 uint `gorm:"primarykey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Name               string         `gorm:"type:varchar(255);unique"`
	Price              uint
	Description        string
	Picture            string
	AccessOfflineClass *bool
	AccessOnlineClass  *bool
	AccessTrainer      *bool
	AccessGym          *bool
}
