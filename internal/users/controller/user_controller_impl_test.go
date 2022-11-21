package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	authServiceMock "github.com/Group10CapstoneProject/Golang/internal/auth/service/mock"
	userServiceMock "github.com/Group10CapstoneProject/Golang/internal/users/service/mock"
	validatorMock "github.com/Group10CapstoneProject/Golang/utils/validator/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteUserController struct {
	suite.Suite
	userServiceMock *userServiceMock.UserServiceMock
	authServiceMock *authServiceMock.AuthServiceMock
	validaorMock    *validatorMock.CustomValidatorMock
	userController  UserController
	echoNew         *echo.Echo
}

func (s *suiteUserController) SetupSuit() {
	s.userServiceMock = new(userServiceMock.UserServiceMock)
	s.authServiceMock = new(authServiceMock.AuthServiceMock)
	s.validaorMock = new(validatorMock.CustomValidatorMock)
	s.userController = NewUserController(s.userServiceMock, s.authServiceMock)
	s.echoNew = echo.New()
	s.echoNew.Validator = s.validaorMock
}

func (s *suiteUserController) TearDown() {
	s.userServiceMock = nil
	s.validaorMock = nil
	s.userController = nil
	s.echoNew = nil
}

func (s *suiteUserController) TestSignUp() {
	testCase := []struct {
		Name           string
		Body           map[string]interface{}
		ValidatorErr   error
		ExpectedStatus int
		ExpectedBody   map[string]interface{}
		ExpectedErr    error
		CreateUserErr  error
	}{
		{
			Name: "success",
			Body: map[string]interface{}{
				"name":     "test",
				"email":    "test@gmail.com",
				"password": "123456",
			},
			ValidatorErr:   nil,
			ExpectedStatus: 200,
			CreateUserErr:  nil,
			ExpectedErr:    nil,
			ExpectedBody: map[string]interface{}{
				"message": "sign up success",
			},
		},
		{
			Name: "validator error",
			Body: map[string]interface{}{
				"name":     "test",
				"email":    "test@gmail.com",
				"password": "123456",
			},
			ValidatorErr:   echo.NewHTTPError(http.StatusBadRequest, "validator error"),
			ExpectedStatus: 0,
			CreateUserErr:  nil,
			ExpectedErr:    echo.NewHTTPError(http.StatusBadRequest, "validator error"),
			ExpectedBody:   nil,
		},
		{
			Name: "create user error",
			Body: map[string]interface{}{
				"name":     "test",
				"email":    "test@gmail.com",
				"password": "123456",
			},
			ValidatorErr:   nil,
			ExpectedStatus: 0,
			CreateUserErr:  errors.New("creaet user error"),
			ExpectedErr:    echo.NewHTTPError(http.StatusBadRequest, "creaet user error"),
			ExpectedBody:   nil,
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuit()

			// setup request
			body, err := json.Marshal(v.Body)
			s.NoError(err)

			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctx := s.echoNew.NewContext(r, w)
			ctx.SetPath("/signup")

			// define mock
			s.userServiceMock.On("CreateUser").Return(v.CreateUserErr)
			s.validaorMock.On("Validate").Return(v.ValidatorErr)

			err = s.userController.Signup(ctx)

			if err != nil {
				s.Equal(v.ExpectedErr, err)
			} else {
				s.NoError(err)
				bodyRes := map[string]interface{}{}
				err = json.NewDecoder(w.Result().Body).Decode(&bodyRes)
				s.NoError(err)

				s.Equal(v.ExpectedStatus, w.Result().StatusCode)
				s.Equal(v.ExpectedBody, bodyRes)
			}

			s.TearDown()
		})
	}
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(suiteUserController))
}
