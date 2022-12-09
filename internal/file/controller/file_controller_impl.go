package controller

import (
	"fmt"
	"net/http"

	"github.com/Group10CapstoneProject/Golang/internal/file/dto"
	fileServ "github.com/Group10CapstoneProject/Golang/internal/file/service"
	"github.com/labstack/echo/v4"
)

type fileControllerImpl struct {
	fileService fileServ.FileService
}

func (d *fileControllerImpl) Upload(c echo.Context) error {
	title := c.FormValue("title")
	form, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	src, err := form.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()
	body := dto.FileStoreRequest{
		Title:    title,
		FileName: form.Filename,
		File:     src,
	}
	if err := c.Validate(body); err != nil {
		return err
	}
	fmt.Println(body.FileName)
	result, err := d.fileService.Upload(body, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success upload",
		"data":    result,
	})
}

func NewFileController(fileService fileServ.FileService) FileController {
	return &fileControllerImpl{
		fileService: fileService,
	}
}
