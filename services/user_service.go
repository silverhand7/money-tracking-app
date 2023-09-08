package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator"
	"github.com/silverhand7/money-tracking-app/exceptions"
	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/domain"
	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
	"github.com/silverhand7/money-tracking-app/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepositoryContract
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *UserService) GetAll(ctx context.Context) []responses.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)

	var userResponses []responses.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, responses.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return userResponses
}

func (service *UserService) Create(ctx context.Context, request requests.UserCreateRequest) responses.UserResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user := domain.User{
		Name:      request.Name,
		Email:     request.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (service *UserService) FindById(ctx context.Context, userId int32) responses.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	userResponse := responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse
}

func (service *UserService) Update(ctx context.Context, request requests.UserUpdateRequest) responses.UserResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	user.Name = request.Name
	user.Email = request.Email
	user.UpdatedAt = time.Now()
	user = service.UserRepository.Update(ctx, tx, user)

	return responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (service *UserService) UpdatePassword(ctx context.Context, request requests.UserUpdatePasswordRequest) responses.UserResponse {
	panic("not implemented") // TODO: Implement
}

func (service *UserService) Delete(ctx context.Context, userId int32) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user.ID)
}
