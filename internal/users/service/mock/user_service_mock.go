package mock

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) CreateUser(user *dto.NewUser, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func (m *UserServiceMock) CreateAdmin(user *dto.NewUser, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func (m *UserServiceMock) CreateSuperadmin() error {
	args := m.Called()
	return args.Error(0)
}
