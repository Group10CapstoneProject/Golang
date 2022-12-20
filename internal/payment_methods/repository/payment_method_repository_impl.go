package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
)

type paymentMethodRepositoryImpl struct {
	db *gorm.DB
}

// CreatePaymentMethod implements PaymentMethodRepository
func (r *paymentMethodRepositoryImpl) CreatePaymentMethod(body *model.PaymentMethod, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			return myerrors.ErrDuplicateRecord
		}
		return err
	}
	return nil
}

// DeletePaymentMethod implements PaymentMethodRepository
func (r *paymentMethodRepositoryImpl) DeletePaymentMethod(body *model.PaymentMethod, ctx context.Context) error {
	err := r.db.WithContext(ctx).Preload("Member").Preload("OnlineClass").Preload("OfflineClass").Preload("Trainer").First(body).Error
	if err != nil {
		return err
	}
	if len(body.Member) > 0 || len(body.OnlineClass) > 0 || len(body.OfflineClass) > 0 || len(body.Trainer) > 0 {
		return myerrors.ErrRecordIsUsed
	}
	res := r.db.WithContext(ctx).Unscoped().Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// FindPaymentMethodById implements PaymentMethodRepository
func (r *paymentMethodRepositoryImpl) FindPaymentMethodById(id uint, ctx context.Context) (*model.PaymentMethod, error) {
	paymentMethods := model.PaymentMethod{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&paymentMethods).Error
	return &paymentMethods, err
}

// FindPaymentMethods implements PaymentMethodRepository
func (r *paymentMethodRepositoryImpl) FindPaymentMethods(access bool, ctx context.Context) ([]model.PaymentMethod, error) {
	paymentMethods := []model.PaymentMethod{}
	res := r.db.WithContext(ctx).Model(&model.PaymentMethod{})
	if !access {
		res.Where("id != ?", 0).Find(&paymentMethods)
	} else {
		res.Find(&paymentMethods)
	}
	err := res.Error
	if err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

// UpdatePaymentMethod implements PaymentMethodRepository
func (r *paymentMethodRepositoryImpl) UpdatePaymentMethod(body *model.PaymentMethod, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(body).Updates(body)
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "Error 1062:") {
			return myerrors.ErrDuplicateRecord
		}
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

func NewPaymentMethodRepository(database *gorm.DB) PaymentMethodRepository {
	return &paymentMethodRepositoryImpl{
		db: database,
	}
}
