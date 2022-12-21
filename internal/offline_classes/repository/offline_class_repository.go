package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type OfflineClassRepository interface {
	// offline class
	CreateOfflineClass(body *model.OfflineClass, ctx context.Context) error
	FindOfflineClasses(cond *model.OfflineClass, title string, ctx context.Context) ([]model.OfflineClass, error)
	FindOfflineClassById(id uint, ctx context.Context) (*model.OfflineClass, error)
	UpdateOfflineClass(body *model.OfflineClass, ctx context.Context) error
	OperationOfflineClassSlot(body *model.OfflineClass, operation string, ctx context.Context) error
	DeleteOfflineClass(body *model.OfflineClass, ctx context.Context) error
	// offline class category
	CreateOfflineClassCategory(body *model.OfflineClassCategory, ctx context.Context) error
	FindOfflineClassCategories(cond *model.OfflineClassCategory, ctx context.Context) ([]model.OfflineClassCategory, error)
	FindOfflineClassCategoryById(id uint, ctx context.Context) (*model.OfflineClassCategory, error)
	UpdateOfflineClassCategory(body *model.OfflineClassCategory, ctx context.Context) error
	DeleteOfflineClassCategory(body *model.OfflineClassCategory, ctx context.Context) error
	// offline class booking
	CreateOfflineClassBooking(body *model.OfflineClassBooking, ctx context.Context) (*model.OfflineClassBooking, error)
	FindOfflineClassBookings(page *model.Pagination, cond *model.OfflineClassBooking, ctx context.Context) ([]model.OfflineClassBooking, int, error)
	ReadOfflineClassBookings(cond *model.OfflineClassBooking, ctx context.Context) ([]model.OfflineClassBooking, error)
	FindOfflineClassBookingByUser(userId uint, ctx context.Context) ([]model.OfflineClassBooking, error)
	FindOfflineClassBookingById(id uint, ctx context.Context) (*model.OfflineClassBooking, error)
	UpdateOfflineClassBooking(body *model.OfflineClassBooking, ctx context.Context) error
	DeleteOfflineClassBooking(body *model.OfflineClassBooking, ctx context.Context) error
}
