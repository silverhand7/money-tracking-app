package domain

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int32
	Type      string
	Name      string
	Icon      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}
