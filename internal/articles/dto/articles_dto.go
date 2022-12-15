package dto

import (
	"github.com/Group10CapstoneProject/Golang/model"
)

// articles store and update request
type ArticlesStoreRequest struct {
	ID          uint
	Title       string `json:"title" validate:"required,gte=1"`
	Description string `json:"description" validate:"required,gte=1"`
	Picture     string `json:"picture" validate:"required,gte=1"`
	Content     string `json:"content" validate:"required,gte=1"`
}

func (u *ArticlesStoreRequest) ToModel() *model.Articles {
	return &model.Articles{
		ID:          u.ID,
		Title:       u.Title,
		Description: u.Description,
		Picture:     u.Picture,
		Content:     u.Content,
	}
}

type ArticlesUpdateRequest struct {
	ID          uint `json:"id,omitempty"`
	Title       uint `json:"title,omitempty" validate:"omitempty,gte=1"`
	Description uint `json:"Description,omitempty" validate:"omitempty,gte=1"`
	Picture     uint `json:"Picture,omitempty" validate:"omitempty,gte=1"`
	Content     uint `json:"Content" validate:"omitempty,required,gte=1"`
}

func (u *ArticlesUpdateRequest) ToModel() *model.Articles {
	return &model.Articles{
		ID:          u.ID,
		Title:       u.ToModel().Title,
		Description: u.ToModel().Description,
		Picture:     u.ToModel().Picture,
		Content:     u.ToModel().Content,
	}
}

// articles resource
type ArticlesResource struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Content     string `json:"content"`
}

type ArticlesResources []ArticlesResource

func (u *ArticlesResources) FromModel(m []model.Articles) {
	for _, each := range m {
		var resource ArticlesResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}

func (u *ArticlesResource) FromModel(m *model.Articles) {
	u.ID = m.ID
	u.Title = m.Title
	u.Description = m.Description
	u.Picture = m.Picture
	u.Content = m.Content

}

type ArticlesResponses struct {
	Articles ArticlesResources `json:"articles"`
	Page     uint              `json:"page"`
	Limit    uint              `json:"limit"`
	Count    uint              `json:"count"`
}

// articles detail resource
type ArticlesDetailResource struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Content     string `json:"content"`
}
