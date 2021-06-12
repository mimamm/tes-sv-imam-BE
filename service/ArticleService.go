package service

import (
	"article/models"
	"article/repository"
	"net/http"

	"github.com/labstack/echo"
)

type ResponseModel struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func CreateArticle(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.ArticleModel)
	if err := c.Bind(U); err != nil {
		return nil
	}
	Res = (*ResponseModel)(repository.CreateArticle(U))
	return c.JSON(http.StatusOK, Res)
}
