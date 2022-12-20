package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/google/uuid"
)

// offline class resource
type OfflineClassResource struct {
	ID                       uint      `json:"id"`
	Title                    string    `json:"title"`
	Time                     time.Time `json:"time"`
	Price                    uint      `json:"price"`
	Duration                 uint      `json:"duration"`
	Slot                     uint      `json:"slot"`
	SlotBooked               uint      `json:"slot_booked"`
	Picture                  string    `json:"picture"`
	OfflineClassCategoryID   uint      `json:"offline_class_category_id"`
	OfflineClassCategoryName string    `json:"offline_class_category_name"`
	TrainerID                uint      `json:"trainer_id"`
	TrainerName              string    `json:"trainer_name"`
}

func (u *OfflineClassResource) FromModel(m *model.OfflineClass) {
	u.ID = m.ID
	u.Title = m.Title
	u.Time = m.Time
	u.Price = m.Price
	u.Duration = m.Duration
	u.Slot = m.Slot
	u.SlotBooked = m.SlotBooked
	u.Picture = m.Picture
	u.OfflineClassCategoryID = m.OfflineClassCategoryID
	u.OfflineClassCategoryName = m.OfflineClassCategory.Name
	u.TrainerID = m.TrainerID
	u.TrainerName = m.Trainer.Name
}

type OfflineClassResources []OfflineClassResource

