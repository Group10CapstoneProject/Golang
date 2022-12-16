package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/google/uuid"
)

// member store and update request
type MemberStoreRequest struct {
	UserID          uint
	MemberTypeID    uint  `json:"member_type_id" validate:"required,gte=1"`
	Duration        uint  `json:"duration" validate:"required,gte=1"`
	PaymentMethodID *uint `json:"payment_method_id" validate:"required,gte=1"`
	Total           uint  `json:"total" validate:"required,gte=1"`
}

func (u *MemberStoreRequest) ToModel() *model.Member {
	return &model.Member{
		UserID:          u.UserID,
		MemberTypeID:    u.MemberTypeID,
		Duration:        u.Duration,
		PaymentMethodID: u.PaymentMethodID,
		Total:           u.Total,
	}
}

type MemberUpdateRequest struct {
	ID              uint  `json:"id,omitempty"`
	MemberTypeID    uint  `json:"member_type_id,omitempty" validate:"omitempty,gte=1"`
	Duration        uint  `json:"duration,omitempty" validate:"omitempty,gte=1"`
	PaymentMethodID *uint `json:"payment_method_id,omitempty" validate:"omitempty,gte=1"`
	Total           uint  `json:"total" validate:"omitempty,required,gte=1"`
}

func (u *MemberUpdateRequest) ToModel() *model.Member {
	return &model.Member{
		ID:              u.ID,
		MemberTypeID:    u.MemberTypeID,
		Duration:        u.Duration,
		PaymentMethodID: u.PaymentMethodID,
		Total:           u.Total,
	}
}

// member resource
type MemberResource struct {
	ID             uint             `json:"id"`
	UserName       string           `json:"user_name"`
	UserEmail      string           `json:"user_email"`
	MemberTypeName string           `json:"member_type_name"`
	ExpiredAt      time.Time        `json:"expired_at"`
	ActivedAt      time.Time        `json:"actived_at"`
	Duration       uint             `json:"duration"`
	Status         model.StatusType `json:"status"`
}

func (u *MemberResource) FromModel(m *model.Member) {
	u.ID = m.ID
	u.UserName = m.User.Name
	u.UserEmail = m.User.Email
	u.MemberTypeName = m.MemberType.Name
	u.ExpiredAt = m.ExpiredAt
	u.ActivedAt = m.ActivedAt
	u.Duration = m.Duration
	u.Status = m.Status
}

type MemberResources []MemberResource

func (u *MemberResources) FromModel(m []model.Member) {
	for _, each := range m {
		var resource MemberResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

type MemberResponses struct {
	Members MemberResources `json:"members"`
	Page    uint            `json:"page"`
	Limit   uint            `json:"limit"`
	Count   uint            `json:"count"`
}

// member detail resource
type MemberDetailResource struct {
	ID            uint                  `json:"id"`
	User          UserResource          `json:"user"`
	MemberType    MemberTypeResource    `json:"member_type"`
	ExpiredAt     time.Time             `json:"expired_at"`
	ActivedAt     time.Time             `json:"actived_at"`
	Duration      uint                  `json:"duration"`
	ProofPayment  string                `json:"proof_payment"`
	PaymentMethod PaymentMethodResource `json:"payment_method"`
	Total         uint                  `json:"total"`
	Code          uuid.UUID             `json:"code"`
	Status        model.StatusType      `json:"status"`
}

func (u *MemberDetailResource) FromModel(m *model.Member) {
	memberType := MemberTypeResource{}
	memberType.FromModel(&m.MemberType)
	paymentMethod := PaymentMethodResource{}
	paymentMethod.FromModel(&m.PaymentMethod)
	user := UserResource{}
	user.FromModel(&m.User)

	u.ID = m.ID
	u.User = user
	u.MemberType = memberType
	u.ExpiredAt = m.ExpiredAt
	u.ActivedAt = m.ActivedAt
	u.Duration = m.Duration
	u.ProofPayment = m.ProofPayment
	u.PaymentMethod = paymentMethod
	u.Total = m.Total
	u.Code = m.Code
	u.Status = m.Status
}

type PaymentMethodResource struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PaymentNumber string `json:"payment_number"`
}

func (u *PaymentMethodResource) FromModel(m *model.PaymentMethod) {
	u.ID = *m.ID
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

type SetStatusMember struct {
	ID     uint
	Status model.StatusType `json:"status" validate:"required,status"`
}

func (u *SetStatusMember) ToModel() *model.Member {
	return &model.Member{
		ID:     u.ID,
		Status: u.Status,
	}
}
