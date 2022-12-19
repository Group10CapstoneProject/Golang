package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	memberRepo "github.com/Group10CapstoneProject/Golang/internal/members/repository"
	notifRepo "github.com/Group10CapstoneProject/Golang/internal/notifications/repository"
	"github.com/Group10CapstoneProject/Golang/internal/trainers/dto"
	trainerRepo "github.com/Group10CapstoneProject/Golang/internal/trainers/repository"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/imgkit"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/google/uuid"
)

type trainerServiceImpl struct {
	trainerRepository      trainerRepo.TrainerRepository
	memberRepository       memberRepo.MemberRepository
	notificationRepository notifRepo.NotificationRepository
	imagekitService        imgkit.ImagekitService
}

// CreateSkill implements TrainerService
func (s *trainerServiceImpl) CreateSkill(request *dto.SkillStoreRequest, ctx context.Context) error {
	skill := request.ToModel()
	err := s.trainerRepository.CreateSkill(skill, ctx)
	if err != nil {
		return err
	}
	return nil
}

// CreateTrainer implements TrainerService
func (s *trainerServiceImpl) CreateTrainer(request *dto.TrainerStoreRequest, ctx context.Context) error {
	trainer := request.ToModel()
	fmt.Println(trainer)
	err := s.trainerRepository.CreateTrainer(trainer, ctx)
	if err != nil {
		return err
	}
	return nil
}

// CreateTrainerBooking implements TrainerService
func (s *trainerServiceImpl) CreateTrainerBooking(request *dto.TrainerBookingStoreRequest, ctx context.Context) (uint, error) {
	trainerBooking := request.ToModel()
	t := time.Now()
	if *trainerBooking.PaymentMethodID == 0 {
		member, err := s.memberRepository.ReadMembers(&model.Member{
			UserID: trainerBooking.UserID,
			Status: model.ACTIVE,
		}, ctx)
		if err != nil {
			return 0, err
		}
		if len(member) == 0 || !*member[0].MemberType.AccessTrainer {
			return 0, myerrors.ErrPaymentMethod
		}
		trainerBooking.ExpiredAt = time.Date(trainerBooking.Time.Year(), trainerBooking.Time.Month(), trainerBooking.Time.Day(), 23, 59, 59, 0, t.Location())
		trainerBooking.ActivedAt = time.Now()
		trainerBooking.ProofPayment = "https://ik.imagekit.io/rnwxyz/gymmember.png"
		trainerBooking.Code = uuid.New()
		trainerBooking.Status = model.ACTIVE
	} else {
		trainerBooking.ExpiredAt = time.Now().Add(24 * time.Hour)
		trainerBooking.ExpiredAt = time.Date(trainerBooking.ExpiredAt.Year(), trainerBooking.ExpiredAt.Month(), trainerBooking.ExpiredAt.Day(), 23, 59, 59, 0, t.Location())
		trainerBooking.ActivedAt = time.Date(0001, 1, 1, 0, 0, 0, 0, t.Location())
		trainerBooking.Status = model.WAITING
	}
	result, err := s.trainerRepository.CreateTrainerBooking(trainerBooking, ctx)
	if err != nil {
		return 0, err
	}
	return result.ID, nil
}

