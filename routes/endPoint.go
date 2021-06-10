package routes

import (
	"article/service"

	"github.com/labstack/echo"
)

//RoleEndPoint function
func Endpoint() {
	e := echo.New()
	//roles endpoint
	e.POST("/article/", service.CreateArticle)

	e.Logger.Fatal(e.Start(":1323"))
}
