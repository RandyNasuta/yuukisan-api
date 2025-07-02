package structs

import "time"

type RoleResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy uint      `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}
