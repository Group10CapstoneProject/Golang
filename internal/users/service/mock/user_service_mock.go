package mock

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	"github.com/Group10CapstoneProject/Golang/model"
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

func (m *UserServiceMock) CreateSuperadmin(superadmin *model.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *UserServiceMock) FindUsers(page model.Pagination, role string, ctx context.Context) (*dto.PageResponse, error) {
	args := m.Called()
	return args.Get(0).(*dto.PageResponse), args.Error(1)
}

func (m *UserServiceMock) FindUser(userId *uint, ctx context.Context) (*dto.UserResponse, error) {
	args := m.Called()
	return args.Get(0).(*dto.UserResponse), args.Error(1)
}
