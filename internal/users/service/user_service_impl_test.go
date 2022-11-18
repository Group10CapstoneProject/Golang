package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	userRepositoryMock "github.com/Group10CapstoneProject/Golang/internal/users/repository/mock"
	passwordMock "github.com/Group10CapstoneProject/Golang/utils/password/mock"
	"github.com/stretchr/testify/suite"
)

type suiteUserService struct {
	suite.Suite
	userRepositoryMock *userRepositoryMock.UserRepositoryMock
	passwordMock       *passwordMock.PasswordMock
	userService        UserService
}

func (s *suiteUserService) SetupSuit() {
	s.userRepositoryMock = new(userRepositoryMock.UserRepositoryMock)
	s.passwordMock = new(passwordMock.PasswordMock)
	s.userService = NewUserService(s.userRepositoryMock, s.passwordMock)
}

func (s *suiteUserService) TearDown() {
	s.userRepositoryMock = nil
	s.passwordMock = nil
	s.userService = nil
}

func (s *suiteUserService) TestCreateUser() {
	testCase := []struct {
		Name            string
		Body            dto.NewUser
		ExpectedErr     error
		CreateUserErr   error
		HashPasswordRes string
		HashPasswordErr error
	}{
		{
			Name: "success",
			Body: dto.NewUser{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			ExpectedErr:     nil,
			CreateUserErr:   nil,
			HashPasswordRes: "hashpassword",
			HashPasswordErr: nil,
		},
		{
			Name: "hash password error",
			Body: dto.NewUser{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			ExpectedErr:     errors.New("hash password error"),
			CreateUserErr:   nil,
			HashPasswordRes: "",
			HashPasswordErr: errors.New("hash password error"),
		},
		{
			Name: "create user error",
			Body: dto.NewUser{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			ExpectedErr:     errors.New("create user error"),
			CreateUserErr:   errors.New("create user error"),
			HashPasswordRes: "hashpassword",
			HashPasswordErr: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuit()

			// set mock result or error
			s.passwordMock.On("HashPassword").Return(v.HashPasswordRes, v.HashPasswordErr)
			s.userRepositoryMock.On("CreateUser").Return(v.CreateUserErr)

			err := s.userService.CreateUser(&v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}
func (s *suiteUserService) TestCreateAdmin() {
	testCase := []struct {
		Name            string
		Body            dto.NewUser
		ExpectedErr     error
		CreateUserErr   error
		HashPasswordRes string
		HashPasswordErr error
	}{
		{
			Name: "success",
			Body: dto.NewUser{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			ExpectedErr:     nil,
			CreateUserErr:   nil,
			HashPasswordRes: "hashpassword",
			HashPasswordErr: nil,
		},
		{
			Name: "hash password error",
			Body: dto.NewUser{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			ExpectedErr:     errors.New("hash password error"),
			CreateUserErr:   nil,
			HashPasswordRes: "",
			HashPasswordErr: errors.New("hash password error"),
		},
		{
			Name: "create user error",
			Body: dto.NewUser{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			ExpectedErr:     errors.New("create user error"),
			CreateUserErr:   errors.New("create user error"),
			HashPasswordRes: "hashpassword",
			HashPasswordErr: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuit()

			// set mock result or error
			s.passwordMock.On("HashPassword").Return(v.HashPasswordRes, v.HashPasswordErr)
			s.userRepositoryMock.On("CreateUser").Return(v.CreateUserErr)

			err := s.userService.CreateAdmin(&v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}
func (s *suiteUserService) TestCreateSuperadmin() {
	testCase := []struct {
		Name            string
		ExpectedErr     error
		CheckUserRes    bool
		CheckUserErr    error
		CreateUserErr   error
		HashPasswordRes string
		HashPasswordErr error
	}{
		{
			Name:            "success",
			ExpectedErr:     nil,
			CheckUserRes:    true,
			CheckUserErr:    nil,
			CreateUserErr:   nil,
			HashPasswordRes: "hashpassword",
			HashPasswordErr: nil,
		},
		{
			Name:            "hash password error",
			ExpectedErr:     errors.New("hash password error"),
			CheckUserRes:    true,
			CheckUserErr:    nil,
			CreateUserErr:   nil,
			HashPasswordRes: "",
			HashPasswordErr: errors.New("hash password error"),
		},
		{
			Name:            "create user error",
			ExpectedErr:     errors.New("create user error"),
			CheckUserRes:    true,
			CheckUserErr:    nil,
			CreateUserErr:   errors.New("create user error"),
			HashPasswordRes: "hashpassword",
			HashPasswordErr: nil,
		},
		{
			Name:            "check user error",
			ExpectedErr:     errors.New("create user error"),
			CheckUserRes:    false,
			CheckUserErr:    errors.New("create user error"),
			CreateUserErr:   nil,
			HashPasswordRes: "",
			HashPasswordErr: nil,
		},
		{
			Name:            "user not empty",
			ExpectedErr:     nil,
			CheckUserRes:    false,
			CheckUserErr:    nil,
			CreateUserErr:   nil,
			HashPasswordRes: "",
			HashPasswordErr: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuit()

			// set mock result or error
			s.passwordMock.On("HashPassword").Return(v.HashPasswordRes, v.HashPasswordErr)
			s.userRepositoryMock.On("CreateUser").Return(v.CreateUserErr)
			s.userRepositoryMock.On("CheckUserIsEmpty").Return(v.CheckUserRes, v.CheckUserErr)

			err := s.userService.CreateSuperadmin()

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(suiteUserService))
}
