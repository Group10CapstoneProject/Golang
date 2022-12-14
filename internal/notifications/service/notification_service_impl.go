package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/notifications/dto"
	notifRepo "github.com/Group10CapstoneProject/Golang/internal/notifications/repository"
)

type notificationServiceImpl struct {
	notificationRepository notifRepo.NotificationRepository
}

// PullNotifications implements NotificationService
func (s *notificationServiceImpl) PullNotifications(ctx context.Context) (dto.NotificationResponse, error) {
	notifications, count, err := s.notificationRepository.FindNotifications(ctx)
	if err != nil {
		return dto.NotificationResponse{}, err
	}
	var result dto.NotificationResources
	result.FromModel(notifications)

	return dto.NotificationResponse{Notifications: result, Count: uint(count)}, nil
}

// ReadNotification implements NotificationService
func (s *notificationServiceImpl) ReadNotification(notif *dto.NotificationReadRequest, ctx context.Context) error {
	err := s.notificationRepository.DeleteNotification(notif.ToModel(), ctx)
	return err
}

func NewNotificationService(notificationRepository notifRepo.NotificationRepository) NotificationService {
	return &notificationServiceImpl{
		notificationRepository: notificationRepository,
	}
}
