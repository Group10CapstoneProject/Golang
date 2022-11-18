package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

// CreateUser implements UserRepostiory
func (r *userRepositoryImpl) CreateUser(model *model.User, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(model).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return myerrors.ErrEmailAlredyExist
		}
		return err
	}
	return nil
}

// CheckUserIsEmpty implements UserRepostiory
func (r *userRepositoryImpl) CheckUserIsEmpty(ctx context.Context) (bool, error) {
	err := r.db.WithContext(ctx).First(&model.User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: database,
	}
}
