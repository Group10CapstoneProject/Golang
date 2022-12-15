package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/google/uuid"
)

// trainer booking store and update request
type TrainerBookingStoreRequest struct {
	UserID          uint
	TrainerID       uint   `json:"trainer_id" validate:"required,gte=1"`
	Time            string `json:"time" validate:"required,mytime"`
	PaymentMethodID uint   `json:"payment_method_id" validate:"required,gte=1"`
	Total           uint   `json:"total" validate:"required,gte=1"`
}

func (u *TrainerBookingStoreRequest) ToModel() *model.TrainerBooking {
	time, err := time.Parse(constans.FormatTime, u.Time)
	if err != nil {
		return nil
	}

	return &model.TrainerBooking{
		UserID:          u.UserID,
		TrainerID:       u.TrainerID,
		Time:            time,
		PaymentMethodID: u.PaymentMethodID,
		Total:           u.Total,
	}
}

type TrainerBookingUpdateRequest struct {
	ID              uint   `json:"id,omitempty"`
	UserID          uint   `json:"user_id,omitempty"`
	TrainerID       uint   `json:"trainer_id,omitempty" validate:"omitempty,gte=1"`
	Time            string `json:"time,omitempty" validate:"omitemptys,gte=1"`
	PaymentMethodID uint   `json:"payment_method_id,omitempty" validate:"omitempty,gte=1"`
	ProofPayment    string `json:"proof_payment,omitempty" validate:"omitempty,url"`
	Total           uint   `json:"total,omitempty" validate:"omitempty,gte=1"`
}

func (u *TrainerBookingUpdateRequest) ToModel() *model.TrainerBooking {
	time, err := time.Parse(constans.FormatTime, u.Time)
	if err != nil {
		return nil
	}
	return &model.TrainerBooking{
		ID:              u.ID,
		UserID:          u.UserID,
		TrainerID:       u.TrainerID,
		Time:            time,
		PaymentMethodID: u.PaymentMethodID,
		ProofPayment:    u.ProofPayment,
		Total:           u.Total,
	}
}

// trainer store and update request
type TrainerStoreRequest struct {
	Name         string               `json:"name" validate:"required,personname"`
	Email        string               `json:"email" validate:"required,email"`
	Phone        string               `json:"phone" validate:"required,min=11,max=13"`
	Dob          string               `json:"dob" validate:"required,dob"`
	Gender       string               `json:"gender" validate:"required,alpha,min=1,max=1"`
	Price        uint                 `json:"price" validate:"required,gte=1"`
	DailySlot    uint                 `json:"daily_slot" validate:"required,gte=1"`
	Description  string               `json:"description,omitempty"`
	TrainerSkill TrainerSkillRequests `json:"trainer_skill" validate:"required"`
	Picture      string               `json:"picture" validate:"required,url"`
}

func (u *TrainerStoreRequest) ToModel() *model.Trainer {
	dob, err := time.Parse(constans.FormatDate, u.Dob)
	if err != nil {
		return nil
	}
	trainerSkill := u.TrainerSkill.ToModel()

	return &model.Trainer{
		Name:         u.Name,
		Email:        u.Email,
		Phone:        u.Phone,
		Dob:          dob,
		Gender:       u.Gender,
		Price:        u.Price,
		DailySlot:    u.DailySlot,
		Description:  u.Description,
		TrainerSkill: trainerSkill,
		Picture:      u.Picture,
	}
}

type TrainerUpdateRequest struct {
	ID           uint
	Name         string               `json:"name,omitempty" validate:"omitempty,personname"`
	Email        string               `json:"email,omitempty" validate:"omitempty,email"`
	Phone        string               `json:"phone,omitempty" validate:"omitempty,min=11,max=13"`
	Dob          string               `json:"dob,omitempty" validate:"omitempty,dob"`
	Gender       string               `json:"gender,omitempty" validate:"omitempty,alpha,min=1,max=1"`
	Price        uint                 `json:"price,omitempty" validate:"omitempty,gte=1"`
	DailySlot    uint                 `json:"daily_slot,omitempty" validate:"omitempty,gte=1"`
	Description  string               `json:"description,omitempty"`
	TrainerSkill TrainerSkillRequests `json:"trainer_skill,omitempty" validate:"omitempty"`
	Picture      string               `json:"picture,omitempty" validate:"omitempty,url"`
}

func (u *TrainerUpdateRequest) ToModel() *model.Trainer {
	dob, err := time.Parse(constans.FormatDate, u.Dob)
	if err != nil {
		return nil
	}
	trainerSkill := u.TrainerSkill.ToModel()

	return &model.Trainer{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Phone:        u.Phone,
		Dob:          dob,
		Gender:       u.Gender,
		Price:        u.Price,
		DailySlot:    u.DailySlot,
		Description:  u.Description,
		TrainerSkill: trainerSkill,
		Picture:      u.Picture,
	}
}

// trainer skill store and update request
type TrainerSkillRequest struct {
	SkillID uint `json:"skill_id" validate:"required,gte=1"`
}

func (u *TrainerSkillRequest) ToModel() *model.TrainerSkill {
	return &model.TrainerSkill{
		SkillID: u.SkillID,
	}
}

type TrainerSkillRequests []TrainerSkillRequest

func (u *TrainerSkillRequests) ToModel() []model.TrainerSkill {
	var trainerSkill []model.TrainerSkill
	for _, skill := range *u {
		a := skill.ToModel()
		trainerSkill = append(trainerSkill, *a)
	}
	return trainerSkill
}

// skill store and update request
type SkillStoreRequest struct {
	Name        string `json:"name" validate:"required,name"`
	Description string `json:"description,omitempty"`
}

func (u *SkillStoreRequest) ToModel() *model.Skill {
	return &model.Skill{
		Name:        u.Name,
		Description: u.Description,
	}
}

type SkillUpdateRequest struct {
	ID          uint
	Name        string `json:"name,omitempty" validate:"omitempty,name"`
	Description string `json:"description,omitempty"`
}

func (u *SkillUpdateRequest) ToModel() *model.Skill {
	return &model.Skill{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

// set status booking
type SetStatusTrainerBooking struct {
	ID     uint
	Status model.StatusType `json:"status" validate:"required,status"`
}

func (u *SetStatusTrainerBooking) ToModel() *model.TrainerBooking {
	return &model.TrainerBooking{
		ID:     u.ID,
		Status: u.Status,
	}
}

// filter trainer
type FilterTrainer struct {
	Name       string `json:"name,omitempty" validate:"omitempty,personname"`
	Date       string `json:"date,omitempty" validate:"omitempty,mydate"`
	Gender     string `json:"gender,omitempty" validate:"omitempty,min=1,max=1"`
	PriceOrder string `json:"price,omitempty" validate:"omitempty,ordertype"`
}

// take booking
type TakeTrainerBooking struct {
	Email string    `json:"email" validate:"required,email"`
	Code  uuid.UUID `json:"code" validate:"required"`
}

func (u *TakeTrainerBooking) ToModel() *model.TrainerBooking {
	return &model.TrainerBooking{
		User: model.User{
			Email: u.Email,
		},
		Code: u.Code,
	}
}
