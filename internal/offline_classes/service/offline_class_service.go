package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/offline_classes/dto"
	"github.com/Group10CapstoneProject/Golang/model"
)

type OfflineClassService interface {
	// offline class booking
	CreateOfflineClassBooking(request *dto.OfflineClassBookingStoreRequest, ctx context.Context) (uint, error)
	FindOfflineClassBookings(page *model.Pagination, ctx context.Context) (*dto.OfflineClassBookingResponses, error)
	CheckOfflineClassBookings(request *dto.TakeOfflineClassBooking, ctx context.Context) (*dto.OfflineClassBookingResource, error)
	FindOfflineClassBookingById(id uint, ctx context.Context) (*dto.OfflineClassBookingDetailResource, error)
	FindOfflineClassBookingByUser(userId uint, ctx context.Context) (*dto.OfflineClassBookingResources, error)
	UpdateOfflineClassBooking(request *dto.OfflineClassBookingUpdateRequest, ctx context.Context) error
	SetStatusOfflineClassBooking(request *dto.SetStatusOfflineClassBooking, ctx context.Context) error
	TakeOfflineClassBooking(request *dto.TakeOfflineClassBooking, ctx context.Context) error
	OfflineClassPayment(request *model.PaymentRequest, ctx context.Context) error
	CancelOfflineClassBooking(id uint, userId uint, ctx context.Context) error
	DeleteOfflineClassBooking(id uint, ctx context.Context) error

	// offline class
	CreateOfflineClass(request *dto.OfflineClassStoreRequest, ctx context.Context) error
	FindOfflineClasses(ctx context.Context) (*dto.OfflineClassResources, error)
	CheckAccessOfflineClass(userId uint, offlineClassId uint, ctx context.Context) (bool, error)
	FindOfflineClassById(id uint, ctx context.Context) (*dto.OfflineClassDetailResource, error)
	UpdateOfflineClass(request *dto.OfflineClassUpdateRequest, ctx context.Context) error
	DeleteOfflineClass(id uint, ctx context.Context) error

	// offline class category
	CreateOfflineClassCategory(request *dto.OfflineClassCategoryStoreRequest, ctx context.Context) error
	FindOfflineClassCategories(ctx context.Context) (*dto.OfflineClassCategoryResources, error)
	FindOfflineClassCategoryById(id uint, ctx context.Context) (*dto.OfflineClassByCategoryResource, error)
	UpdateOfflineClassCategory(request *dto.OfflineClassCategoryUpdateRequest, ctx context.Context) error
	DeleteOfflineClassCategory(id uint, ctx context.Context) error
}
