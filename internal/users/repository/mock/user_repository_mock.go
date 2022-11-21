package mock

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(model *model.User, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func (m *UserRepositoryMock) CheckUserIsEmpty(ctx context.Context) (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *UserRepositoryMock) FindUserByEmail(email *string, ctx context.Context) (*model.User, error) {
	args := m.Called()
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) FindUserByID(id *uint, ctx context.Context) (*model.User, error) {
	args := m.Called()
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) FindUsers(page *model.Pagination, ctx context.Context) ([]model.User, int, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Get(1).(int), args.Error(2)
}
