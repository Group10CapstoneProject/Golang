package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/google/uuid"
)

// offline class booking store and update request
type OfflineClassBookingStoreRequest struct {
	UserID          uint
	OfflineCLassID  uint  `json:"offline_class_id" validate:"required,gte=1"`
	PaymentMethodID *uint `json:"payment_method_id" validate:"required,gte=0"`
	Total           uint  `json:"total" validate:"required,gte=1"`
}

func (u *OfflineClassBookingStoreRequest) ToModel() *model.OfflineClassBooking {
	return &model.OfflineClassBooking{
		UserID:          u.UserID,
		OfflineClassID:  u.OfflineCLassID,
		PaymentMethodID: u.PaymentMethodID,
		Total:           u.Total,
	}
}

type OfflineClassBookingUpdateRequest struct {
	ID              uint
	UserID          uint   `json:"user_id,omitempty"`
	OfflineCLassID  uint   `json:"offline_class_id,,omitempty" validate:"omitempty,gte=1"`
	PaymentMethodID *uint  `json:"payment_method_id,omitempty" validate:"omitempty,gte=0"`
	ProofPayment    string `json:"proof_payment,omitempty" validate:"omitempty,url"`
	Total           uint   `json:"total,omitempty" validate:"omitempty,gte=1"`
}

func (u *OfflineClassBookingUpdateRequest) ToModel() *model.OfflineClassBooking {
	return &model.OfflineClassBooking{
		ID:              u.ID,
		UserID:          u.UserID,
		OfflineClassID:  u.OfflineCLassID,
		PaymentMethodID: u.PaymentMethodID,
		ProofPayment:    u.ProofPayment,
		Total:           u.Total,
	}
}

// offline class store and update request
type OfflineClassStoreRequest struct {
	Title                  string `json:"title" validate:"required,name"`
	Time                   string `json:"time" validate:"required,mytime"`
	Duration               uint   `json:"duration" validate:"required,gte=1"`
	Slot                   uint   `json:"slot" validate:"required,gte=1"`
	Price                  uint   `json:"price" validate:"required,gte=1"`
	Picture                string `json:"picture" validate:"required,url"`
	Description            string `json:"description,omitempty"`
	TrainerID              uint   `json:"trainer_id" validate:"required,gte=1"`
	Location               string `json:"location" validate:"required"`
	OfflineClassCategoryID uint   `json:"offline_class_category_id" validate:"required,gte=1"`
}

func (u *OfflineClassStoreRequest) ToModel() *model.OfflineClass {
	zone, _ := time.Now().Zone()
	time, err := time.Parse(constans.FormatTime, u.Time+" "+zone)
	if err != nil {
		return nil
	}

	return &model.OfflineClass{
		Title:                  u.Title,
		Time:                   time,
		Duration:               u.Duration,
		Slot:                   u.Slot,
		Price:                  u.Price,
		Picture:                u.Picture,
		Description:            u.Description,
		TrainerID:              u.TrainerID,
		Location:               u.Location,
		OfflineClassCategoryID: u.OfflineClassCategoryID,
	}
}

type OfflineClassUpdateRequest struct {
	ID                     uint
	Title                  string `json:"title" validate:"omitempty,name"`
	Time                   string `json:"time" validate:"omitempty,mytime"`
	Duration               uint   `json:"duration" validate:"omitempty,gte=1"`
	Slot                   uint   `json:"slot" validate:"omitempty,gte=1"`
	Price                  uint   `json:"price" validate:"omitempty,gte=1"`
	Picture                string `json:"picture" validate:"omitempty,url"`
	Description            string `json:"description,omitempty"`
	Location               string `json:"location" validate:"omitempty"`
	TrainerID              uint   `json:"trainer_id" validate:"omitempty,gte=1"`
	OfflineClassCategoryID uint   `json:"offline_class_category_id" validate:"omitempty,gte=1"`
}

func (u *OfflineClassUpdateRequest) ToModel() *model.OfflineClass {
	zone, _ := time.Now().Zone()
	time, err := time.Parse(constans.FormatTime, u.Time+" "+zone)
	if err != nil {
		return nil
	}

	return &model.OfflineClass{
		ID:                     u.ID,
		Title:                  u.Title,
		Time:                   time,
		Duration:               u.Duration,
		Slot:                   u.Slot,
		Price:                  u.Price,
		Picture:                u.Picture,
		Description:            u.Description,
		Location:               u.Location,
		TrainerID:              u.TrainerID,
		OfflineClassCategoryID: u.OfflineClassCategoryID,
	}
}

// offline class category store and update request
type OfflineClassCategoryStoreRequest struct {
	Name        string `json:"name" validate:"required,name"`
	Description string `json:"description,omitempty"`
	Picture     string `json:"picture" validate:"required,url"`
}

func (u *OfflineClassCategoryStoreRequest) ToModel() *model.OfflineClassCategory {
	return &model.OfflineClassCategory{
		Name:        u.Name,
		Description: u.Description,
		Picture:     u.Picture,
	}
}

type OfflineClassCategoryUpdateRequest struct {
	ID          uint
	Name        string `json:"name,omitempty" validate:"omitempty,name"`
	Description string `json:"description,omitempty"`
	Picture     string `json:"picture,omitempty" validate:"omitempty,url"`
}

func (u *OfflineClassCategoryUpdateRequest) ToModel() *model.OfflineClassCategory {
	return &model.OfflineClassCategory{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
		Picture:     u.Picture,
	}
}

// set status booking
type SetStatusOfflineClassBooking struct {
	ID     uint
	Status model.StatusType `json:"status" validate:"required,status"`
}

func (u *SetStatusOfflineClassBooking) ToModel() *model.OfflineClassBooking {
	return &model.OfflineClassBooking{
		ID:     u.ID,
		Status: u.Status,
	}
}

// take booking
type TakeOfflineClassBooking struct {
	Email string    `json:"email" validate:"required,email"`
	Code  uuid.UUID `json:"code" validate:"required"`
}

func (u *TakeOfflineClassBooking) ToModel() *model.OfflineClassBooking {
	return &model.OfflineClassBooking{
		User: model.User{
			Email: u.Email,
		},
		Code: u.Code,
	}
}

type OfflineClassFilter struct {
	Time         string `json:"time,omitempty" validate:"omitempty,mydate"`
	CategoryId   uint   `json:"category_id,omitempty" validate:"omitempty,gte=1"`
	OrderByPrice string `json:"order_by_price,omitempty" validate:"omitempty,ordertype"`
	Q            string
}
