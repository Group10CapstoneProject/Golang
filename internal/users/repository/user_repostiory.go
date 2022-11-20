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
}
