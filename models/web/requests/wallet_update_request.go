package requests

type WalletUpdateRequest struct {
	ID       int32  `validate:"required" json:"id"`
	Name     string `validate:"required,max=200,min=1" json:"name"`
	Icon     string `json:"icon"`
	Currency string `validate:"required" json:"currency"`
	Balance  int64  `json:"balance"`
	UserID   int32  `validate:"required" json:"user_id"`
}
