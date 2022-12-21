package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/dashboard/dto"
)

type DashboardService interface {
	GetDashboard(ctx context.Context) (*dto.DashboardDto, error)
}
