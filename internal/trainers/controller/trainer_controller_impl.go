package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/constans"
	memberServ "github.com/Group10CapstoneProject/Golang/internal/members/service"
	dtoNotif "github.com/Group10CapstoneProject/Golang/internal/notifications/dto"
	notifServ "github.com/Group10CapstoneProject/Golang/internal/notifications/service"
	"github.com/Group10CapstoneProject/Golang/internal/trainers/dto"
	trainerServ "github.com/Group10CapstoneProject/Golang/internal/trainers/service"
	"github.com/Group10CapstoneProject/Golang/model"
	jwtServ "github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type trainerControllerImpl struct {
	memberService       memberServ.MemberService
	trainerService      trainerServ.TrainerService
	jwtService          jwtServ.JWTService
	notificationService notifServ.NotificationService
}

// CreateSkill implements TrainerController
func (d *trainerControllerImpl) CreateSkill(c echo.Context) error {
	var skill dto.SkillStoreRequest
	if err := c.Bind(&skill); err != nil {
		return err
	}
	if err := c.Validate(skill); err != nil {
		return err
	}
	if err := d.trainerService.CreateSkill(&skill, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new trainer skill success created",
	})
}

// CreateTrainer implements TrainerController
func (d *trainerControllerImpl) CreateTrainer(c echo.Context) error {
	var trainer dto.TrainerStoreRequest
	if err := c.Bind(&trainer); err != nil {
		return err
	}
	if err := c.Validate(trainer); err != nil {
		return err
	}
	if err := d.trainerService.CreateTrainer(&trainer, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new trainer success created",
	})
}

// CreateTrainerBooking implements TrainerController
func (d *trainerControllerImpl) CreateTrainerBooking(c echo.Context) error {
	var trainerBooking dto.TrainerBookingStoreRequest
	if err := c.Bind(&trainerBooking); err != nil {
		return err
	}
	if err := c.Validate(trainerBooking); err != nil {
		return err
	}
	claims := d.jwtService.GetClaims(&c)
	trainerBooking.UserID = uint(claims["user_id"].(float64))
	id, err := d.trainerService.CreateTrainerBooking(&trainerBooking, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new offline class booking success created",
		"data":    echo.Map{"id": id},
	})
}

// DeleteSkill implements TrainerController
func (d *trainerControllerImpl) DeleteSkill(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.trainerService.DeleteSkill(uint(id), c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete trainer skill",
	})
}

// DeleteTrainer implements TrainerController
func (d *trainerControllerImpl) DeleteTrainer(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.trainerService.DeleteTrainer(uint(id), c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete trainer",
	})
}

// DeleteTrainerBooking implements TrainerController
func (d *trainerControllerImpl) DeleteTrainerBooking(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.trainerService.DeleteTrainerBooking(uint(id), c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete trainer booking",
	})
}

// GetSkillDetail implements TrainerController
func (d *trainerControllerImpl) GetSkillDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	skill, err := d.trainerService.FindSkillById(uint(id), c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get trainer skill",
		"data":    skill,
	})
}

// GetSkills implements TrainerController
func (d *trainerControllerImpl) GetSkills(c echo.Context) error {
	skills, err := d.trainerService.FindSkills(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get trainer skills",
		"data":    skills,
	})
}

// GetTrainerBookingDetail implements TrainerController
func (d *trainerControllerImpl) GetTrainerBookingDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	trainerBooking, err := d.trainerService.FindTrainerBookingById(uint(id), c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	claims := d.jwtService.GetClaims(&c)
	if claims["role"].(string) == constans.Role_superadmin || claims["role"].(string) == constans.Role_admin {
		notif := dtoNotif.NotificationReadRequest{
			TransactionID: uint(id),
			Title:         "Offline Class",
		}
		if err := d.notificationService.ReadNotification(&notif, c.Request().Context()); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success trainer booking",
		"data":    trainerBooking,
	})
}

// GetTrainerBookingUser implements TrainerController
func (d *trainerControllerImpl) GetTrainerBookingUser(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	trainerBooking, err := d.trainerService.FindTrainerBookingByUser(uint(claims["user_id"].(float64)), c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get trainer booking",
		"data":    trainerBooking,
	})
}

// GetTrainerBookings implements TrainerController
func (d *trainerControllerImpl) GetTrainerBookings(c echo.Context) error {
	var query model.Pagination
	query.NewPageQuery(c)

	trainerBookings, err := d.trainerService.FindTrainerBookings(&query, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get trainer bookings",
		"data":    trainerBookings,
	})
}

