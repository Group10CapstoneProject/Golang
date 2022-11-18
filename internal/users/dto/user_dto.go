package dto

import "github.com/Group10CapstoneProject/Golang/model"

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *NewUser) ToModel() *model.User {
	return &model.User{
		Name:  u.Name,
		Email: u.Email,
	}
}
