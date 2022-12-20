package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type ArticlesRepository interface {
	// Articles
	CreateArticles(body *model.Articles, ctx context.Context) error
	FindArticles(ctx context.Context) ([]model.Articles, error)
	FindArticlesById(id uint, ctx context.Context) (*model.Articles, error)
	UpdateArticles(body *model.Articles, ctx context.Context) error
	DeleteArticles(body *model.Articles, ctx context.Context) error
}
