package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/model"
)

type NotificationReadRequest struct {
	UserID        uint
	TransactionID uint
	Title         string
}

func (u *NotificationReadRequest) ToModel() *model.Notification {
	return &model.Notification{
		UserID:        u.UserID,
		TransactionID: u.TransactionID,
		Title:         u.Title,
	}
}

type NotificationResource struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UserID          uint      `json:"user_id"`
	UserName        string    `json:"user_name"`
	TransactionID   uint      `json:"transaction_id"`
	TransactionType string    `json:"transaction_type"`
	Title           string    `json:"title"`
}

func (u *NotificationResource) FromModel(m *model.Notification) {
	u.ID = m.ID
	u.CreatedAt = m.CreatedAt
	u.UserID = m.UserID
	u.UserName = m.User.Name
	u.TransactionID = m.TransactionID
	u.TransactionType = m.TransactionType
	u.Title = m.Title
}

type NotificationResources []NotificationResource

func (u *NotificationResources) FromModel(m []model.Notification) {
	for _, each := range m {
		var resource NotificationResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

type NotificationResponse struct {
	Notifications NotificationResources `json:"notifications"`
	Count         uint                  `json:"count"`
}
