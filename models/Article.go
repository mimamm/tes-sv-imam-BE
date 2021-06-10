package models

type ArticleModel struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Content     string `json:"content" validate:"required"`
	Category    string `json:"category" validate:"required"`
	CreatedDate string `json:"created_date" validate:"required"`
	UpdatedDate string `json:"updated_date" validate:"required"`
	Status      string `json:"status" validate:"required"`
}
