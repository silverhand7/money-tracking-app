package services

import (
	"context"

	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
)

type WalletServiceContract interface {
	Create(ctx context.Context, request requests.WalletCreateRequest) responses.WalletResponse
	Update(ctx context.Context, request requests.WalletUpdateRequest) responses.WalletResponse
	Delete(ctx context.Context, walletId int32, userId int32)
	FindById(ctx context.Context, walletId int32, userId int32) responses.WalletResponse
	GetAll(ctx context.Context, userId int32) []responses.WalletResponse
	GetWalletTransactions(ctx context.Context, walletId int32, userId int32) []responses.TransactionResponse
}
