package controller

import "github.com/labstack/echo/v4"

type TrainerController interface {
	// trainer booking
	CreateTrainerBooking(c echo.Context) error
	GetTrainerBookings(c echo.Context) error
	GetTrainerBookingDetail(c echo.Context) error
	GetTrainerBookingUser(c echo.Context) error
	UpdateTrainerBooking(c echo.Context) error
	SetStatusTrainerBooking(c echo.Context) error
	TrainerBookingPayment(c echo.Context) error
	DeleteTrainerBooking(c echo.Context) error
	CheckTrainerBooking(c echo.Context) error
	TakeTrainerBooking(c echo.Context) error

	// trainer
	CreateTrainer(c echo.Context) error
	GetTrainers(c echo.Context) error
	GetTrainerDetail(c echo.Context) error
	UpdateTrainer(c echo.Context) error
	DeleteTrainer(c echo.Context) error

	// skill
	CreateSkill(c echo.Context) error
	GetSkills(c echo.Context) error
	GetSkillDetail(c echo.Context) error
	UpdateSkill(c echo.Context) error
	DeleteSkill(c echo.Context) error
}
