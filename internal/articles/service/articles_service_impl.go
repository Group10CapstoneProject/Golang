package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/articles/dto"
	articlesRepo "github.com/Group10CapstoneProject/Golang/internal/articles/repository"
	"github.com/Group10CapstoneProject/Golang/model"
)

type articlesServiceImpl struct {
	articlesRepository articlesRepo.ArticlesRepository
}

// CreateArticles implements ArticlesService

func (s *articlesServiceImpl) CreateArticles(request *dto.ArticlesStoreRequest, ctx context.Context) error {
	articles := request.ToModel()
	err := s.articlesRepository.CreateArticles(articles, ctx)
	return err
}

// DeleteArticles implements ArticlesService
func (s *articlesServiceImpl) DeleteArticles(id uint, ctx context.Context) error {
	articles := model.Articles{
		ID: id,
	}
	err := s.articlesRepository.DeleteArticles(&articles, ctx)
	return err
}

// FindArticlesById implements ArticlesService
func (s *articlesServiceImpl) FindArticlesById(id uint, ctx context.Context) (*dto.ArticlesResource, error) {
	articles, err := s.articlesRepository.FindArticlesById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.ArticlesResource
	result.FromModel(articles)
	return &result, nil
}

// FindArticles implements ArticlesService
func (s *articlesServiceImpl) FindArticles(ctx context.Context) (*dto.ArticlesResource, error) {
	articles, err := s.articlesRepository.FindArticles(ctx)
	if err != nil {
		return nil, err
	}
	var result dto.ArticlesResource
	result.FromModel(articles)
	return &result, nil
}

// UpdateArticles implements ArticlesService
func (s *articlesServiceImpl) UpdateArticles(request *dto.ArticlesUpdateRequest, ctx context.Context) error {
	articles := request.ToModel()
	err := s.articlesRepository.UpdateArticles(articles, ctx)
	return err
}

func NewArticlesService(articlesRepository articlesRepo.ArticlesRepository) ArticlesService {
	return &articlesServiceImpl{
		articlesRepository: articlesRepository,
	}
}
