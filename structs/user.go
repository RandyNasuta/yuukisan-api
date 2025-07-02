package structs

import "time"

type User struct {
	Id              uint      `json:"id"`
	RoleId          uint      `json:"role_id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	PhotoProfile    string    `json:"photo_profile"`
	Bio             string    `json:"bio"`
	LastLoginAt     string    `json:"last_login_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
