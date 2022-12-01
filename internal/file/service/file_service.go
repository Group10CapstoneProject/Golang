package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/file/dto"
)

type FileService interface {
	Upload(request dto.FileStoreRequest, ctx context.Context) (url dto.UrlResource, err error)
}
