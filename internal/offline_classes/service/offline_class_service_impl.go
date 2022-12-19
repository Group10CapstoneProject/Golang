package service

import (
	"bytes"
	"context"
	"io"
	"time"

	notifRepo "github.com/Group10CapstoneProject/Golang/internal/notifications/repository"
	"github.com/Group10CapstoneProject/Golang/internal/offline_classes/dto"
	offlineClassRepo "github.com/Group10CapstoneProject/Golang/internal/offline_classes/repository"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/imgkit"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/google/uuid"
)

type offlineClassServiceImpl struct {
	offlineClassRepository offlineClassRepo.OfflineClassRepository
	notificationRepository notifRepo.NotificationRepository
	imagekitService        imgkit.ImagekitService
}

// CheckAccessOfflineClass implements OfflineClassService
func (s *offlineClassServiceImpl) CheckAccessOfflineClass(userId uint, offlineClassId uint, ctx context.Context) (bool, error) {
	panic("unimplemented")
}

// CreateOfflineClass implements OfflineClassService
func (s *offlineClassServiceImpl) CreateOfflineClass(request *dto.OfflineClassStoreRequest, ctx context.Context) error {
	offlineClass := request.ToModel()
	err := s.offlineClassRepository.CreateOfflineClass(offlineClass, ctx)
	return err
}

// CreateOfflineClassBooking implements OfflineClassService
func (s *offlineClassServiceImpl) CreateOfflineClassBooking(request *dto.OfflineClassBookingStoreRequest, ctx context.Context) (uint, error) {
	offlineClassBooking := request.ToModel()
	t := time.Now()
	exp := time.Now().Add(24 * time.Hour)
	offlineClassBooking.ExpiredAt = time.Date(exp.Year(), exp.Month(), exp.Day(), 23, 59, 59, 0, exp.Location())
	offlineClassBooking.ActivedAt = time.Date(0001, 1, 1, 0, 0, 0, 0, t.Location())
	offlineClassBooking.Status = model.WAITING
	result, err := s.offlineClassRepository.CreateOfflineClassBooking(offlineClassBooking, ctx)
	if err != nil {
		return 0, err
	}
	err = s.offlineClassRepository.OperationOfflineClassSlot(&model.OfflineClass{ID: offlineClassBooking.OfflineClassID}, "increment", ctx)
	if err != nil {
		return 0, err
	}
	return result.ID, nil
}

// CreateOfflineClassCategory implements OfflineClassService
func (s *offlineClassServiceImpl) CreateOfflineClassCategory(request *dto.OfflineClassCategoryStoreRequest, ctx context.Context) error {
	offlineClassCategory := request.ToModel()
	err := s.offlineClassRepository.CreateOfflineClassCategory(offlineClassCategory, ctx)
	return err
}

// DeleteOfflineClass implements OfflineClassService
func (s *offlineClassServiceImpl) DeleteOfflineClass(id uint, ctx context.Context) error {
	offlineClass := model.OfflineClass{
		ID: id,
	}
	err := s.offlineClassRepository.DeleteOfflineClass(&offlineClass, ctx)
	return err
}

