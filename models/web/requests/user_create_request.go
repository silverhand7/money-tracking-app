package requests

type UserCreateRequest struct {
	Name     string `validate:"required,max=200,min=1" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,passwordValidator" json:"password"`
}
