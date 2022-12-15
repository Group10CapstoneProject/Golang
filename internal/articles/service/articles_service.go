package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/articles/dto"
	"github.com/Group10CapstoneProject/Golang/model"
)

type ArticlesService interface {
	// articles
	CreateArticles(request *dto.ArticlesStoreRequest, ctx context.Context) error
	FindArticles(page *model.Pagination, ctx context.Context) (*dto.ArticlesResponses, error)
	FindArticlesById(id uint, ctx context.Context) (*dto.ArticlesResource, error)
	UpdateArticles(request *dto.ArticlesUpdateRequest, ctx context.Context) error
	DeleteArticles(id uint, ctx context.Context) error
}
