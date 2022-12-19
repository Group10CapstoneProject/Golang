package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/internal/articles/dto"
	articlesServ "github.com/Group10CapstoneProject/Golang/internal/articles/service"
	"github.com/labstack/echo/v4"
)

type articlesControllerImpl struct {
	articlesService articlesServ.ArticlesService
}

// CreateArticles implements ArticlesController
func (d *articlesControllerImpl) CreateArticles(c echo.Context) error {
	var articles dto.ArticlesStoreRequest
	if err := c.Bind(&articles); err != nil {
		return err
	}
	if err := c.Validate(articles); err != nil {
		return err
	}
	if err := d.articlesService.CreateArticles(&articles, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new articles succcess created",
	})
}

// DeleteArticles implements ArticlesController
func (d *articlesControllerImpl) DeleteArticles(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.articlesService.DeleteArticles(uint(id), c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete articles",
	})
}

// GetArticlesDetail implements ArticlesController
func (d *articlesControllerImpl) GetArticlesDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	articles, err := d.articlesService.FindArticlesById(uint(id), c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get articles detail",
		"data":    articles,
	})
}

// GetArticles implements ArticlesController
func (d *articlesControllerImpl) GetArticles(c echo.Context) error {
	articles, err := d.articlesService.FindArticles(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get articles",
		"data":    articles,
	})
}

// UpdateArticles implements ArticlesController
func (d *articlesControllerImpl) UpdateArticles(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var articles dto.ArticlesUpdateRequest
	if err := c.Bind(&articles); err != nil {
		return err
	}
	if err := c.Validate(articles); err != nil {
		return err
	}
	articles.ID = uint(id)
	if err := d.articlesService.UpdateArticles(&articles, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update articles",
	})
}

func NewArticlesController(articlesService articlesServ.ArticlesService) ArticlesController {
	return &articlesControllerImpl{
		articlesService: articlesService,
	}
}
