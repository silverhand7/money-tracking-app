package requests

type TransactionCreateRequest struct {
	WalletID   int32  `validate:"required" json:"wallet_id"`
	CategoryID int32  `validate:"required" json:"category_id"`
	Nominal    int64  `json:"nominal"`
	Note       string `json:"note"`
	DateTime   string `validate:"required" json:"date_time"`
	UserID     int32  `validate:"required" json:"user_id"`
}
