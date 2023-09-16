package domain

import (
	"time"
)

type Transaction struct {
	ID         int32
	WalletID   int32
	CategoryID int32
	Nominal    int64
	DateTime   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