func (u *OfflineClassResources) FromModel(m []model.OfflineClass) {
	for _, each := range m {
		var resource OfflineClassResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

type OfflineClassDetailResource struct {
	ID                   uint                         `json:"id"`
	Title                string                       `json:"title"`
	Time                 time.Time                    `json:"time"`
	Price                uint                         `json:"price"`
	Duration             uint                         `json:"duration"`
	Slot                 uint                         `json:"slot"`
	SlotBooked           uint                         `json:"slot_booked"`
	Picture              string                       `json:"picture"`
	TrainerID            uint                         `json:"trainer_id"`
	Location             string                       `json:"location"`
	Description          string                       `json:"description"`
	AccessClass          bool                         `json:"access_class"`
	OfflineClassCategory OfflineClassCategoryResource `json:"offline_class_category"`
	Trainer              TrainerResource              `json:"trainer"`
}

func (u *OfflineClassDetailResource) FromModel(m *model.OfflineClass) {
	category := OfflineClassCategoryResource{}
	category.FromModel(&m.OfflineClassCategory)
	trainer := TrainerResource{}
	trainer.FromModel(&m.Trainer)

	u.ID = m.ID
	u.Title = m.Title
	u.Time = m.Time
	u.Price = m.Price
	u.Duration = m.Duration
	u.Slot = m.Slot
	u.SlotBooked = m.SlotBooked
	u.Picture = m.Picture
	u.TrainerID = m.TrainerID
	u.Location = m.Location
	u.Description = m.Description
	u.OfflineClassCategory = category
	u.Trainer = trainer
}

type OfflineClassResponses struct {
	OfflineClasses OfflineClassResources `json:"offline_classes"`
	Count          uint                  `json:"count"`
}

// offline class booking
type OfflineClassBookingResource struct {
	ID                uint             `json:"id"`
	UserName          string           `json:"user_name"`
	UserEmail         string           `json:"user_email"`
	OfflineClassTitle string           `json:"offline_class_title"`
	ExpiredAt         time.Time        `json:"expired_at"`
	ActivedAt         time.Time        `json:"actived_at"`
	Status            model.StatusType `json:"status"`
}

func (u *OfflineClassBookingResource) FromModel(m *model.OfflineClassBooking) {
	u.ID = m.ID
	u.UserName = m.User.Name
	u.UserEmail = m.User.Email
	u.OfflineClassTitle = m.OfflineClass.Title
	u.ExpiredAt = m.ExpiredAt
	u.ActivedAt = m.ActivedAt
	u.Status = m.Status
}

type OfflineClassBookingResources []OfflineClassBookingResource

func (u *OfflineClassBookingResources) FromModel(m []model.OfflineClassBooking) {
	for _, each := range m {
		var resource OfflineClassBookingResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

type OfflineClassBookingResponses struct {
	OfflineClassBookings OfflineClassBookingResources `json:"offline_class_bookings"`
	Page                 uint                         `json:"page"`
	Limit                uint                         `json:"limit"`
	Count                uint                         `json:"count"`
}

type OfflineClassBookingDetailResource struct {
	ID            uint                  `json:"id"`
	User          UserResource          `json:"user"`
	OfflineClass  OfflineClassResource  `json:"offline_class"`
	ExpiredAt     time.Time             `json:"expired_at"`
	ActivedAt     time.Time             `json:"actived_at"`
	ProofPayment  string                `json:"proof_payment"`
	PaymentMethod PaymentMethodResource `json:"payment_method"`
	Code          uuid.UUID             `json:"code"`
	Total         uint                  `json:"total"`
	Status        model.StatusType      `json:"status"`
}

func (u *OfflineClassBookingDetailResource) FromModel(m *model.OfflineClassBooking) {
	offlineClass := OfflineClassResource{}
	offlineClass.FromModel(&m.OfflineClass)
	paymentMethod := PaymentMethodResource{}
	paymentMethod.FromModel(&m.PaymentMethod)
	user := UserResource{}
	user.FromModel(&m.User)

	u.ID = m.ID
	u.User = user
	u.OfflineClass = offlineClass
	u.ExpiredAt = m.ExpiredAt
	u.ActivedAt = m.ActivedAt
	u.ProofPayment = m.ProofPayment
	u.PaymentMethod = paymentMethod
	u.Total = m.Total
	u.Code = m.Code
	u.Status = m.Status
}

type PaymentMethodResource struct {
	ID            *uint  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PaymentNumber string `json:"payment_number"`
}

func (u *PaymentMethodResource) FromModel(m *model.PaymentMethod) {
	u.ID = m.ID
	u.Name = m.Name
	u.Description = m.Description
	u.PaymentNumber = m.PaymentNumber
}

type UserResource struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *UserResource) FromModel(m *model.User) {
	u.ID = m.ID
	u.Name = m.Name
	u.Email = m.Email
}

// offline class category resource
type OfflineClassCategoryResource struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Picture           string `json:"picture"`
	OfflineClassCount uint   `json:"offline_class_count"`
}

func (u *OfflineClassCategoryResource) FromModel(m *model.OfflineClassCategory) {
	count := len(m.OfflineClass)

	u.ID = m.ID
	u.Name = m.Name
	u.Picture = m.Picture
	u.Description = m.Description
	u.OfflineClassCount = uint(count)
}

type OfflineClassCategoryResources []OfflineClassCategoryResource

func (u *OfflineClassCategoryResources) FromModel(m []model.OfflineClassCategory) {
	for _, each := range m {
		var resource OfflineClassCategoryResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

// offline class category resource
type OfflineClassByCategoryResource struct {
	ID                uint                  `json:"id"`
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	OfflineClassCount uint                  `json:"offline_class_count"`
	Picture           string                `json:"picture"`
	OfflineClasses    OfflineClassResources `json:"offline_classes"`
}

func (u *OfflineClassByCategoryResource) FromModel(m *model.OfflineClassCategory) {
	count := len(m.OfflineClass)
	offlineClasses := OfflineClassResources{}
	offlineClasses.FromModel(m.OfflineClass)

	u.ID = m.ID
	u.Name = m.Name
	u.Picture = m.Picture
	u.Description = m.Description
	u.OfflineClassCount = uint(count)
	u.OfflineClasses = offlineClasses
}

// trainer resource
type TrainerResource struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Age     uint   `json:"age"`
	Gender  string `json:"gender"`
	Picture string `json:"picture"`
}

func (u *TrainerResource) FromModel(m *model.Trainer) {
	age := uint(time.Now().Year() - m.Dob.Year())
	if time.Now().Month() < m.Dob.Month() {
		age--
	} else if time.Now().Month() == m.Dob.Month() && time.Now().Day() < m.Dob.Day() {
		age--
	}

	u.ID = m.ID
	u.Name = m.Name
	u.Email = m.Email
	u.Phone = m.Phone
	u.Age = age
	u.Gender = m.Gender
	u.Picture = m.Picture
}
