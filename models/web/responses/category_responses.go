package responses

import (
	"database/sql"
	"time"
)

type CategoryResponse struct {
	ID        int32          `json:"id"`
	Type      string         `json:"type"`
	Name      string         `json:"name"`
	Icon      sql.NullString `json:"icon"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
