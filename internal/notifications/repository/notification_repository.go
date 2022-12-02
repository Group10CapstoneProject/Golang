package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type NotificationRepository interface {
	CreateNotification(body *model.Notification, ctx context.Context) error
	FindNotifications(ctx context.Context) ([]model.Notification, int, error)
	DeleteNotification(body *model.Notification, ctx context.Context) error
}
