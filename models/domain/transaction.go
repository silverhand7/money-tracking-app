package domain

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID         int32
	WalletID   int32
	CategoryID int32
	Nominal    sql.NullInt64
	DateTime   sql.NullTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
