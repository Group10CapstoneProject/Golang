package dto

import "github.com/Group10CapstoneProject/Golang/model"

// payment method request
type PaymentMethodStoreRequest struct {
	Name          string `json:"name" validate:"required,name"`
	PaymentNumber string `json:"payment_number" validate:"required,number"`
	Picture       string `json:"picture" validate:"required,url"`
	Description   string `json:"description,omitempty"`
}

func (u *PaymentMethodStoreRequest) ToModel() *model.PaymentMethod {
	return &model.PaymentMethod{
		Name:          u.Name,
		Picture:       u.Picture,
		PaymentNumber: u.PaymentNumber,
		Description:   u.Description,
	}
}

type PaymentMethodUpdateRequest struct {
	ID            uint
	Name          string `json:"name" validate:"omitempty,name"`
	PaymentNumber string `json:"payment_number" validate:"omitempty,number"`
	Picture       string `json:"picture" validate:"omitempty,url"`
	Description   string `json:"description,omitempty"`
}

func (u *PaymentMethodUpdateRequest) ToModel() *model.PaymentMethod {
	return &model.PaymentMethod{
		ID:            &u.ID,
		Name:          u.Name,
		Picture:       u.Picture,
		PaymentNumber: u.PaymentNumber,
		Description:   u.Description,
	}
}

// payment method resource
type PaymentMethodResource struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	PaymentNumber string `json:"payment_number"`
	Picture       string `json:"picture"`
	Description   string `json:"description"`
}

func (u *PaymentMethodResource) FromModel(m *model.PaymentMethod) {
	u.ID = *m.ID
	u.Name = m.Name
	u.Picture = m.Picture
	u.PaymentNumber = m.PaymentNumber
	u.Description = m.Description
}

type PaymentMethodResources []PaymentMethodResource

func (u *PaymentMethodResources) FromModel(m []model.PaymentMethod) {
	for _, each := range m {
		var resource PaymentMethodResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}
