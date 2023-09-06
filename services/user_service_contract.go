package services

import (
	"context"

	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
)

type UserServiceContract interface {
	Create(ctx context.Context, request requests.UserCreateRequest) responses.UserResponse
	Update(ctx context.Context, request requests.UserUpdateRequest) responses.UserResponse
	UpdatePassword(ctx context.Context, request requests.UserUpdatePasswordRequest) responses.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) responses.UserResponse
	GetAll(ctx context.Context) []responses.UserResponse
}