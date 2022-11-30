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
		if strings.Contains(err.Error(), "Duplicate entry") {
			if err := r.CheckPaymentMethodIsDeleted(body); err == nil {
				return nil
			}
			return myerrors.ErrDuplicateRecord
		}
		return err
	}
	return nil
}

// CheckPaymentMethodIsDeleted implements PaymentMethodRepository
func (r *paymentMethodRepositoryImpl) CheckPaymentMethodIsDeleted(body *model.PaymentMethod) error {
	paymentMethod := model.PaymentMethod{}
	err := r.db.Where("name = ?", body.Name).First(&model.PaymentMethod{}).Error
	if err == nil {
		return myerrors.ErrDuplicateRecord
	}
	err = r.db.Unscoped().Where("name = ?", body.Name).First(&paymentMethod).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}
	body.ID = paymentMethod.ID

	if err := r.UpdatePaymentMethod(body, context.Background()); err != nil {
		return err
	}
	return nil
}

// DeletePaymentMethod implements PaymentMethodRepository
func (r *paymentMethodRepositoryImpl) DeletePaymentMethod(body *model.PaymentMethod, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
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
func (r *paymentMethodRepositoryImpl) FindPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error) {
	paymentMethods := []model.PaymentMethod{}
	err := r.db.WithContext(ctx).Find(&paymentMethods).Error
	return paymentMethods, err
}

// UpdatePaymentMethod implements PaymentMethodRepository
func (r *paymentMethodRepositoryImpl) UpdatePaymentMethod(body *model.PaymentMethod, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(body).Updates(body)
	if res.Error != nil {
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
