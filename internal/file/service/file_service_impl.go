package service

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/Group10CapstoneProject/Golang/internal/file/dto"
	"github.com/Group10CapstoneProject/Golang/utils/imgkit"
)

type fileServiceImpl struct {
	imgkit imgkit.ImagekitService
}

func (s *fileServiceImpl) Upload(request dto.FileStoreRequest, ctx context.Context) (u dto.UrlResource, err error) {
	title := request.Title
	buf := bytes.NewBuffer(nil)

	if _, err = io.Copy(buf, request.File); err != nil {
		return dto.UrlResource{}, err
	}

	result, err := s.imgkit.Upload(title, buf.Bytes())
	if err != nil {
		return dto.UrlResource{}, err
	}
	if result == "" {
		return dto.UrlResource{}, errors.New("internal server error")
	}

	u.Url = result
	return u, nil
}

func NewFileService(imgkit imgkit.ImagekitService) FileService {
	return &fileServiceImpl{
		imgkit: imgkit,
	}
}
