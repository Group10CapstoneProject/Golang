package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type OnlineClassRepository interface {
	// online class booking
	CreateOnlineClassBooking(body *model.OnlineClassBooking, ctx context.Context) (*model.OnlineClassBooking, error)
	FindOnlineClassBookings(page *model.Pagination, ctx context.Context) ([]model.OnlineClassBooking, int, error)
	FindOnlineClassBookingById(id uint, ctx context.Context) (*model.OnlineClassBooking, error)
	FindOnlineClassBookingByUser(userId uint, ctx context.Context) ([]model.OnlineClassBooking, error)
	UpdateOnlineClassBooking(body *model.OnlineClassBooking, ctx context.Context) error
	DeleteOnlineClassBooking(body *model.OnlineClassBooking, ctx context.Context) error
	ReadOnlineClassBooking(cond *model.OnlineClassBooking, ctx context.Context) ([]model.OnlineClassBooking, error)

	// online class
	CreateOnlineClass(body *model.OnlineClass, ctx context.Context) error
	FindOnlineClasses(ctx context.Context) ([]model.OnlineClass, error)
	FindOnlineClassById(id uint, ctx context.Context) (*model.OnlineClass, error)
	UpdateOnlineClass(body *model.OnlineClass, ctx context.Context) error
	DeleteOnlineClass(body *model.OnlineClass, ctx context.Context) error

	// online class category
	CreateOnlineClassCategory(body *model.OnlineClassCategory, ctx context.Context) error
	FindOnlineClassCategories(ctx context.Context) ([]model.OnlineClassCategory, error)
	FindOnlineClassCategoryById(id uint, ctx context.Context) (*model.OnlineClassCategory, error)
	UpdateOnlineClassCategory(body *model.OnlineClassCategory, ctx context.Context) error
	DeleteOnlineClassCategory(body *model.OnlineClassCategory, ctx context.Context) error
}
