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
	PaymentMethodID *uint  `json:"payment_method_id" validate:"required,gte=0"`
	Total           uint   `json:"total" validate:"required,gte=1"`
}

func (u *TrainerBookingStoreRequest) ToModel() *model.TrainerBooking {
	zone, _ := time.Now().Zone()
	time, err := time.Parse(constans.FormatTime, u.Time+" "+zone)
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
	ID              uint
	UserID          uint   `json:"user_id,omitempty" validate:"omitempty,gte=1"`
	TrainerID       uint   `json:"trainer_id,omitempty" validate:"omitempty,gte=1"`
	Time            string `json:"time,omitempty" validate:"omitempty,mytime"`
	PaymentMethodID *uint  `json:"payment_method_id,omitempty" validate:"omitempty,gte=1"`
	ProofPayment    string `json:"proof_payment,omitempty" validate:"omitempty,url"`
	Total           uint   `json:"total,omitempty" validate:"omitempty,gte=1"`
}

func (u *TrainerBookingUpdateRequest) ToModel() *model.TrainerBooking {
	zone, _ := time.Now().Zone()
	time, err := time.Parse(constans.FormatTime, u.Time+" "+zone)
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
	Name         string `json:"name" validate:"required,personname"`
	Email        string `json:"email" validate:"required,email"`
	Phone        string `json:"phone" validate:"required,numeric,min=11,max=13"`
	Dob          string `json:"dob" validate:"required,dob"`
	Gender       string `json:"gender" validate:"required,gender"`
	Price        uint   `json:"price" validate:"required,gte=1"`
	DailySlot    uint   `json:"daily_slot" validate:"required,gte=1"`
	Description  string `json:"description,omitempty"`
	TrainerSkill []uint `json:"skills" validate:"required"`
	Picture      string `json:"picture" validate:"required,url"`
}

func (u *TrainerStoreRequest) ToModel() *model.Trainer {
	dob, err := time.Parse(constans.FormatDate, u.Dob)
	if err != nil {
		return nil
	}
	var trainerSkill []model.TrainerSkill
	for _, id := range u.TrainerSkill {
		a := &model.TrainerSkill{
			SkillID: id,
		}
		trainerSkill = append(trainerSkill, *a)
	}

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
	Name         string `json:"name,omitempty" validate:"omitempty,personname"`
	Email        string `json:"email,omitempty" validate:"omitempty,email"`
	Phone        string `json:"phone,omitempty" validate:"omitempty,numeric,min=11,max=13"`
	Dob          string `json:"dob,omitempty" validate:"omitempty,dob"`
	Gender       string `json:"gender,omitempty" validate:"omitempty,gender"`
	Price        uint   `json:"price,omitempty" validate:"omitempty,gte=1"`
	DailySlot    uint   `json:"daily_slot,omitempty" validate:"omitempty,gte=1"`
	Description  string `json:"description,omitempty"`
	TrainerSkill []uint `json:"skills,omitempty" validate:"omitempty"`
	Picture      string `json:"picture,omitempty" validate:"omitempty,url"`
}

func (u *TrainerUpdateRequest) ToModel() *model.Trainer {
	dob, err := time.Parse(constans.FormatDate, u.Dob)
	if err != nil {
		return nil
	}
	var trainerSkill []model.TrainerSkill
	for _, id := range u.TrainerSkill {
		a := &model.TrainerSkill{
			SkillID: id,
		}
		trainerSkill = append(trainerSkill, *a)
	}

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
	Gender     string `json:"gender,omitempty" validate:"omitempty,gender"`
	PriceOrder string `json:"order_by_price,omitempty" validate:"omitempty,ordertype"`
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
