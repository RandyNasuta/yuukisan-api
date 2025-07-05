package structs

import "time"

type Article struct {
	Id          uint      `json:"id"`
	CategoryId  uint      `json:"category_id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
	Cover       string    `json:"cover"`
	ImageBase64 string    `json:"image_base64"`
	Writer      string    `json:"writer"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   uint      `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ArticleCreateRequest struct {
	CategoryId  uint      `json:"category_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Slug        string    `json:"slug" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	Cover       string    `json:"cover" binding:"required"`
	ImageBase64 string    `json:"image_base64" binding:"required"`
	Writer      string    `json:"writer" binding:"required"`
	CreatedBy   uint      `json:"created_by" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	UpdatedBy   uint      `json:"updated_by" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
}

type ArticleUpdateRequest struct {
	CategoryId  uint      `json:"category_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Slug        string    `json:"slug" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	Cover       string    `json:"cover" binding:"required"`
	ImageBase64 string    `json:"image_base64" binding:"required"`
	Writer      string    `json:"writer" binding:"required"`
	CreatedBy   uint      `json:"created_by" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	UpdatedBy   uint      `json:"updated_by" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
}
