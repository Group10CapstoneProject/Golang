package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/model"
)

// trainer resource
type TrainerResource struct {
	ID           uint                  `json:"id"`
	Name         string                `json:"name"`
	Phone        string                `json:"phone"`
	Age          uint                  `json:"age"`
	Gender       string                `json:"gender"`
	Price        uint                  `json:"price"`
	DailySlot    uint                  `json:"daily_slot"`
	Description  string                `json:"description"`
	Picture      string                `json:"picture"`
	TrainerSkill TrainerSkillResources `json:"trainer_skill"`
}

func (u *TrainerResource) FromModel(m *model.Trainer) {
	age := uint(time.Now().Year() - m.Dob.Year())
	if time.Now().Month() < m.Dob.Month() {
		age--
	} else if time.Now().Month() == m.Dob.Month() && time.Now().Day() < m.Dob.Day() {
		age--
	}
	trainerSkill := TrainerSkillResources{}
	trainerSkill.FromModel(m.TrainerSkill)

	u.ID = m.ID
	u.Name = m.Name
	u.Phone = m.Phone
	u.Age = age
	u.Gender = m.Gender
	u.Price = m.Price
	u.DailySlot = m.DailySlot
	u.Description = m.Description
	u.Picture = m.Picture
	u.TrainerSkill = trainerSkill
}

type TrainerResources []TrainerResource

func (u *TrainerResources) FromModel(m []model.Trainer) {
	for _, each := range m {
		var resource TrainerResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

type TrainerDetailResource struct {
	ID            uint                  `json:"id"`
	Name          string                `json:"name"`
	Phone         string                `json:"phone"`
	Age           uint                  `json:"age"`
	Dob           time.Time             `json:"dob"`
	Gender        string                `json:"gender"`
	Price         uint                  `json:"price"`
	DailySlot     uint                  `json:"daily_slot"`
	Description   string                `json:"description"`
	Picture       string                `json:"picture"`
	TrainerSkill  TrainerSkillResources `json:"trainer_skill"`
	ClientActive  uint                  `json:"client_active"`
	AccessTrainer bool                  `json:"access_trainer"`
}

func (u *TrainerDetailResource) FromModel(m *model.Trainer) {
	age := uint(time.Now().Year() - m.Dob.Year())
	if time.Now().Month() < m.Dob.Month() {
		age--
	} else if time.Now().Month() == m.Dob.Month() && time.Now().Day() < m.Dob.Day() {
		age--
	}
	count := len(m.TrainerBooking)
	trainerSkill := TrainerSkillResources{}
	trainerSkill.FromModel(m.TrainerSkill)

	u.ID = m.ID
	u.Name = m.Name
	u.Phone = m.Phone
	u.Age = age
	u.Dob = m.Dob
	u.Gender = m.Gender
	u.Price = m.Price
	u.DailySlot = m.DailySlot
	u.Description = m.Description
	u.Picture = m.Picture
	u.TrainerSkill = trainerSkill
	u.ClientActive = uint(count)
}

// trainer skill resource
type TrainerSkillResource struct {
	SkillID   uint   `json:"skill_id"`
	SkillName string `json:"skill_name"`
}

func (u *TrainerSkillResource) FromModel(m *model.TrainerSkill) {
	u.SkillID = m.SkillID
	u.SkillName = m.Skill.Name
}

type TrainerSkillResources []TrainerSkillResource

func (u *TrainerSkillResources) FromModel(m []model.TrainerSkill) {
	for _, each := range m {
		var resource TrainerSkillResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

// online class booking
type TrainerBookingResource struct {
	ID          uint             `json:"id"`
	UserName    string           `json:"user_name"`
	UserEmail   string           `json:"user_email"`
	TrainerName string           `json:"trainer_name"`
	ExpiredAt   time.Time        `json:"expired_at"`
	ActivedAt   time.Time        `json:"actived_at"`
	Time        time.Time        `json:"time"`
	Status      model.StatusType `json:"status"`
}

func (u *TrainerBookingResource) FromModel(m *model.TrainerBooking) {
	u.ID = m.ID
	u.UserName = m.User.Name
	u.UserEmail = m.User.Email
	u.TrainerName = m.Trainer.Name
	u.Time = m.Time
	u.ExpiredAt = m.ExpiredAt
	u.ActivedAt = m.ActivedAt
	u.Status = m.Status
}

type TrainerBookingResources []TrainerBookingResource

func (u *TrainerBookingResources) FromModel(m []model.TrainerBooking) {
	for _, each := range m {
		var resource TrainerBookingResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

type TrainerBookingResponses struct {
	TrainerBookings TrainerBookingResources `json:"trainer_bookings"`
	Page            uint                    `json:"page"`
	Limit           uint                    `json:"limit"`
	Count           uint                    `json:"count"`
}

type TrainerBookingDetailResource struct {
	ID            uint                  `json:"id"`
	User          UserResource          `json:"user"`
	Trainer       TrainerResource       `json:"trainer"`
	ExpiredAt     time.Time             `json:"expired_at"`
	ActivedAt     time.Time             `json:"actived_at"`
	Time          time.Time             `json:"time"`
	ProofPayment  string                `json:"proof_payment"`
	PaymentMethod PaymentMethodResource `json:"payment_method"`
	Total         uint                  `json:"total"`
	Status        model.StatusType      `json:"status"`
}

func (u *TrainerBookingDetailResource) FromModel(m *model.TrainerBooking) {
	trainer := TrainerResource{}
	trainer.FromModel(&m.Trainer)
	paymentMethod := PaymentMethodResource{}
	paymentMethod.FromModel(&m.PaymentMethod)
	user := UserResource{}
	user.FromModel(&m.User)

	u.ID = m.ID
	u.User = user
	u.Trainer = trainer
	u.ExpiredAt = m.ExpiredAt
	u.ActivedAt = m.ActivedAt
	u.Time = m.Time
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
type SkillResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (u *SkillResource) FromModel(m *model.Skill) {
	u.ID = m.ID
	u.Name = m.Name
	u.Description = m.Description
}

type SkillResources []SkillResource

func (u *SkillResources) FromModel(m []model.Skill) {
	for _, each := range m {
		var resource SkillResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}
