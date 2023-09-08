package responses

import (
	"time"
)

type CategoryResponse struct {
	ID        int32     `json:"id"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
