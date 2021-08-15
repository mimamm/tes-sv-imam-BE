package service

import (
	"article/models"
	"article/repository"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

type ResponseModel struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func CreateArticle(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.ArticlePayload)
	if err := c.Bind(U); err != nil {
		return nil
	}
	if err := c.Validate(U); err != nil {
		return err
	}
	if statusValidation(U.Status) {
		Res = (*ResponseModel)(repository.CreateArticle(U))
	} else {
		Res = &ResponseModel{400, "Status: Draft, Publish, or Thrash"}
	}
	return c.JSON(http.StatusOK, Res)
}

func statusValidation(status string) bool {
	if strings.EqualFold(status, "Draft") ||
		strings.EqualFold(status, "Publish") ||
		strings.EqualFold(status, "Thrash") {
		return true
	} else {
		return false
	}
}

//Read All with Pagination
func ReadArticleWithPagination(c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))
	result := repository.ReadArticleWithPagination(limit, offset)
	return c.JSON(http.StatusOK, result)
}

//ReadAll
func ReadAllArticle(c echo.Context) error {
	result := repository.ReadAllArticle()
	return c.JSON(http.StatusOK, result)
}
func DefaultRoute(c echo.Context) error {
	result := "Hello"
	return c.JSON(http.StatusOK, result)
}

// Read by Id
func ReadArticleById(c echo.Context) error {
	id := c.Param("id")
	data, _ := strconv.Atoi(id)
	result := repository.ReadArticleById(data)
	return c.JSON(http.StatusOK, result)
}

// Delete
func DeleteArticle(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.Param("id")
	data, _ := strconv.Atoi(id)
	Res = (*ResponseModel)(repository.DeleteArticle(data))
	return c.JSON(http.StatusOK, Res)
}

// Update
func UpdateArticle(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.Param("id")
	data, _ := strconv.Atoi(id)
	U := new(models.ArticlePayload)
	if err := c.Bind(U); err != nil {
		return nil
	}
	if err := c.Validate(U); err != nil {
		return err
	}
	if statusValidation(U.Status) {
		Res = (*ResponseModel)(repository.UpdateArticle(U, data))
	} else {
		Res = &ResponseModel{400, "Status: Draft, Publish, or Thrash"}
	}
	return c.JSON(http.StatusOK, Res)
}
