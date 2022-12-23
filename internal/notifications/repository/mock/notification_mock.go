package mock

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/stretchr/testify/mock"
)

type NotificationRepositoryMock struct {
	mock.Mock
}

func (m *NotificationRepositoryMock) CreateNotification(body *model.Notification, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func (m *NotificationRepositoryMock) FindNotifications(ctx context.Context) ([]model.Notification, int, error) {
	args := m.Called()
	return args.Get(0).([]model.Notification), args.Int(1), args.Error(2)
}

func (m *NotificationRepositoryMock) DeleteNotification(body *model.Notification, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}
