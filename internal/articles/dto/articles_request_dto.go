package dto

import (
	"github.com/Group10CapstoneProject/Golang/model"
)

// articles store and update request
type ArticlesStoreRequest struct {
	Title   string `json:"title" validate:"required,name"`
	Picture string `json:"picture" validate:"required,url"`
	Content string `json:"content" validate:"required"`
}

func (u *ArticlesStoreRequest) ToModel() *model.Articles {
	return &model.Articles{
		Title:   u.Title,
		Picture: u.Picture,
		Content: u.Content,
	}
}

type ArticlesUpdateRequest struct {
	ID      uint
	Title   string `json:"title,omitempty" validate:"omitempty,name"`
	Picture string `json:"Picture,omitempty" validate:"omitempty,url"`
	Content string `json:"Content,omitempty" validate:"omitempty,omitempty"`
}

func (u *ArticlesUpdateRequest) ToModel() *model.Articles {
	return &model.Articles{
		ID:      u.ID,
		Title:   u.Title,
		Picture: u.Picture,
		Content: u.Content,
	}
}
