package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/articles/dto"
)

type ArticlesService interface {
	// articles
	CreateArticles(request *dto.ArticlesStoreRequest, ctx context.Context) error
	FindArticles(ctx context.Context) (*dto.ArticlesResources, error)
	FindArticlesById(id uint, ctx context.Context) (*dto.ArticlesDetailResource, error)
	UpdateArticles(request *dto.ArticlesUpdateRequest, ctx context.Context) error
	DeleteArticles(id uint, ctx context.Context) error
}
