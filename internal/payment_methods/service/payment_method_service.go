package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/payment_methods/dto"
)

type PaymentMethodService interface {
	CreatePaymentMethod(request *dto.PaymentMethodStoreRequest, ctx context.Context) error
	FindPaymentMethods(access bool, ctx context.Context) (*dto.PaymentMethodResources, error)
	FindPaymentMethodById(id uint, ctx context.Context) (*dto.PaymentMethodResource, error)
	UpdatePaymentMethod(request *dto.PaymentMethodUpdateRequest, ctx context.Context) error
	DeletePaymentMethod(id uint, ctx context.Context) error
}
