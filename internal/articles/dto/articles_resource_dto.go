package dto

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/model"
)

// articles resource
type ArticlesResource struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Picture   string    `json:"picture"`
	UpdatedAt time.Time `json:"updated_at"`
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
	u.Picture = m.Picture
	u.UpdatedAt = m.UpdatedAt
}

// articles detail resource
type ArticlesDetailResource struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Picture   string    `json:"picture"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *ArticlesDetailResource) FromModel(m *model.Articles) {
	u.ID = m.ID
	u.Title = m.Title
	u.Picture = m.Picture
	u.UpdatedAt = m.UpdatedAt
	u.Content = m.Content
}
