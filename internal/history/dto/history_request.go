package dto

import "github.com/Group10CapstoneProject/Golang/model"

type HistoryActivityRequest struct {
	UserID uint             `json:"user_id"`
	Status model.StatusType `json:"status" validate:"required,activity"`
	Type   string           `json:"type"`
}

type HistoryOrderRequest struct {
	UserID uint   `json:"user_id"`
	Type   string `json:"type"`
}
