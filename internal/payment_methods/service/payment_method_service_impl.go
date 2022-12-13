package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/payment_methods/dto"
	paymentMethodRepo "github.com/Group10CapstoneProject/Golang/internal/payment_methods/repository"
	"github.com/Group10CapstoneProject/Golang/model"
)

type paymentMehtodServiceImpl struct {
	paymentMethodRepository paymentMethodRepo.PaymentMethodRepository
}

// CreatePaymentMethod implements PaymentMethodService
func (s *paymentMehtodServiceImpl) CreatePaymentMethod(request *dto.PaymentMethodStoreRequest, ctx context.Context) error {
	paymentMethod := request.ToModel()
	err := s.paymentMethodRepository.CreatePaymentMethod(paymentMethod, ctx)
	return err
}

// DeletePaymentMethod implements PaymentMethodService
func (s *paymentMehtodServiceImpl) DeletePaymentMethod(id uint, ctx context.Context) error {
	paymentMethod := model.PaymentMethod{
		ID: id,
	}
	err := s.paymentMethodRepository.DeletePaymentMethod(&paymentMethod, ctx)
	return err
}

// FindPaymentMethodById implements PaymentMethodService
func (s *paymentMehtodServiceImpl) FindPaymentMethodById(id uint, ctx context.Context) (*dto.PaymentMethodResource, error) {
	paymentMethod, err := s.paymentMethodRepository.FindPaymentMethodById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.PaymentMethodResource
	result.FromModel(paymentMethod)
	return &result, nil
}

// FindPaymentMethods implements PaymentMethodService
func (s *paymentMehtodServiceImpl) FindPaymentMethods(ctx context.Context) (*dto.PaymentMethodResources, error) {
	paymentMethods, err := s.paymentMethodRepository.FindPaymentMethods(ctx)
	if err != nil {
		return nil, err
	}
	var result dto.PaymentMethodResources
	result.FromModel(paymentMethods)
	return &result, nil
}

// UpdatePaymentMethod implements PaymentMethodService
func (s *paymentMehtodServiceImpl) UpdatePaymentMethod(request *dto.PaymentMethodUpdateRequest, ctx context.Context) error {
	paymentMethod := request.ToModel()
	err := s.paymentMethodRepository.UpdatePaymentMethod(paymentMethod, ctx)
	return err
}

func NewPaymentMethodService(paymentMethodRepository paymentMethodRepo.PaymentMethodRepository) PaymentMethodService {
	return &paymentMehtodServiceImpl{
		paymentMethodRepository: paymentMethodRepository,
	}
}
