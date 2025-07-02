package structs

import "time"

type Article struct {
	Id         uint      `json:"id"`
	CategoryId uint      `json:"category_id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	Cover      string    `json:"cover"`
	Writer     string    `json:"writer"`
	CreatedBy  uint      `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedBy  uint      `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at"`
}
