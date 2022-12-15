package service

import (
	"bytes"
	"context"
	"io"
	"strings"
	"time"

	notifRepo "github.com/Group10CapstoneProject/Golang/internal/notifications/repository"
	"github.com/Group10CapstoneProject/Golang/internal/online_classes/dto"
	onlineClassRepo "github.com/Group10CapstoneProject/Golang/internal/online_classes/repository"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/imgkit"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
)

type onlineClassServiceImpl struct {
	onlineClassRepository  onlineClassRepo.OnlineClassRepository
	notificationRepository notifRepo.NotificationRepository
	imagekitService        imgkit.ImagekitService
}

// CreateOnlineClass implements OnlineClassService
func (s *onlineClassServiceImpl) CreateOnlineClass(request *dto.OnlineClassStoreRequest, ctx context.Context) error {
	onlineClass := request.ToModel()
	path := strings.Split(onlineClass.Link, "/")
	onlineClass.Path = path[(len(path) - 1)]
	err := s.onlineClassRepository.CreateOnlineClass(onlineClass, ctx)
	return err
}

// CreateOnlineClassBooking implements OnlineClassService
func (s *onlineClassServiceImpl) CreateOnlineClassBooking(request *dto.OnlineClassBookingStoreRequest, ctx context.Context) (uint, error) {
	onlineClassBooking := request.ToModel()
	onlineClassBooking.ExpiredAt = time.Now().Add(24 * time.Hour)
	onlineClassBooking.ActivedAt = time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	onlineClassBooking.Status = model.WAITING
	result, err := s.onlineClassRepository.CreateOnlineClassBooking(onlineClassBooking, ctx)
	if err != nil {
		return 0, err
	}
	return result.ID, nil
}

// CreateOnlineClassCategory implements OnlineClassService
func (s *onlineClassServiceImpl) CreateOnlineClassCategory(request *dto.OnlineClassCategoryStoreRequest, ctx context.Context) error {
	onlineClassCategory := request.ToModel()
	err := s.onlineClassRepository.CreateOnlineClassCategory(onlineClassCategory, ctx)
	return err
}

// DeleteOnlineClass implements OnlineClassService
func (s *onlineClassServiceImpl) DeleteOnlineClass(id uint, ctx context.Context) error {
	onlineClass := model.OnlineClass{
		ID: id,
	}
	err := s.onlineClassRepository.DeleteOnlineClass(&onlineClass, ctx)
	return err
}

// DeleteOnlineClassBooking implements OnlineClassService
func (s *onlineClassServiceImpl) DeleteOnlineClassBooking(id uint, ctx context.Context) error {
	onlineClassBooking := model.OnlineClassBooking{
		ID: id,
	}
	err := s.onlineClassRepository.DeleteOnlineClassBooking(&onlineClassBooking, ctx)
	return err
}

// DeleteOnlineClassCategory implements OnlineClassService
func (s *onlineClassServiceImpl) DeleteOnlineClassCategory(id uint, ctx context.Context) error {
	onlineClassCategory := model.OnlineClassCategory{
		ID: id,
	}
	err := s.onlineClassRepository.DeleteOnlineClassCategory(&onlineClassCategory, ctx)
	return err
}

// FindOnlineClassBookingById implements OnlineClassService
func (s *onlineClassServiceImpl) FindOnlineClassBookingById(id uint, ctx context.Context) (*dto.OnlineClassBookingDetailResource, error) {
	onlineClassBooking, err := s.onlineClassRepository.FindOnlineClassBookingById(id, ctx)
	if err != nil {
		return nil, err
	}
	if onlineClassBooking.Status != model.INACTIVE && time.Now().After(onlineClassBooking.ExpiredAt) {
		onlineClassBooking.Status = model.INACTIVE
		body := model.OnlineClassBooking{
			ID:     onlineClassBooking.ID,
			Status: model.INACTIVE,
		}
		err := s.onlineClassRepository.UpdateOnlineClassBooking(&body, ctx)
		if err != nil {
			return nil, err
		}
	}
	var result dto.OnlineClassBookingDetailResource
	result.FromModel(onlineClassBooking)
	return &result, nil
}

