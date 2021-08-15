package service

import (
	"article/repository"
	"net/http"

	"github.com/labstack/echo"
)

type DatabaseResponse struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func GenerateDatabase(c echo.Context) error {
	Res := &DatabaseResponse{400, "Bad Request"}
	Res = (*DatabaseResponse)(repository.GenerateDatabase())
	return c.JSON(http.StatusOK, Res)
}

func InsertExampleData(c echo.Context) error {
	Res := &DatabaseResponse{400, "Bad Request"}
	Res = (*DatabaseResponse)(repository.InsertExampleData())
	return c.JSON(http.StatusOK, Res)
}
