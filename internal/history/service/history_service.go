package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/history/dto"
)

type HistoryService interface {
	FindHistoryActivity(request *dto.HistoryActivityRequest, ctx context.Context) ([]dto.HistoryResource, error)
	FindHistoryOrder(request *dto.HistoryOrderRequest, ctx context.Context) ([]dto.HistoryResource, error)
}
