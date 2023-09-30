package domain

import (
	"database/sql"
	"time"
)

type WalletTransaction struct {
	ID         int32
	WalletID   int32
	CategoryID int32
	Nominal    int64
	Type       string
	Note       sql.NullString
	DateTime   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
