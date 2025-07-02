package models

import "time"

type Role struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   uint      `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}
