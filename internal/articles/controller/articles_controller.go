package controller

import (
	"github.com/labstack/echo/v4"
)

type ArticlesController interface {
	CreateArticles(c echo.Context) error
	GetArticles(c echo.Context) error
	GetArticlesDetail(c echo.Context) error
	UpdateArticles(c echo.Context) error
	DeleteArticles(c echo.Context) error
}
