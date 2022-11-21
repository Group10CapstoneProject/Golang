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

// FindUserByEmail implements UserRepository
func (r *userRepositoryImpl) FindUserByEmail(email *string, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", *email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserById implements UserRepository
func (r *userRepositoryImpl) FindUserByID(id *uint, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", *id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUsers implements UserRepository
func (r *userRepositoryImpl) FindUsers(page *model.Pagination, ctx context.Context) ([]model.User, int, error) {
	var users []model.User
	var count int64
	if page.Q != "" {
		res := r.db.WithContext(ctx).Model(&model.User{}).Where("name = ? OR email = ?", "%"+page.Q+"%", "%"+page.Q+"%").
			Offset(page.Limit * page.Page).Limit(page.Limit).Find(&users)
		if res.Error != nil {
			return nil, 0, res.Error
		}
		res.Count(&count)
		return users, int(count), nil
	}
	res := r.db.WithContext(ctx).Model(&model.User{}).
		Offset(page.Limit * page.Page).Limit(page.Limit).Find(&users)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	res.Count(&count)
	return users, int(count), nil
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: database,
	}
}
