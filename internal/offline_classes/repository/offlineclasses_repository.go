package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type OfflineClassRepository interface {
	// offline class
	CreateOfflineClass(body *model.OfflineClass, ctx context.Context) error
	FindOfflineClass(page *model.Pagination, ctx context.Context) ([]model.OfflineClass, int, error)
	FindOfflineClassById(id uint, ctx context.Context) (*model.OfflineClass, error)
	UpdateOfflineClass(body *model.OfflineClass, ctx context.Context) error
	DeleteOfflineClass(body *model.OfflineClass, ctx context.Context) error
}
