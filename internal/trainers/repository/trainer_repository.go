package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type TrainerRepository interface {
	// trainer booking
	CreateTrainerBooking(body *model.TrainerBooking, ctx context.Context) (*model.TrainerBooking, error)
	FindTrainerBookings(page *model.Pagination, ctx context.Context) ([]model.TrainerBooking, int, error)
	FindTrainerBookingById(id uint, ctx context.Context) (*model.TrainerBooking, error)
	FindTrainerBookingByUser(userId uint, ctx context.Context) ([]model.TrainerBooking, error)
	UpdateTrainerBooking(body *model.TrainerBooking, ctx context.Context) error
	DeleteTrainerBooking(body *model.TrainerBooking, ctx context.Context) error
	ReadTrainerBooking(cond *model.TrainerBooking, ctx context.Context) ([]model.TrainerBooking, error)

	// trainer
	CreateTrainer(body *model.Trainer, ctx context.Context) error
	FindTrainers(cond *model.Trainer, priceOrder string, date string, ctx context.Context) ([]model.Trainer, error)
	FindTrainerById(id uint, ctx context.Context) (*model.Trainer, error)
	UpdateTrainer(body *model.Trainer, ctx context.Context) error
	DeleteTrainer(body *model.Trainer, ctx context.Context) error

	// skill
	CreateSkill(body *model.Skill, ctx context.Context) error
	FindSkills(ctx context.Context) ([]model.Skill, error)
	FindSkillById(id uint, ctx context.Context) (*model.Skill, error)
	CheckSkillIsDeleted(body *model.Skill) error
	UpdateSkill(body *model.Skill, ctx context.Context) error
	DeleteSkill(body *model.Skill, ctx context.Context) error
}
