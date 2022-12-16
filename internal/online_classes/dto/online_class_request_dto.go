package dto

import (
	"github.com/Group10CapstoneProject/Golang/model"
)

// online class booking store and update request
type OnlineClassBookingStoreRequest struct {
	UserID          uint
	OnlineCLassID   uint  `json:"online_class_id" validate:"required,gte=1"`
	Duration        uint  `json:"duration" validate:"required,gte=1"`
	PaymentMethodID *uint `json:"payment_method_id" validate:"required,gte=0"`
	Total           uint  `json:"total" validate:"required,gte=1"`
}

func (u *OnlineClassBookingStoreRequest) ToModel() *model.OnlineClassBooking {
	return &model.OnlineClassBooking{
		UserID:          u.UserID,
		OnlineClassID:   u.OnlineCLassID,
		Duration:        u.Duration,
		PaymentMethodID: u.PaymentMethodID,
		Total:           u.Total,
	}
}

type OnlineClassBookingUpdateRequest struct {
	ID              uint   `json:"id,omitempty"`
	UserID          uint   `json:"user_id,omitempty"`
	OnlineCLassID   uint   `json:"onlinc_class_id,,omitempty" validate:"omitempty,gte=1"`
	Duration        uint   `json:"duration,,omitempty" validate:"omitemptys,gte=1"`
	PaymentMethodID *uint  `json:"payment_method_id,omitempty" validate:"omitempty,gte=0"`
	ProofPayment    string `json:"proof_payment,omitempty" validate:"omitempty,url"`
	Total           uint   `json:"total,omitempty" validate:"omitempty,gte=1"`
}

func (u *OnlineClassBookingUpdateRequest) ToModel() *model.OnlineClassBooking {
	return &model.OnlineClassBooking{
		ID:              u.ID,
		UserID:          u.UserID,
		OnlineClassID:   u.OnlineCLassID,
		Duration:        u.Duration,
		PaymentMethodID: u.PaymentMethodID,
		ProofPayment:    u.ProofPayment,
		Total:           u.Total,
	}
}

// online class store and update request
type OnlineClassStoreRequest struct {
	Title                 string `json:"title" validate:"required,name"`
	Link                  string `json:"link" validate:"required,url"`
	Price                 uint   `json:"price" validate:"required,gte=1"`
	Description           string `json:"description,omitempty"`
	OnlineClassCategoryID uint   `json:"online_class_category_id" validate:"required,gte=1"`
	Tools                 string `json:"tools,omitempty"`
	TargetArea            string `json:"target_area,omitempty"`
	Duration              uint   `json:"duration" validate:"required,gte=1"`
	Level                 string `json:"level,omitempty"`
	Picture               string `json:"picture" validate:"required,url"`
}

func (u *OnlineClassStoreRequest) ToModel() *model.OnlineClass {
	return &model.OnlineClass{
		Title:                 u.Title,
		Link:                  u.Link,
		Price:                 u.Price,
		Description:           u.Description,
		OnlineClassCategoryID: u.OnlineClassCategoryID,
		Tools:                 u.Tools,
		TargetArea:            u.TargetArea,
		Duration:              u.Duration,
		Level:                 u.Level,
		Picture:               u.Picture,
	}
}

type OnlineClassUpdateRequest struct {
	ID                    uint
	Title                 string `json:"title" validate:"omitempty,name"`
	Link                  string `json:"link" validate:"omitempty,url"`
	Price                 uint   `json:"price" validate:"omitempty,gte=1"`
	Description           string `json:"description,omitempty"`
	OnlineClassCategoryID uint   `json:"online_class_category_id" validate:"omitempty,gte=1"`
	Tools                 string `json:"tools,omitempty"`
	TargetArea            string `json:"target_area,omitempty"`
	Duration              uint   `json:"duration" validate:"omitempty,gte=1"`
	Level                 string `json:"level,omitempty"`
	Picture               string `json:"picture" validate:"omitempty,url"`
}

func (u *OnlineClassUpdateRequest) ToModel() *model.OnlineClass {
	return &model.OnlineClass{
		ID:                    u.ID,
		Title:                 u.Title,
		Link:                  u.Link,
		Price:                 u.Price,
		Description:           u.Description,
		OnlineClassCategoryID: u.OnlineClassCategoryID,
		Tools:                 u.Tools,
		TargetArea:            u.TargetArea,
		Duration:              u.Duration,
		Level:                 u.Level,
		Picture:               u.Picture,
	}
}

// online class category store and update request
type OnlineClassCategoryStoreRequest struct {
	Name        string `json:"name" validate:"required,name"`
	Description string `json:"description,omitempty"`
	Picture     string `json:"picture" validate:"required,url"`
}

func (u *OnlineClassCategoryStoreRequest) ToModel() *model.OnlineClassCategory {
	return &model.OnlineClassCategory{
		Name:        u.Name,
		Description: u.Description,
		Picture:     u.Picture,
	}
}

type OnlineClassCategoryUpdateRequest struct {
	ID          uint
	Name        string `json:"name,omitempty" validate:"omitempty,name"`
	Description string `json:"description,omitempty"`
	Picture     string `json:"picture,omitempty" validate:"omitempty,url"`
}

func (u *OnlineClassCategoryUpdateRequest) ToModel() *model.OnlineClassCategory {
	return &model.OnlineClassCategory{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
		Picture:     u.Picture,
	}
}

// set status booking
type SetStatusOnlineClassBooking struct {
	ID     uint
	Status model.StatusType `json:"status" validate:"required,status"`
}

func (u *SetStatusOnlineClassBooking) ToModel() *model.OnlineClassBooking {
	return &model.OnlineClassBooking{
		ID:     u.ID,
		Status: u.Status,
	}
}
