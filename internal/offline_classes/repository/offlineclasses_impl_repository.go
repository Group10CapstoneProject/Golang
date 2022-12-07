package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
)

type offlineclassRepositoryImpl struct {
	db *gorm.DB
}

// CreateOfflineClass implements OfflineClassRepository
func (r *offlineclassRepositoryImpl) CreateOfflineClass(body *model.OfflineClass, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if err := r.CheckOfflineClassIsDeleted(body); err == nil {
				return nil
			}
			return myerrors.ErrDuplicateRecord
		}
		return err
	}
	return nil
}

// CheckOfflineClassTypeIsDeleted implements OfflineClassRepository
func (r *offlineclassRepositoryImpl) CheckOfflineClassIsDeleted(body *model.OfflineClass) error {
	offlineclass := model.OfflineClass{}
	err := r.db.Where("title = ?", body.title).First(&model.OfflineClass{}).Error
	if err == nil {
		return myerrors.ErrDuplicateRecord
	}
	err = r.db.Unscoped().Where("title = ?", body.title).First(&offlineclass).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}
	body.ID = offlineclass.ID

	if err := r.UpdateOfflineClass(body, context.Background()); err != nil {
		return err
	}
	return nil
}

// DeleteOfflineClass implements OfflineClassRepository
func (r *offlineclassRepositoryImpl) DeleteOfflineClass(body *model.OfflineClass, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// FindOfflineClass implements OfflineClassRepository
func (r *offlineclassRepositoryImpl) FindOfflineClass(ctx context.Context) ([]model.OfflineClass, error) {
	offlineClass := []model.OfflineClass{}
	err := r.db.WithContext(ctx).Find(&offlineClass).Error
	return offlineClass, err
}

// FindOfflineClassById implements OfflineClassRepository
func (r *offlineclassRepositoryImpl) FindOfflineClassById(id uint, ctx context.Context) (*model.OfflineClass, error) {
	offlineclass := model.OfflineClass{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("title").
		Preload("slot").
		Preload("slot_ready").
		Preload("price").
		First(&offlineclass).Error
	return &offlineclass, err
}

// UpdateOfflineClass implements OfflineClassRepository
func (r *offlineclassRepositoryImpl) UpdateOfflineClass(body *model.OfflineClass, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(body).Updates(body)
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "Duplicate entry") {
			return myerrors.ErrDuplicateRecord
		}
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

func NewOfflineClassRepository(database *gorm.DB) OfflineClassRepository {
	return &offlineclassRepositoryImpl{
		db: database,
	}
}
