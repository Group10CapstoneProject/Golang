package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type PaymentMethodRepository interface {
	CreatePaymentMethod(body *model.PaymentMethod, ctx context.Context) error
	FindPaymentMethods(access bool, ctx context.Context) ([]model.PaymentMethod, error)
	FindPaymentMethodById(id uint, ctx context.Context) (*model.PaymentMethod, error)
	CheckPaymentMethodIsDeleted(body *model.PaymentMethod) error
	UpdatePaymentMethod(body *model.PaymentMethod, ctx context.Context) error
	DeletePaymentMethod(body *model.PaymentMethod, ctx context.Context) error
}
