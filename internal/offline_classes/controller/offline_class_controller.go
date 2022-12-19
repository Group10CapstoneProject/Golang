package controller

import "github.com/labstack/echo/v4"

type OfflineClassController interface {
	// offline class booking
	CreateOfflineClassBooking(c echo.Context) error
	GetOfflineClassBookings(c echo.Context) error
	GetOfflineClassBookingDetail(c echo.Context) error
	GetOfflineClassBookingUser(c echo.Context) error
	UpdateOfflineClassBooking(c echo.Context) error
	SetStatusOfflineClassBooking(c echo.Context) error
	OfflineClassBookingPayment(c echo.Context) error
	DeleteOfflineClassBooking(c echo.Context) error
	TakeOfflineClassBooking(c echo.Context) error
	CancelOfflineClassBooking(c echo.Context) error
	CheckOfflineClassBooking(c echo.Context) error

	// offline class
	CreateOfflineClass(c echo.Context) error
	GetOfflineClasses(c echo.Context) error
	GetOfflineClassDetail(c echo.Context) error
	UpdateOfflineClass(c echo.Context) error
	DeleteOfflineClass(c echo.Context) error

	// offline class category
	CreateOfflineClassCategory(c echo.Context) error
	GetOfflineClassCategories(c echo.Context) error
	GetOfflineClassCategoryDetail(c echo.Context) error
	UpdateOfflineClassCategory(c echo.Context) error
	DeleteOfflineClassCategory(c echo.Context) error
}
