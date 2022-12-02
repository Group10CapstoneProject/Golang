package dto

import "mime/multipart"

type PaymMemberStoreRequest struct {
	ID       uint           `validate:"required"`
	UserID   uint           `validate:"required"`
	FileName string         `validate:"required,image"`
	File     multipart.File `validate:"required,file"`
}
