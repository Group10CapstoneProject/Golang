package controller

import "github.com/labstack/echo/v4"

type OnlineClassController interface {
	// online class booking
	CreateOnlineClassBooking(c echo.Context) error
	GetOnlineClassBookings(c echo.Context) error
	GetOnlineClassBookingDetail(c echo.Context) error
	GetOnlineClassBookingUser(c echo.Context) error
	UpdateOnlineClassBooking(c echo.Context) error
	SetStatusOnlineClassBooking(c echo.Context) error
	OnlineClassBookingPayment(c echo.Context) error
	DeleteOnlineClassBooking(c echo.Context) error
	CancelOnlineClassBooking(c echo.Context) error

	// online class
	CreateOnlineClass(c echo.Context) error
	GetOnlineClasses(c echo.Context) error
	GetOnlineClassDetail(c echo.Context) error
	UpdateOnlineClass(c echo.Context) error
	DeleteOnlineClass(c echo.Context) error

	// online class category
	CreateOnlineClassCategory(c echo.Context) error
	GetOnlineClassCategories(c echo.Context) error
	GetOnlineClassCategoryDetail(c echo.Context) error
	UpdateOnlineClassCategory(c echo.Context) error
	DeleteOnlineClassCategory(c echo.Context) error
}
