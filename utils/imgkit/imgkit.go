package imgkit

import (
	"context"

	"github.com/codedius/imagekit-go"
	"github.com/google/uuid"
)

type ImagekitService interface {
	Upload(title string, file interface{}) (url string, err error)
}
type imagekitServiceImpl struct {
	PRIVATE_KEY string
	PUBLIC_KEY  string
}

func (i *imagekitServiceImpl) Upload(title string, file interface{}) (url string, err error) {
	opts := imagekit.Options{
		PrivateKey: i.PRIVATE_KEY,
		PublicKey:  i.PUBLIC_KEY,
	}
	ik, err := imagekit.NewClient(&opts)
	if err != nil {
		return "", err
	}
	name := uuid.New().String()
	ur := imagekit.UploadRequest{
		File:              file,
		FileName:          name,
		UseUniqueFileName: false,
		Tags:              []string{},
		Folder:            title,
		IsPrivateFile:     false,
		CustomCoordinates: "",
		ResponseFields:    nil,
	}

	upr, err := ik.Upload.ServerUpload(context.Background(), &ur)
	if err != nil {
		return "", err
	}

	return upr.URL, nil
}

func NewImageKitService(privkey string, pubkey string) ImagekitService {
	return &imagekitServiceImpl{
		PRIVATE_KEY: privkey,
		PUBLIC_KEY:  pubkey,
	}
}
