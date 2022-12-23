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

// DeleteUser implements UserRepository
func (r *userRepositoryImpl) DeleteUser(user *model.User, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(user)
	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// UpdateUser implements UserRepository
func (r *userRepositoryImpl) UpdateUser(user *model.User, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(user).Updates(user)
	err := res.Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			return myerrors.ErrEmailAlredyExist
		}
		return err
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// CreateUser implements UserRepostiory
func (r *userRepositoryImpl) CreateUser(model *model.User, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(model).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
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
func (r *userRepositoryImpl) FindUsers(page *model.Pagination, role string, ctx context.Context) ([]model.User, int, error) {
	var users []model.User
	var count int64
	offset := (page.Limit * page.Page) - page.Limit

	if page.Q != "" {
		err := r.db.WithContext(ctx).Model(&model.User{}).
			Where("(name LIKE ? OR email LIKE ?) AND role = ?", "%"+page.Q+"%", "%"+page.Q+"%", role).
			Offset(offset).
			Limit(page.Limit).
			Order("id DESC").
			Find(&users).
			Error
		if err != nil {
			return nil, 0, err
		}

		err = r.db.WithContext(ctx).Model(&model.User{}).Where("(name LIKE ? OR email LIKE ?) AND role = ?", "%"+page.Q+"%", "%"+page.Q+"%", role).Count(&count).Error
		if err != nil {
			return nil, 0, err
		}

		return users, int(count), nil
	}

	err := r.db.WithContext(ctx).Model(&model.User{}).
		Where("role = ?", role).
		Offset(offset).
		Limit(page.Limit).
		Order("id DESC").
		Find(&users).
		Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Model(&model.User{}).Where("role = ?", role).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return users, int(count), nil
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: database,
	}
}
