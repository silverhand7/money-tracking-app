package domain

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID         int32
	WalletID   int32
	CategoryID int32
	Nominal    int64
	DateTime   time.Time
	Note       sql.NullString
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
