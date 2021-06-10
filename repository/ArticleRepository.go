package repository

import (
	"fmt"

	"article/driver"
	"article/models"
)

type ResponseModel struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func CreateArticle(L *models.ArticleModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec(`INSERT INTO posts(title, content, category, status) VALUES (?, ?, ?, ?)`,
		L.Title, L.Content, L.Category, L.Status)

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("insert success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}
