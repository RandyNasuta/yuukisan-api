package models

import "time"

type Article struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	CategoryId  uint      `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryId" json:"category"`
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
