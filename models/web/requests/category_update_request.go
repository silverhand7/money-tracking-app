package requests

type CategoryUpdateRequest struct {
	ID   int32  `validate:"required" json:"id"`
	Type string `validate:"required,categoryTypeValidator" json:"type"`
	Name string `validate:"required,max=200,min=1" json:"name"`
	Icon string `validate:"required" json:"icon"`
}
