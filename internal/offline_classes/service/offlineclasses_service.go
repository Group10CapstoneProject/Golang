package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/members/dto"
	"github.com/Group10CapstoneProject/Golang/model"
)

type OfflineClassService interface {
	// offline class
	CreateOfflineClass(request *dto.OfflineClassRequest, ctx context.Context) error
	FindOfflineClass(page *model.Pagination, ctx context.Context) (*dto.OfflineClassResponses, error)
	FindOfflineClassById(id uint, ctx context.Context) (*dto.OfflineClassDetailResource, error)
	UpdateOfflineClass(request *dto.OfflineClassUpdateRequest, ctx context.Context) error
	DeleteOfflineClass(id uint, ctx context.Context) error
}
