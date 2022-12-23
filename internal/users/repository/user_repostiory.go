package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type UserRepository interface {
	CreateUser(model *model.User, ctx context.Context) error
	CheckUserIsEmpty(ctx context.Context) (bool, error)
	FindUserByEmail(email *string, ctx context.Context) (*model.User, error)
	FindUserByID(id *uint, ctx context.Context) (*model.User, error)
	FindUsers(page *model.Pagination, role string, ctx context.Context) ([]model.User, int, error)
	UpdateUser(user *model.User, ctx context.Context) error
	DeleteUser(user *model.User, ctx context.Context) error
}
