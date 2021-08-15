package routes

import (
	"article/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

//Endpoint Function
func Endpoint() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{"GET", "POST", "PUT", "UPDATE", "DELETE", "PATCH"},
	}))

	// Generate Datase
	e.POST("/database/set", service.GenerateDatabase)
	e.POST("/database/example-data", service.InsertExampleData)

	// Article Endpoint
	e.POST("/article/", service.CreateArticle)
	e.GET("/article/:limit/:offset", service.ReadArticleWithPagination)
	e.GET("/article/:id", service.ReadArticleById)
	e.DELETE("/article/:id", service.DeleteArticle)
	e.PATCH("/article/:id", service.UpdateArticle)
	e.GET("/article/all/", service.ReadAllArticle)
	e.GET("/", service.DefaultRoute)

	e.Logger.Fatal(e.Start(":1323"))
}
