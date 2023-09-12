package responses

import "time"

type WalletResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Currency  string    `json:"currency"`
	Balance   int32     `json:"balance"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