// DeleteSkill implements TrainerService
func (s *trainerServiceImpl) DeleteSkill(id uint, ctx context.Context) error {
	skill := model.Skill{
		ID: id,
	}
	err := s.trainerRepository.DeleteSkill(&skill, ctx)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTrainer implements TrainerService
func (s *trainerServiceImpl) DeleteTrainer(id uint, ctx context.Context) error {
	trainer := model.Trainer{
		ID: id,
	}
	err := s.trainerRepository.DeleteTrainer(&trainer, ctx)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTrainerBooking implements TrainerService
func (s *trainerServiceImpl) DeleteTrainerBooking(id uint, ctx context.Context) error {
	trainerBooking := model.TrainerBooking{
		ID: id,
	}
	err := s.trainerRepository.DeleteTrainerBooking(&trainerBooking, ctx)
	if err != nil {
		return err
	}
	return nil
}

// FindSkillById implements TrainerService
func (s *trainerServiceImpl) FindSkillById(id uint, ctx context.Context) (*dto.SkillResource, error) {
	skill, err := s.trainerRepository.FindSkillById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.SkillResource
	result.FromModel(skill)
	return &result, nil
}

// FindSkills implements TrainerService
func (s *trainerServiceImpl) FindSkills(ctx context.Context) (*dto.SkillResources, error) {
	skills, err := s.trainerRepository.FindSkills(ctx)
	if err != nil {
		return nil, err
	}
	var result dto.SkillResources
	result.FromModel(skills)
	return &result, nil
}

// FindTrainerBookingById implements TrainerService
func (s *trainerServiceImpl) FindTrainerBookingById(id uint, ctx context.Context) (*dto.TrainerBookingDetailResource, error) {
	trainerBooking, err := s.trainerRepository.FindTrainerBookingById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.TrainerBookingDetailResource
	result.FromModel(trainerBooking)
	return &result, nil
}

// FindTrainerBookingByUser implements TrainerService
func (s *trainerServiceImpl) FindTrainerBookingByUser(userId uint, ctx context.Context) (*dto.TrainerBookingResources, error) {
	trainerBooking, err := s.trainerRepository.FindTrainerBookingByUser(userId, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.TrainerBookingResources
	result.FromModel(trainerBooking)
	return &result, nil
}

// FindTrainerBookings implements TrainerService
func (s *trainerServiceImpl) FindTrainerBookings(page *model.Pagination, ctx context.Context) (*dto.TrainerBookingResponses, error) {
	trainerBookings, count, err := s.trainerRepository.FindTrainerBookings(page, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.TrainerBookingResources
	result.FromModel(trainerBookings)

	response := dto.TrainerBookingResponses{
		TrainerBookings: result,
		Page:            uint(page.Page),
		Limit:           uint(page.Limit),
		Count:           uint(count),
	}
	return &response, nil
}

// FindTrainerById implements TrainerService
func (s *trainerServiceImpl) FindTrainerById(id uint, ctx context.Context) (*dto.TrainerDetailResource, error) {
	trainer, err := s.trainerRepository.FindTrainerById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.TrainerDetailResource
	result.FromModel(trainer)

	return &result, nil
}

// FindTrainers implements TrainerService
func (s *trainerServiceImpl) FindTrainers(cond *dto.FilterTrainer, ctx context.Context) (*dto.TrainerResources, error) {
	trainers, err := s.trainerRepository.FindTrainers(&model.Trainer{
		Name:   cond.Name,
		Gender: cond.Gender,
	}, cond.PriceOrder, cond.Date, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.TrainerResources
	result.FromModel(trainers)

	return &result, nil
}

// SetStatusTrainerBooking implements TrainerService
func (s *trainerServiceImpl) SetStatusTrainerBooking(request *dto.SetStatusTrainerBooking, ctx context.Context) error {
	check, err := s.trainerRepository.FindTrainerBookingById(request.ID, ctx)
	if err != nil {
		return err
	}
	trainerBooking := request.ToModel()
	t := time.Now()
	if trainerBooking.Status == model.ACTIVE && check.Status != model.ACTIVE && check.Status != model.INACTIVE {
		trainerBooking.ExpiredAt = time.Date(check.Time.Year(), check.Time.Month(), check.Time.Day(), 23, 59, 59, 0, t.Location())
		trainerBooking.ActivedAt = time.Now()
		trainerBooking.Code = uuid.New()
	} else if trainerBooking.Status == model.REJECT && check.Status != model.REJECT {
		trainerBooking.ExpiredAt = time.Now()
	}

	if time.Now().After(check.ExpiredAt) {
		return myerrors.ErrOrderExpired
	}

	err = s.trainerRepository.UpdateTrainerBooking(trainerBooking, ctx)
	return err
}

// TrainerPayment implements TrainerService
func (s *trainerServiceImpl) TrainerPayment(request *model.PaymentRequest, ctx context.Context) error {
	// check offline class booking id
	id := request.ID
	trainerBooking, err := s.trainerRepository.FindTrainerBookingById(id, ctx)
	if err != nil {
		return err
	}
	if trainerBooking.UserID != request.UserID {
		return myerrors.ErrPermission
	}
	if trainerBooking.ProofPayment != "" {
		return myerrors.ErrAlredyPaid
	}
	// create file buffer
	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, request.File); err != nil {
		return err
	}
	url, err := s.imagekitService.Upload("trainer_booking", buf.Bytes())
	if err != nil {
		return err
	}
	if url == "" {
		return myerrors.ErrFailedUpload
	}
	// update tainer booking
	exp := time.Now().Add(24 * time.Hour)
	body := model.TrainerBooking{
		ID:           id,
		ProofPayment: url,
		ExpiredAt:    time.Date(exp.Year(), exp.Month(), exp.Day(), 23, 59, 59, 0, exp.Location()),
		Status:       model.PENDING,
	}
	err = s.trainerRepository.UpdateTrainerBooking(&body, ctx)
	if err != nil {
		return err
	}
	// push or create notification
	notif := model.Notification{
		UserID:          trainerBooking.UserID,
		TransactionID:   id,
		TransactionType: "/trainers/bookings/details",
		Title:           "Trainer",
	}
	if err := s.notificationRepository.CreateNotification(&notif, ctx); err != nil {
		return err
	}
	return nil
}

// UpdateSkill implements TrainerService
func (s *trainerServiceImpl) UpdateSkill(request *dto.SkillUpdateRequest, ctx context.Context) error {
	skill := request.ToModel()
	err := s.trainerRepository.UpdateSkill(skill, ctx)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTrainer implements TrainerService
func (s *trainerServiceImpl) UpdateTrainer(request *dto.TrainerUpdateRequest, ctx context.Context) error {
	trainer := request.ToModel()
	err := s.trainerRepository.UpdateTrainer(trainer, ctx)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTrainerBooking implements TrainerService
func (s *trainerServiceImpl) UpdateTrainerBooking(request *dto.TrainerBookingUpdateRequest, ctx context.Context) error {
	trainerBooking := request.ToModel()
	err := s.trainerRepository.UpdateTrainerBooking(trainerBooking, ctx)
	if err != nil {
		return err
	}
	return nil
}

// CheckTrainerBooking implements TrainerService
func (s *trainerServiceImpl) CheckTrainerBooking(request *dto.TakeTrainerBooking, ctx context.Context) (*dto.TrainerBookingResource, error) {
	cond := request.ToModel()
	trainerBooking, err := s.trainerRepository.ReadTrainerBooking(cond, ctx)
	if err != nil {
		return nil, err
	}
	if len(trainerBooking) == 0 {
		return nil, myerrors.ErrRecordNotFound
	}

	if trainerBooking[0].Status == model.DONE {
		return nil, myerrors.ErrAlredyTake
	} else if trainerBooking[0].Status != model.ACTIVE {
		return nil, myerrors.ErrRecordNotFound
	}

	var result dto.TrainerBookingResource
	result.FromModel(&trainerBooking[0])

	return &result, nil
}

// TakeTrainerBooking implements TrainerService
func (s *trainerServiceImpl) TakeTrainerBooking(request *dto.TakeTrainerBooking, ctx context.Context) error {
	cond := request.ToModel()
	trainerBooking, err := s.trainerRepository.ReadTrainerBooking(cond, ctx)
	if err != nil {
		return err
	}
	if len(trainerBooking) == 0 {
		return myerrors.ErrRecordNotFound
	}

	if trainerBooking[0].Status == model.DONE {
		return myerrors.ErrAlredyTake
	} else if trainerBooking[0].Status != model.ACTIVE {
		return myerrors.ErrRecordNotFound
	}

	trainerBookingUpdate := model.TrainerBooking{
		ID:     trainerBooking[0].ID,
		Status: model.DONE,
	}
	err = s.trainerRepository.UpdateTrainerBooking(&trainerBookingUpdate, ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewTrainerService(trainerRepository trainerRepo.TrainerRepository, memberRepository memberRepo.MemberRepository, notificationRepository notifRepo.NotificationRepository, imagekitService imgkit.ImagekitService) TrainerService {
	return &trainerServiceImpl{
		trainerRepository:      trainerRepository,
		notificationRepository: notificationRepository,
		imagekitService:        imagekitService,
		memberRepository:       memberRepository,
	}
}
