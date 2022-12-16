package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	userRepository "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/password"
)

type userServiceImpl struct {
	userRepository userRepository.UserRepository
	password       password.PasswordHash
}

// FindAdmins implements UserService
func (*userServiceImpl) FindAdmins(page model.Pagination, ctx context.Context) (*dto.PageResponse, error) {
	panic("unimplemented")
}

// FindUser implements UserService
func (r *userServiceImpl) FindUser(userId *uint, ctx context.Context) (*dto.UserResponse, error) {
	user, err := r.userRepository.FindUserByID(userId, ctx)
	if err != nil {
		return nil, err
	}
	var userResponse dto.UserResponse
	userResponse.FromModel(user)
	return &userResponse, err
}

// FindUsers implements UserService
func (r *userServiceImpl) FindUsers(page model.Pagination, role string, ctx context.Context) (*dto.PageResponse, error) {
	users, count, err := r.userRepository.FindUsers(&page, role, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.UsersResponse
	result.FromModel(users)

	response := dto.PageResponse{
		Users: result,
		Page:  uint(page.Page),
		Limit: uint(page.Limit),
		Count: count,
	}
	return &response, nil
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
func (s *userServiceImpl) CreateSuperadmin(superadmin *model.User) error {
	isEmpty, err := s.userRepository.CheckUserIsEmpty(context.Background())
	if err != nil {
		return err
	}
	if !isEmpty {
		return nil
	}
	hashPassword, err := s.password.HashPassword(superadmin.Password)
	if err != nil {
		return err
	}
	model := model.User{
		ID:       1,
		Name:     superadmin.Name,
		Email:    superadmin.Email,
		Role:     constans.Role_superadmin,
		Password: hashPassword,
	}
	err = s.userRepository.CreateUser(&model, context.Background())
	return err
}

func NewUserService(userRepository userRepository.UserRepository, password password.PasswordHash) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
		password:       password,
	}
}
