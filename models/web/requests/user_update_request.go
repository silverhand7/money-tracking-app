package requests

type UserUpdateRequest struct {
	ID    int32  `validate:"required" json:"id"`
	Name  string `validate:"required,max=200,min=1" json:"name"`
	Email string `validate:"required,email" json:"email"`
}
