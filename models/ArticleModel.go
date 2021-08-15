package models

import "database/sql"

type ArticleModel struct {
	Id          int            `json:"id" validate:"required"`
	Title       string         `json:"title" validate:"required"`
	Content     string         `json:"content" validate:"required"`
	Category    string         `json:"category" validate:"required"`
	CreatedDate string         `json:"created_date" validate:"required"`
	UpdatedDate sql.NullString `json:"updated_date" validate:"omitempty"`
	Status      string         `json:"status" validate:"required"`
}
