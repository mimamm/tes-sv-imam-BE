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

func CreateArticle(L *models.ArticlePayload) *ResponseModel {
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

func ReadArticleWithPagination(limit int, offset int) []models.ArticlePayload {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.ArticlePayload

	items, err := db.Query("SELECT title, content, category, status FROM posts LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	for items.Next() {
		var each = models.ArticlePayload{}
		var err = items.Scan(&each.Title, &each.Content, &each.Category, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

func ReadAllArticle() []models.ArticleModel {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.ArticleModel

	items, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	for items.Next() {
		var each = models.ArticleModel{}
		var err = items.Scan(&each.Id, &each.Title, &each.Content, &each.Category, &each.CreatedDate, &each.UpdatedDate, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

func ReadArticleById(id int) []models.ArticlePayload {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.ArticlePayload

	items, err := db.Query(`SELECT title, content, category, status FROM posts WHERE id = ?`, id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	for items.Next() {
		var each = models.ArticlePayload{}
		var err = items.Scan(&each.Title, &each.Content, &each.Category, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

func DeleteArticle(id int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("UPDATE posts SET status = 'Thrash', updated_date = CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed Delete Data"}
		return Res
	}
	Res = &ResponseModel{200, "Success Delete Data"}
	return Res
}

func UpdateArticle(L *models.ArticlePayload, id int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("UPDATE posts SET title = ?, content = ?, category = ?, status =?, updated_date = CURRENT_TIMESTAMP WHERE id = ?",
		L.Title, L.Content, L.Category, L.Status, id)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}
