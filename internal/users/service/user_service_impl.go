package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	userRepo "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/password"
)

type userServiceImpl struct {
	userRepository userRepo.UserRepository
	password       password.PasswordHash
}

// CreateUser implements UserService
func (s *userServiceImpl) CreateUser(user *dto.NewUser, ctx context.Context) error {
	hashPassword, err := s.password.HashPassword(user.Password)
	if err != nil {
		return err
	}
	userModel := user.ToModel()
	userModel.Password = hashPassword
	userModel.Role = constans.Role_user
	err = s.userRepository.CreateUser(userModel, ctx)
	return err
}

// CreateAdmin implements UserService
func (s *userServiceImpl) CreateAdmin(user *dto.NewUser, ctx context.Context) error {
	hashPassword, err := s.password.HashPassword(user.Password)
	if err != nil {
		return err
	}
	userModel := user.ToModel()
	userModel.Password = hashPassword
	userModel.Role = constans.Role_admin
	err = s.userRepository.CreateUser(userModel, ctx)
	return err
}

// CreateSuperadmin implements UserService
func (s *userServiceImpl) CreateSuperadmin() error {
	isEmpty, err := s.userRepository.CheckUserIsEmpty(context.Background())
	if err != nil {
		return err
	}
	if !isEmpty {
		return nil
	}
	hashPassword, err := s.password.HashPassword(constans.Superadmin_password)
	if err != nil {
		return err
	}
	superadmin := model.User{
		Name:     constans.Superadmin_name,
		Email:    constans.Superadmin_email,
		Role:     constans.Role_superadmin,
		Password: hashPassword,
	}
	err = s.userRepository.CreateUser(&superadmin, context.Background())
	return err
}

func NewUserService(userRepository userRepo.UserRepository, password password.PasswordHash) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
		password:       password,
	}
}
