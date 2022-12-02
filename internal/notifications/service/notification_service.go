package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/notifications/dto"
)

type NotificationService interface {
	ReadNotification(notif *dto.NotificationReadRequest, ctx context.Context) error
	PullNotifications(ctx context.Context) (dto.NotificationResponse, error)
}
