package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/online_classes/dto"
	"github.com/Group10CapstoneProject/Golang/model"
)

type OnlineClassService interface {
	// online class booking
	CreateOnlineClassBooking(request *dto.OnlineClassBookingStoreRequest, ctx context.Context) (uint, error)
	FindOnlineClassBookings(page *model.Pagination, ctx context.Context) (*dto.OnlineClassBookingResponses, error)
	FindOnlineClassBookingById(id uint, ctx context.Context) (*dto.OnlineClassBookingDetailResource, error)
	FindOnlineClassBookingByUser(userId uint, ctx context.Context) (*dto.OnlineClassBookingResources, error)
	UpdateOnlineClassBooking(request *dto.OnlineClassBookingUpdateRequest, ctx context.Context) error
	SetStatusOnlineClassBooking(request *dto.SetStatusOnlineClassBooking, ctx context.Context) error
	CancelOnlineClassBooking(id uint, userId uint, ctx context.Context) error
	OnlineClassPayment(request *model.PaymentRequest, ctx context.Context) error
	DeleteOnlineClassBooking(id uint, ctx context.Context) error

	// online class
	CreateOnlineClass(request *dto.OnlineClassStoreRequest, ctx context.Context) error
	FindOnlineClasses(q string, ctx context.Context) (*dto.OnlineClassResources, error)
	CheckAccessOnlineClass(userId uint, onlineClassId uint, ctx context.Context) (bool, error)
	FindOnlineClassById(id uint, ctx context.Context) (*dto.OnlineClassDetailResource, error)
	UpdateOnlineClass(request *dto.OnlineClassUpdateRequest, ctx context.Context) error
	DeleteOnlineClass(id uint, ctx context.Context) error

	// online class category
	CreateOnlineClassCategory(request *dto.OnlineClassCategoryStoreRequest, ctx context.Context) error
	FindOnlineClassCategories(ctx context.Context) (*dto.OnlineClassCategoryResources, error)
	FindOnlineClassCategoryById(id uint, ctx context.Context) (*dto.OnlineClassByCategoryResource, error)
	UpdateOnlineClassCategory(request *dto.OnlineClassCategoryUpdateRequest, ctx context.Context) error
	DeleteOnlineClassCategory(id uint, ctx context.Context) error
}