// DeleteOfflineClassBooking implements OfflineClassService
func (s *offlineClassServiceImpl) DeleteOfflineClassBooking(id uint, ctx context.Context) error {
	offlineClassBooking := model.OfflineClassBooking{
		ID: id,
	}
	err := s.offlineClassRepository.DeleteOfflineClassBooking(&offlineClassBooking, ctx)
	if err != nil {
		return err
	}
	err = s.offlineClassRepository.OperationOfflineClassSlot(&model.OfflineClass{ID: offlineClassBooking.OfflineClassID}, "decrement", ctx)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOfflineClassCategory implements OfflineClassService
func (s *offlineClassServiceImpl) DeleteOfflineClassCategory(id uint, ctx context.Context) error {
	offlineClassCategory := model.OfflineClassCategory{
		ID: id,
	}
	err := s.offlineClassRepository.DeleteOfflineClassCategory(&offlineClassCategory, ctx)
	return err
}

// FindOfflineClassBookingById implements OfflineClassService
func (s *offlineClassServiceImpl) FindOfflineClassBookingById(id uint, ctx context.Context) (*dto.OfflineClassBookingDetailResource, error) {
	offlineClassBooking, err := s.offlineClassRepository.FindOfflineClassBookingById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OfflineClassBookingDetailResource
	result.FromModel(offlineClassBooking)
	return &result, nil
}

// FindOfflineClassBookingByUser implements OfflineClassService
func (s *offlineClassServiceImpl) FindOfflineClassBookingByUser(userId uint, ctx context.Context) (*dto.OfflineClassBookingResources, error) {
	offlineClassBooking, err := s.offlineClassRepository.FindOfflineClassBookingByUser(userId, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OfflineClassBookingResources
	result.FromModel(offlineClassBooking)
	return &result, nil
}

// FindOfflineClassBookings implements OfflineClassService
func (s *offlineClassServiceImpl) FindOfflineClassBookings(page *model.Pagination, ctx context.Context) (*dto.OfflineClassBookingResponses, error) {
	offlineClassBookings, count, err := s.offlineClassRepository.FindOfflineClassBookings(page, &model.OfflineClassBooking{}, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OfflineClassBookingResources
	result.FromModel(offlineClassBookings)

	response := dto.OfflineClassBookingResponses{
		OfflineClassBookings: result,
		Page:                 uint(page.Page),
		Limit:                uint(page.Limit),
		Count:                uint(count),
	}
	return &response, nil
}

// FindOfflineClassById implements OfflineClassService
func (s *offlineClassServiceImpl) FindOfflineClassById(id uint, ctx context.Context) (*dto.OfflineClassDetailResource, error) {
	offlineClass, err := s.offlineClassRepository.FindOfflineClassById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OfflineClassDetailResource
	result.FromModel(offlineClass)
	return &result, nil
}

// FindOfflineClassCategories implements OfflineClassService
func (s *offlineClassServiceImpl) FindOfflineClassCategories(ctx context.Context) (*dto.OfflineClassCategoryResources, error) {
	offlineClassCategories, err := s.offlineClassRepository.FindOfflineClassCategories(&model.OfflineClassCategory{}, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OfflineClassCategoryResources
	result.FromModel(offlineClassCategories)
	return &result, nil
}

// FindOfflineClassCategoryById implements OfflineClassService
func (s *offlineClassServiceImpl) FindOfflineClassCategoryById(id uint, ctx context.Context) (*dto.OfflineClassByCategoryResource, error) {
	offlineClassCategory, err := s.offlineClassRepository.FindOfflineClassCategoryById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OfflineClassByCategoryResource
	result.FromModel(offlineClassCategory)
	return &result, nil
}

// FindOfflineClasses implements OfflineClassService
func (s *offlineClassServiceImpl) FindOfflineClasses(ctx context.Context) (*dto.OfflineClassResources, error) {
	offlineClasses, err := s.offlineClassRepository.FindOfflineClasses(&model.OfflineClass{}, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.OfflineClassResources
	result.FromModel(offlineClasses)
	return &result, nil
}

// OfflineClassPayment implements OfflineClassService
func (s *offlineClassServiceImpl) OfflineClassPayment(request *model.PaymentRequest, ctx context.Context) error {
	// check offline class booking id
	id := request.ID
	offlineClassBooking, err := s.offlineClassRepository.FindOfflineClassBookingById(id, ctx)
	if err != nil {
		return err
	}
	if offlineClassBooking.UserID != request.UserID {
		return myerrors.ErrPermission
	}
	if offlineClassBooking.ProofPayment != "" {
		return myerrors.ErrAlredyPaid
	}
	// create file buffer
	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, request.File); err != nil {
		return err
	}
	url, err := s.imagekitService.Upload("offline_class_booking", buf.Bytes())
	if err != nil {
		return err
	}
	if url == "" {
		return myerrors.ErrFailedUpload
	}
	// update offline class booking
	exp := time.Now().Add(24 * time.Hour)
	body := model.OfflineClassBooking{
		ID:           id,
		ProofPayment: url,
		ExpiredAt:    time.Date(exp.Year(), exp.Month(), exp.Day(), 23, 59, 59, 0, exp.Location()),
		Status:       model.PENDING,
	}
	err = s.offlineClassRepository.UpdateOfflineClassBooking(&body, ctx)
	if err != nil {
		return err
	}
	// push or create notification
	notif := model.Notification{
		UserID:          offlineClassBooking.UserID,
		TransactionID:   id,
		TransactionType: "/offline-classes/bookings/details",
		Title:           "Offline Class",
	}
	if err := s.notificationRepository.CreateNotification(&notif, ctx); err != nil {
		return err
	}
	return nil
}

// SetStatusOfflineClassBooking implements OfflineClassService
func (s *offlineClassServiceImpl) SetStatusOfflineClassBooking(request *dto.SetStatusOfflineClassBooking, ctx context.Context) error {
	check, err := s.offlineClassRepository.FindOfflineClassBookingById(request.ID, ctx)
	if err != nil {
		return err
	}
	offlineClassBooking := request.ToModel()
	if offlineClassBooking.Status == model.ACTIVE && check.Status != model.ACTIVE && check.Status != model.INACTIVE {
		exp := check.OfflineClass.Time
		offlineClassBooking.ExpiredAt = time.Date(exp.Year(), exp.Month(), exp.Day(), 23, 59, 59, 0, exp.Location())
		offlineClassBooking.ActivedAt = time.Now()
		offlineClassBooking.Code = uuid.New()
	} else if offlineClassBooking.Status == model.REJECT && check.Status != model.REJECT {
		offlineClassBooking.ExpiredAt = time.Now()
		err := s.offlineClassRepository.OperationOfflineClassSlot(&model.OfflineClass{ID: check.OfflineClassID}, "decrement", ctx)
		if err != nil {
			return err
		}
	}

	if time.Now().After(check.ExpiredAt) {
		return myerrors.ErrOrderExpired
	}

	err = s.offlineClassRepository.UpdateOfflineClassBooking(offlineClassBooking, ctx)
	return err
}

// UpdateOfflineClass implements OfflineClassService
func (s *offlineClassServiceImpl) UpdateOfflineClass(request *dto.OfflineClassUpdateRequest, ctx context.Context) error {
	offlineClass := request.ToModel()
	err := s.offlineClassRepository.UpdateOfflineClass(offlineClass, ctx)
	return err
}

// UpdateOfflineClassBooking implements OfflineClassService
func (s *offlineClassServiceImpl) UpdateOfflineClassBooking(request *dto.OfflineClassBookingUpdateRequest, ctx context.Context) error {
	offlineClassBooking := request.ToModel()
	err := s.offlineClassRepository.UpdateOfflineClassBooking(offlineClassBooking, ctx)
	return err
}

// UpdateOfflineClassCategory implements OfflineClassService
func (s *offlineClassServiceImpl) UpdateOfflineClassCategory(request *dto.OfflineClassCategoryUpdateRequest, ctx context.Context) error {
	offlineClassCategory := request.ToModel()
	err := s.offlineClassRepository.UpdateOfflineClassCategory(offlineClassCategory, ctx)
	return err
}

// TakeOfflineClassBooking implements OfflineClassService
func (s *offlineClassServiceImpl) TakeOfflineClassBooking(request *dto.TakeOfflineClassBooking, ctx context.Context) error {
	cond := request.ToModel()
	offlineClassBooking, err := s.offlineClassRepository.ReadOfflineClassBookings(cond, ctx)
	if err != nil {
		return err
	}
	if len(offlineClassBooking) == 0 {
		return myerrors.ErrRecordNotFound
	}

	if offlineClassBooking[0].Status == model.DONE {
		return myerrors.ErrAlredyTake
	} else if offlineClassBooking[0].Status != model.ACTIVE {
		return myerrors.ErrRecordNotFound
	}

	offlineClassBookingUpdate := model.OfflineClassBooking{
		ID:     offlineClassBooking[0].ID,
		Status: model.DONE,
	}
	err = s.offlineClassRepository.UpdateOfflineClassBooking(&offlineClassBookingUpdate, ctx)
	if err != nil {
		return err
	}
	return nil
}

// ReadOfflineClassBookings implements OfflineClassService
func (s *offlineClassServiceImpl) CheckOfflineClassBookings(request *dto.TakeOfflineClassBooking, ctx context.Context) (*dto.OfflineClassBookingResource, error) {
	cond := request.ToModel()
	offlineClassBooking, err := s.offlineClassRepository.ReadOfflineClassBookings(cond, ctx)
	if err != nil {
		return nil, err
	}
	if len(offlineClassBooking) == 0 {
		return nil, myerrors.ErrRecordNotFound
	}

	if offlineClassBooking[0].Status == model.DONE {
		return nil, myerrors.ErrAlredyTake
	} else if offlineClassBooking[0].Status != model.ACTIVE {
		return nil, myerrors.ErrRecordNotFound
	}

	var result dto.OfflineClassBookingResource
	result.FromModel(&offlineClassBooking[0])

	return &result, nil
}

func NewOfflineClassService(offlineClassRepository offlineClassRepo.OfflineClassRepository, notificationRepository notifRepo.NotificationRepository, imagekitService imgkit.ImagekitService) OfflineClassService {
	return &offlineClassServiceImpl{
		offlineClassRepository: offlineClassRepository,
		notificationRepository: notificationRepository,
		imagekitService:        imagekitService,
	}
}
