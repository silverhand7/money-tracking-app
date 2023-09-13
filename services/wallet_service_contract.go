package services

import (
	"context"

	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
)

type WalletServiceContract interface {
	Create(ctx context.Context, request requests.WalletCreateRequest) responses.WalletResponse
	Update(ctx context.Context, request requests.WalletUpdateRequest) responses.WalletResponse
	Delete(ctx context.Context, walletId int32)
	FindById(ctx context.Context, walletId int32) responses.WalletResponse
	GetAll(ctx context.Context, apiKey string) []responses.WalletResponse
}
