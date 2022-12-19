package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	"github.com/Group10CapstoneProject/Golang/model"
)

type UserService interface {
	CreateUser(user *dto.NewUser, ctx context.Context) error
	CreateAdmin(user *dto.NewUser, ctx context.Context) error
	FindUser(userId *uint, ctx context.Context) (*dto.UserResponse, error)
	FindUsers(page model.Pagination, role string, ctx context.Context) (*dto.PageResponse, error)
	CreateSuperadmin(superadmin *model.User) error
}