// GetTrainerDetail implements TrainerController
func (d *trainerControllerImpl) GetTrainerDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	trainer, err := d.trainerService.FindTrainerById(uint(id), c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	trainer.AccessTrainer = true
	claims := d.jwtService.GetClaims(&c)

	if claims["role"].(string) == constans.Role_user {
		memberUser, err := d.memberService.FindMemberByUser(uint(claims["user_id"].(float64)), c.Request().Context())
		if err != nil {
			if err != myerrors.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			trainer.AccessTrainer = false
		} else {
			trainer.AccessTrainer = memberUser.MemberType.AccessTrainer
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get trainer",
		"data":    trainer,
	})
}

// GetTrainers implements TrainerController
func (d *trainerControllerImpl) GetTrainers(c echo.Context) error {
	var filter dto.FilterTrainer
	filter.Name = c.QueryParam("name")
	filter.Date = c.QueryParam("date")
	filter.Gender = c.QueryParam("gender")
	filter.PriceOrder = c.QueryParam("price_order")

	if err := c.Validate(filter); err != nil {
		return err
	}
	trainers, err := d.trainerService.FindTrainers(&filter, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get trainers",
		"data":    trainers,
	})
}

// SetStatusTrainerBooking implements TrainerController
func (d *trainerControllerImpl) SetStatusTrainerBooking(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var status dto.SetStatusTrainerBooking
	if err := c.Bind(&status); err != nil {
		return err
	}
	if err := c.Validate(status); err != nil {
		return err
	}
	status.ID = uint(id)
	if err := d.trainerService.SetStatusTrainerBooking(&status, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success set status",
	})
}

// TrainerBookingPayment implements TrainerController
func (d *trainerControllerImpl) TrainerBookingPayment(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	form, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	src, err := form.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()
	claims := d.jwtService.GetClaims(&c)
	body := model.PaymentRequest{
		ID:       uint(id),
		UserID:   uint(claims["user_id"].(float64)),
		FileName: form.Filename,
		File:     src,
	}
	if err := c.Validate(body); err != nil {
		return err
	}
	err = d.trainerService.TrainerPayment(&body, c.Request().Context())
	if err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "payment success",
	})
}

// UpdateSkill implements TrainerController
func (d *trainerControllerImpl) UpdateSkill(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var skill dto.SkillUpdateRequest
	if err := c.Bind(&skill); err != nil {
		return err
	}
	if err := c.Validate(skill); err != nil {
		return err
	}
	skill.ID = uint(id)
	if err := d.trainerService.UpdateSkill(&skill, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update skill",
	})
}

// UpdateTrainer implements TrainerController
func (d *trainerControllerImpl) UpdateTrainer(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var trainer dto.TrainerUpdateRequest
	if err := c.Bind(&trainer); err != nil {
		return err
	}
	if err := c.Validate(trainer); err != nil {
		return err
	}
	trainer.ID = uint(id)
	if err := d.trainerService.UpdateTrainer(&trainer, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update trainer",
	})
}

// UpdateTrainerBooking implements TrainerController
func (d *trainerControllerImpl) UpdateTrainerBooking(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var trainer dto.TrainerBookingUpdateRequest
	if err := c.Bind(&trainer); err != nil {
		return err
	}
	if err := c.Validate(trainer); err != nil {
		return err
	}
	trainer.ID = uint(id)
	if err := d.trainerService.UpdateTrainerBooking(&trainer, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update trainer booking",
	})
}

// TakeTrainerBooking implements TrainerController
func (d *trainerControllerImpl) TakeTrainerBooking(c echo.Context) error {
	emailParam := c.QueryParam("email")
	codeParam := c.QueryParam("code")
	code, err := uuid.Parse(codeParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	condition := dto.TakeTrainerBooking{
		Email: emailParam,
		Code:  code,
	}
	if err := c.Validate(condition); err != nil {
		return err
	}
	err = d.trainerService.TakeTrainerBooking(&condition, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success take trainer booking",
	})
}

// CheckTrainerBooking implements TrainerController
func (d *trainerControllerImpl) CheckTrainerBooking(c echo.Context) error {
	emailParam := c.QueryParam("email")
	codeParam := c.QueryParam("code")
	code, err := uuid.Parse(codeParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	condition := dto.TakeTrainerBooking{
		Email: emailParam,
		Code:  code,
	}
	if err := c.Validate(condition); err != nil {
		return err
	}
	TrainerBooking, err := d.trainerService.CheckTrainerBooking(&condition, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get trainer booking",
		"data":    TrainerBooking,
	})
}

func NewTrainerController(memberService memberServ.MemberService, jwtService jwtServ.JWTService, notificationService notifServ.NotificationService, trainerService trainerServ.TrainerService) TrainerController {
	return &trainerControllerImpl{
		memberService:       memberService,
		trainerService:      trainerService,
		jwtService:          jwtService,
		notificationService: notificationService,
	}
}
