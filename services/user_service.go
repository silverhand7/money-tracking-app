package services

import (
	"context"
	"database/sql"

	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
	"github.com/silverhand7/money-tracking-app/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepositoryContract
	DB             *sql.DB
}

func (service *UserService) GetAll(ctx context.Context) []responses.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)

	var userResponses []responses.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return userResponses
}

func (service *UserService) Create(ctx context.Context, request requests.UserCreateRequest) responses.UserResponse {
	panic("not implemented") // TODO: Implement
}

func (service *UserService) Update(ctx context.Context, request requests.UserUpdateRequest) responses.UserResponse {
	panic("not implemented") // TODO: Implement
}

func (service *UserService) UpdatePassword(ctx context.Context, request requests.UserUpdatePasswordRequest) responses.UserResponse {
	panic("not implemented") // TODO: Implement
}

func (service *UserService) Delete(ctx context.Context, userId int) {
	panic("not implemented") // TODO: Implement
}

func (service *UserService) FindById(ctx context.Context, userId int) responses.UserResponse {
	panic("not implemented") // TODO: Implement
}
