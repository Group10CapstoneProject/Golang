package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
	"gorm.io/gorm"
)

type notificationRepositoryImpl struct {
	db *gorm.DB
}

// CreateNotification implements NotificationRepository
func (r *notificationRepositoryImpl) CreateNotification(body *model.Notification, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	return err
}

// DeleteNotification implements NotificationRepository
func (r *notificationRepositoryImpl) DeleteNotification(body *model.Notification, ctx context.Context) error {
	err := r.db.WithContext(ctx).Delete(body, body).Error
	return err
}

// FindNotification implements NotificationRepository
func (r *notificationRepositoryImpl) FindNotifications(ctx context.Context) ([]model.Notification, int, error) {
	notifications := []model.Notification{}
	var count int64
	err := r.db.WithContext(ctx).
		Preload("User").
		Order("id DESC").
		Find(&notifications).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, err
	}
	return notifications, int(count), nil
}

func NewNotificationRepository(database *gorm.DB) NotificationRepository {
	return &notificationRepositoryImpl{
		db: database,
	}
}
