package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/google/uuid"
)

// member store and update request
type MemberStoreRequest struct {
	UserID          uint
	MemberTypeID    uint `json:"member_type_id" validate:"required,gte=1"`
	Duration        uint `json:"duration" validate:"required,gte=1"`
	PaymentMethodId uint `json:"payment_method_id" validate:"required,gte=1"`
	Total           uint `json:"total" validate:"required,gte=1"`
}

func (u *MemberStoreRequest) ToModel() *model.Member {
	return &model.Member{
		UserID:          u.UserID,
		MemberTypeID:    u.MemberTypeID,
		Duration:        u.Duration,
		PaymentMethodId: u.PaymentMethodId,
		Total:           u.Total,
	}
}

type MemberUpdateRequest struct {
	ID              uint   `json:"id,omitempty"`
	MemberTypeID    uint   `json:"member_type_id,omitempty" validate:"omitempty,gte=1"`
	Duration        uint   `json:"duration,omitempty" validate:"omitempty,gte=1"`
	PaymentMethodId uint   `json:"payment_method_id,omitempty" validate:"omitempty,gte=1"`
	ProofPayment    string `json:"proof_payment,omitempty" validate:"omitempty,url"`
	Total           uint   `json:"total" validate:"omitempty,required,gte=1"`
}

func (u *MemberUpdateRequest) ToModel() *model.Member {
	return &model.Member{
		ID:              u.ID,
		MemberTypeID:    u.MemberTypeID,
		Duration:        u.Duration,
		PaymentMethodId: u.PaymentMethodId,
		ProofPayment:    u.ProofPayment,
		Total:           u.Total,
	}
}

// member resource
type MemberResource struct {
	ID             uint             `json:"id"`
	UserName       string           `json:"user_name"`
	MemberTypeName string           `json:"member_type_name"`
	ExpiredAt      time.Time        `json:"expired_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	Duration       uint             `json:"duration"`
	Status         model.StatusType `json:"status"`
}

func (u *MemberResource) FromModel(m *model.Member) {
	u.ID = m.ID
	u.UserName = m.User.Name
	u.MemberTypeName = m.MemberType.Name
	u.ExpiredAt = m.ExpiredAt
	u.UpdatedAt = m.UpdatedAt
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
	Count   uint            `json:"count"`
}

// member detail resource
type MemberDetailResource struct {
	ID            uint                  `json:"id"`
	UserID        uint                  `json:"user_id"`
	UserName      string                `json:"user_name"`
	UserEmail     string                `json:"user_email"`
	MemberType    MemberTypeResource    `json:"member_type"`
	ExpiredAt     time.Time             `json:"expired_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
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

	u.ID = m.ID
	u.UserID = m.UserID
	u.UserName = m.User.Name
	u.UserEmail = m.User.Email
	u.MemberType = memberType
	u.ExpiredAt = m.ExpiredAt
	u.UpdatedAt = m.UpdatedAt
	u.Duration = m.Duration
	u.ProofPayment = m.ProofPayment
	u.PaymentMethod = paymentMethod
	u.Total = m.Total
	u.Code = m.Code
	u.Status = m.Status
}

type PaymentMethodResource struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	PaymentNumber string `json:"payment_number"`
}

func (u *PaymentMethodResource) FromModel(m *model.PaymentMethod) {
	u.Name = m.Name
	u.Description = m.Description
	u.PaymentNumber = m.PaymentNumber
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
