package dto

// OfflineClass and update request
type OfflineClassRequest struct {
	ID         uint
	Duration   uint `json:"duration" validate:"required,gte=1"`
	slot       uint `json:"slot" validate:"required,gte=1"`
	slot_ready uint `json:"slot_ready" validate:"required,gte=1"`
}
