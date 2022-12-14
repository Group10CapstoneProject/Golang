package dto

import "mime/multipart"

type FileStoreRequest struct {
	Title    string         `validate:"required,title"`
	FileName string         `validate:"required,image"`
	File     multipart.File `validate:"required,file"`
}

type UrlResource struct {
	Url string `json:"url"`
}
