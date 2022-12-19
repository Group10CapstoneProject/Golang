package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
)

type articlesRepositoryImpl struct {
	db *gorm.DB
}

// CreateArticles implements ArticlesRepository
func (r *articlesRepositoryImpl) CreateArticles(body *model.Articles, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	return err
}

// DeleteArticles implements ArticlesRepository
func (r *articlesRepositoryImpl) DeleteArticles(body *model.Articles, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// FindArticles implements ArticlesRepository
func (r *articlesRepositoryImpl) FindArticles(ctx context.Context) ([]model.Articles, error) {
	articles := []model.Articles{}
	err := r.db.WithContext(ctx).Find(&articles).Error
	return articles, err
}

// FindArticlesById implements ArticlesRepository
func (r *articlesRepositoryImpl) FindArticlesById(id uint, ctx context.Context) (*model.Articles, error) {
	articles := model.Articles{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&articles).Error
	return &articles, err
}

// UpdateArticles implements ArticlesRepository
func (r *articlesRepositoryImpl) UpdateArticles(body *model.Articles, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(body).Updates(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

func NewArticlesRepository(database *gorm.DB) ArticlesRepository {
	return &articlesRepositoryImpl{
		db: database,
	}
}
