package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
)

type UserService interface {
	CreateUser(user *dto.NewUser, ctx context.Context) error
	CreateAdmin(user *dto.NewUser, ctx context.Context) error
	CreateSuperadmin() error
}
