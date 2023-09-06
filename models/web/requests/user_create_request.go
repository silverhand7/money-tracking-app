package requests

type UserCreateRequest struct {
	ID       int32
	Name     string
	Email    string
	Password string
}
