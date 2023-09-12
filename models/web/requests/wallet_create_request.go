package requests

type WalletCreateRequest struct {
	Name     string `validate:"required,max=200,min=1" json:"name"`
	Icon     string `json:"icon"`
	Currency string `validate:"required" json:"currency"`
	Balance  int32  `json:"balance"`
	UserID   int32  `validate:"required" json:"user_id"`
}
