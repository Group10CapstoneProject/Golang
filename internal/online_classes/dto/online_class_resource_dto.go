package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/model"
)

// online class resource
type OnlineClassResource struct {
	ID                    uint   `json:"id"`
	Title                 string `json:"title"`
	Price                 uint   `json:"price"`
	Duration              uint   `json:"duration"`
	Level                 string `json:"level"`
	Picture               string `json:"picture"`
	OnlineClassCategoryID uint   `json:"online_class_category_id"`
}

func (u *OnlineClassResource) FromModel(m *model.OnlineClass) {
	u.ID = m.ID
	u.Title = m.Title
	u.Price = m.Price
	u.Duration = m.Duration
	u.Level = m.Level
	u.Picture = m.Picture
	u.OnlineClassCategoryID = m.OnlineClassCategoryID
}

type OnlineClassResources []OnlineClassResource

func (u *OnlineClassResources) FromModel(m []model.OnlineClass) {
	for _, each := range m {
		var resource OnlineClassResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

type OnlineClassDetailResource struct {
	ID                  uint                        `json:"id"`
	Title               string                      `json:"title"`
	Price               uint                        `json:"price"`
	Description         string                      `json:"description"`
	Link                string                      `json:"link"`
	Path                string                      `json:"path"`
	Tools               string                      `json:"tools"`
	TargetArea          string                      `json:"target_area"`
	Duration            uint                        `json:"duration"`
	Level               string                      `json:"level"`
	Picture             string                      `json:"picture"`
	OnlineClassCategory OnlineClassCategoryResource `json:"online_class_category"`
	AccessClass         bool                        `json:"access_class"`
}

func (u *OnlineClassDetailResource) FromModel(m *model.OnlineClass) {
	category := OnlineClassCategoryResource{}
	category.FromModel(&m.OnlineClassCategory)

	u.ID = m.ID
	u.Title = m.Title
	u.Price = m.Price
	u.Description = m.Description
	u.Link = m.Link
	u.Path = m.Path
	u.Tools = m.Tools
	u.TargetArea = m.TargetArea
	u.Duration = m.Duration
	u.Level = m.Level
	u.Picture = m.Picture
	u.OnlineClassCategory = category
}

// online class booking
type OnlineClassBookingResource struct {
	ID               uint             `json:"id"`
	UserName         string           `json:"user_name"`
	UserEmail        string           `json:"user_email"`
	OnlineClassTitle string           `json:"online_class_title"`
	ExpiredAt        time.Time        `json:"expired_at"`
	ActivedAt        time.Time        `json:"actived_at"`
	Duration         uint             `json:"duration"`
	Status           model.StatusType `json:"status"`
}

func (u *OnlineClassBookingResource) FromModel(m *model.OnlineClassBooking) {
	u.ID = m.ID
	u.UserName = m.User.Name
	u.UserEmail = m.User.Email
	u.OnlineClassTitle = m.OnlineClass.Title
	u.Duration = m.Duration
	u.ExpiredAt = m.ExpiredAt
	u.ActivedAt = m.ActivedAt
	u.Status = m.Status
}

type OnlineClassBookingResources []OnlineClassBookingResource

func (u *OnlineClassBookingResources) FromModel(m []model.OnlineClassBooking) {
	for _, each := range m {
		var resource OnlineClassBookingResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

type OnlineClassBookingResponses struct {
	Members OnlineClassBookingResources `json:"online_class_booking"`
	Page    uint                        `json:"page"`
	Limit   uint                        `json:"limit"`
	Count   uint                        `json:"count"`
}

type OnlineClassBookingDetailResource struct {
	ID            uint                      `json:"id"`
	User          UserResource              `json:"user"`
	OnlineClass   OnlineClassDetailResource `json:"online_class"`
	ExpiredAt     time.Time                 `json:"expired_at"`
	ActivedAt     time.Time                 `json:"actived_at"`
	Duration      uint                      `json:"duration"`
	ProofPayment  string                    `json:"proof_payment"`
	PaymentMethod PaymentMethodResource     `json:"payment_method"`
	Total         uint                      `json:"total"`
	Status        model.StatusType          `json:"status"`
}

func (u *OnlineClassBookingDetailResource) FromModel(m *model.OnlineClassBooking) {
	onlineClass := OnlineClassDetailResource{}
	onlineClass.FromModel(&m.OnlineClass)
	paymentMethod := PaymentMethodResource{}
	paymentMethod.FromModel(&m.PaymentMethod)
	user := UserResource{}
	user.FromModel(&m.User)

	u.ID = m.ID
	u.User = user
	u.OnlineClass = onlineClass
	u.ExpiredAt = m.ExpiredAt
	u.ActivedAt = m.ActivedAt
	u.Duration = m.Duration
	u.ProofPayment = m.ProofPayment
	u.PaymentMethod = paymentMethod
	u.Total = m.Total
	u.Status = m.Status
}

type PaymentMethodResource struct {
	ID            uint   `json:"id"`
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

// online class category resource
type OnlineClassCategoryResource struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	OnlineClassCount uint   `json:"online_class_count"`
}

func (u *OnlineClassCategoryResource) FromModel(m *model.OnlineClassCategory) {
	count := len(m.OnlineClass)

	u.ID = m.ID
	u.Name = m.Name
	u.Description = m.Description
	u.OnlineClassCount = uint(count)
}

type OnlineClassCategoryResources []OnlineClassCategoryResource

func (u *OnlineClassCategoryResources) FromModel(m []model.OnlineClassCategory) {
	for _, each := range m {
		var resource OnlineClassCategoryResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

// online class category resource
type OnlineClassByCategoryResource struct {
	ID               uint                 `json:"id"`
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	OnlineClassCount uint                 `json:"online_class_count"`
	OnlineClasses    OnlineClassResources `json:"online_classes"`
}

func (u *OnlineClassByCategoryResource) FromModel(m *model.OnlineClassCategory) {
	count := len(m.OnlineClass)
	onlineClasses := OnlineClassResources{}
	onlineClasses.FromModel(m.OnlineClass)

	u.ID = m.ID
	u.Name = m.Name
	u.Description = m.Description
	u.OnlineClassCount = uint(count)
	u.OnlineClasses = onlineClasses
}
