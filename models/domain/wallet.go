package domain

import (
	"database/sql"
	"time"
)

type Wallet struct {
	ID        int32
	Name      string
	Icon      sql.NullString
	Currency  string
	Balance   int32
	UserID    int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
