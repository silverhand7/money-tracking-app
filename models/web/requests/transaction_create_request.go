package requests

type TransactionCreateRequest struct {
	WalletID   int32  `validate:"required" json:"wallet_id"`
	CategoryID int32  `validate:"required" json:"category_id"`
	Nominal    int64  `json:"nominal"`
	DateTime   string `validate:"required" json:"date_time"`
	UserID     int32  `validate:"required" json:"user_id"`
}
