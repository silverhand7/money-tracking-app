package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/silverhand7/money-tracking-app/exceptions"
	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
	"github.com/silverhand7/money-tracking-app/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository repositories.UserRepositoryContract
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *AuthService) Login(ctx context.Context, request requests.UserLoginRequest) responses.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	helpers.PanicIfError(err)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exceptions.NewNotFoundError("the email or password are invalid"))
	}

	userResponse := responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		ApiKey:    user.ApiKey,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse
}