// FindOnlineClassBookingByUser implements OnlineClassService
func (s *onlineClassServiceImpl) FindOnlineClassBookingByUser(userId uint, ctx context.Context) (*dto.OnlineClassBookingResources, error) {
	onlineClassBooking, err := s.onlineClassRepository.FindOnlineClassBookingByUser(userId, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OnlineClassBookingResources
	result.FromModel(onlineClassBooking)
	return &result, nil
}

// FindOnlineClassBookings implements OnlineClassService
func (s *onlineClassServiceImpl) FindOnlineClassBookings(page *model.Pagination, ctx context.Context) (*dto.OnlineClassBookingResponses, error) {
	onlineClassBookings, count, err := s.onlineClassRepository.FindOnlineClassBookings(page, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OnlineClassBookingResources
	result.FromModel(onlineClassBookings)

	response := dto.OnlineClassBookingResponses{
		OnlineClassBookings: result,
		Page:                uint(page.Page),
		Limit:               uint(page.Limit),
		Count:               uint(count),
	}
	return &response, nil
}

// CheckAccessOnlineClass implements OnlineClassService
func (s *onlineClassServiceImpl) CheckAccessOnlineClass(userId uint, onlineClassId uint, ctx context.Context) (bool, error) {
	cond := model.OnlineClassBooking{
		UserID:        userId,
		OnlineClassID: onlineClassId,
		Status:        model.ACTIVE,
	}
	onlineClassBooking, err := s.onlineClassRepository.ReadOnlineClassBooking(&cond, ctx)
	if err != nil {
		return false, err
	}
	if len(onlineClassBooking) == 0 {
		return false, nil
	}
	return true, nil
}

// FindOnlineClassById implements OnlineClassService
func (s *onlineClassServiceImpl) FindOnlineClassById(id uint, ctx context.Context) (*dto.OnlineClassDetailResource, error) {
	onlineClass, err := s.onlineClassRepository.FindOnlineClassById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OnlineClassDetailResource
	result.FromModel(onlineClass)
	return &result, nil
}

// FindOnlineClassCategoryById implements OnlineClassService
func (s *onlineClassServiceImpl) FindOnlineClassCategoryById(id uint, ctx context.Context) (*dto.OnlineClassByCategoryResource, error) {
	onlineClassCategory, err := s.onlineClassRepository.FindOnlineClassCategoryById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OnlineClassByCategoryResource
	result.FromModel(onlineClassCategory)
	return &result, nil
}

// FindOnlineClassCategories implements OnlineClassService
func (s *onlineClassServiceImpl) FindOnlineClassCategories(ctx context.Context) (*dto.OnlineClassCategoryResources, error) {
	onlineClassCategories, err := s.onlineClassRepository.FindOnlineClassCategories(ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OnlineClassCategoryResources
	result.FromModel(onlineClassCategories)
	return &result, nil
}

// FindOnlineClasss implements OnlineClassService
func (s *onlineClassServiceImpl) FindOnlineClasses(ctx context.Context) (*dto.OnlineClassResources, error) {
	onlineClasses, err := s.onlineClassRepository.FindOnlineClasses(ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OnlineClassResources
	result.FromModel(onlineClasses)
	return &result, nil
}

// SetStatusOnlineClassBooking implements OnlineClassService
func (s *onlineClassServiceImpl) SetStatusOnlineClassBooking(request *dto.SetStatusOnlineClassBooking, ctx context.Context) error {
	check, err := s.onlineClassRepository.FindOnlineClassBookingById(request.ID, ctx)
	if err != nil {
		return err
	}
	onlineClassBooking := request.ToModel()

	if onlineClassBooking.Status == model.ACTIVE && check.Status != model.ACTIVE && check.Status != model.INACTIVE {
		onlineClassBooking.ExpiredAt = time.Now().Add(24 * 30 * time.Duration(check.Duration) * time.Hour)
		onlineClassBooking.ActivedAt = time.Now()
	} else if onlineClassBooking.Status == model.REJECT && check.Status != model.REJECT {
		onlineClassBooking.ExpiredAt = time.Now()
	}

	if time.Now().After(check.ExpiredAt) {
		onlineClassBooking.Status = model.INACTIVE
	}

	err = s.onlineClassRepository.UpdateOnlineClassBooking(onlineClassBooking, ctx)
	return err
}

// UpdateOnlineClass implements OnlineClassService
func (s *onlineClassServiceImpl) UpdateOnlineClass(request *dto.OnlineClassUpdateRequest, ctx context.Context) error {
	onlineClass := request.ToModel()
	err := s.onlineClassRepository.UpdateOnlineClass(onlineClass, ctx)
	return err
}

// UpdateOnlineClassBooking implements OnlineClassService
func (s *onlineClassServiceImpl) UpdateOnlineClassBooking(request *dto.OnlineClassBookingUpdateRequest, ctx context.Context) error {
	onlineClassBooking := request.ToModel()
	err := s.onlineClassRepository.UpdateOnlineClassBooking(onlineClassBooking, ctx)
	return err
}

// UpdateOnlineClassCategory implements OnlineClassService
func (s *onlineClassServiceImpl) UpdateOnlineClassCategory(request *dto.OnlineClassCategoryUpdateRequest, ctx context.Context) error {
	onlineClassCategory := request.ToModel()
	err := s.onlineClassRepository.UpdateOnlineClassCategory(onlineClassCategory, ctx)
	return err
}

// OnlineClassPayment implements OnlineClassService
func (s *onlineClassServiceImpl) OnlineClassPayment(request *model.PaymentRequest, ctx context.Context) error {
	// check online class booking id
	id := request.ID
	onlineClassBooking, err := s.onlineClassRepository.FindOnlineClassBookingById(id, ctx)
	if err != nil {
		return err
	}
	if onlineClassBooking.UserID != request.UserID {
		return myerrors.ErrPermission
	}
	if onlineClassBooking.ProofPayment != "" {
		return myerrors.ErrAlredyPaid
	}
	// create file buffer
	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, request.File); err != nil {
		return err
	}
	url, err := s.imagekitService.Upload("online_class_booking", buf.Bytes())
	if err != nil {
		return err
	}
	if url == "" {
		return myerrors.ErrFailedUpload
	}
	// update online class booking
	body := model.OnlineClassBooking{
		ID:           id,
		ProofPayment: url,
		ExpiredAt:    time.Now().Add(24 * time.Hour),
		Status:       model.PENDING,
	}
	err = s.onlineClassRepository.UpdateOnlineClassBooking(&body, ctx)
	if err != nil {
		return err
	}
	// push or create notification
	notif := model.Notification{
		UserID:          onlineClassBooking.UserID,
		TransactionID:   id,
		TransactionType: "/online-classes/bookings/details",
		Title:           "Online Class",
	}
	if err := s.notificationRepository.CreateNotification(&notif, ctx); err != nil {
		return err
	}
	return nil
}

func NewOnlineClassService(onlineClassRepository onlineClassRepo.OnlineClassRepository, notificationRepository notifRepo.NotificationRepository, imagekitService imgkit.ImagekitService) OnlineClassService {
	return &onlineClassServiceImpl{
		onlineClassRepository:  onlineClassRepository,
		notificationRepository: notificationRepository,
		imagekitService:        imagekitService,
	}
}
