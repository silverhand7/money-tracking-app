package services

import (
	"context"

	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
)

type AuthServiceContract interface {
	Login(ctx context.Context, request requests.UserLoginRequest) responses.UserResponse
}
