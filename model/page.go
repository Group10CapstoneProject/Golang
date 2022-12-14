package model

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Pagination struct {
	Q     string `json:"q,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Page  int    `json:"page,omitempty"`
}

func (p *Pagination) NewPageQuery(c echo.Context) {
	search := c.QueryParam("q")
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 0
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 0
	}

	p.Q = search
	if limit <= 0 {
		p.Limit = 10
	} else {
		p.Limit = limit
	}
	if page <= 0 {
		p.Page = 1
	} else {
		p.Page = page
	}
}
