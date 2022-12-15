package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/trainers/dto"
	"github.com/Group10CapstoneProject/Golang/model"
)

type TrainerService interface {
	// trainer booking
	CreateTrainerBooking(request *dto.TrainerBookingStoreRequest, ctx context.Context) (uint, error)
	FindTrainerBookings(page *model.Pagination, ctx context.Context) (*dto.TrainerBookingResponses, error)
	FindTrainerBookingById(id uint, ctx context.Context) (*dto.TrainerBookingDetailResource, error)
	FindTrainerBookingByUser(userId uint, ctx context.Context) (*dto.TrainerBookingResources, error)
	UpdateTrainerBooking(request *dto.TrainerBookingUpdateRequest, ctx context.Context) error
	SetStatusTrainerBooking(request *dto.SetStatusTrainerBooking, ctx context.Context) error
	TrainerPayment(request *model.PaymentRequest, ctx context.Context) error
	DeleteTrainerBooking(id uint, ctx context.Context) error
	CheckTrainerBooking(request *dto.TakeTrainerBooking, ctx context.Context) (*dto.TrainerBookingResource, error)
	TakeTrainerBooking(request *dto.TakeTrainerBooking, ctx context.Context) error

	// trainer
	CreateTrainer(request *dto.TrainerStoreRequest, ctx context.Context) error
	FindTrainers(cond *dto.FilterTrainer, ctx context.Context) (*dto.TrainerResources, error)
	FindTrainerById(id uint, ctx context.Context) (*dto.TrainerDetailResource, error)
	UpdateTrainer(request *dto.TrainerUpdateRequest, ctx context.Context) error
	DeleteTrainer(id uint, ctx context.Context) error

	// skill
	CreateSkill(request *dto.SkillStoreRequest, ctx context.Context) error
	FindSkills(ctx context.Context) (*dto.SkillResources, error)
	FindSkillById(id uint, ctx context.Context) (*dto.SkillResource, error)
	UpdateSkill(request *dto.SkillUpdateRequest, ctx context.Context) error
	DeleteSkill(id uint, ctx context.Context) error
}
