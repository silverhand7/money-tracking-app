package responses

import (
	"time"
)

type WalletTransactionResponse struct {
	ID         int32     `json:"id"`
	WalletID   int32     `json:"wallet_id"`
	CategoryID int32     `json:"category_id"`
	Nominal    int64     `json:"nominal"`
	Type       string    `json:"type"`
	Note       string    `json:"note"`
	DateTime   time.Time `json:"date_time"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
