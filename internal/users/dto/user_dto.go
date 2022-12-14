package dto

import "github.com/Group10CapstoneProject/Golang/model"

type NewUser struct {
	Name     string `json:"name" validate:"required,personname"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *NewUser) ToModel() *model.User {
	return &model.User{
		Name:  u.Name,
		Email: u.Email,
	}
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (u *UserResponse) FromModel(model *model.User) {
	u.ID = model.ID
	u.Name = model.Name
	u.Email = model.Email
	u.Role = model.Role
}

type UsersResponse []UserResponse

func (u *UsersResponse) FromModel(model []model.User) {
	for _, each := range model {
		var user UserResponse
		user.FromModel(&each)
		*u = append(*u, user)
	}
}

type UpdateUser struct {
	ID    uint
	Name  string `json:"name,omitempty" validate:"omitempty,personname"`
	Email string `json:"email,omitempty" validate:"omitempty,email"`
}

func (u *UpdateUser) ToModel() *model.User {
	return &model.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

type PageResponse struct {
	Users UsersResponse `json:"users"`
	Page  uint          `json:"page"`
	Limit uint          `json:"limit"`
	Count int           `json:"count"`
}
