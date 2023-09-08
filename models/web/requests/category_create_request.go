package requests

type CategoryCreateRequest struct {
	Type string `validate:"required,categoryTypeValidator" json:"type"`
	Name string `validate:"required,max=200,min=1" json:"name"`
	Icon string `json:"icon"`
}
