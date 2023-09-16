package services

import (
	"context"

	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
)

type TransactionServiceContract interface {
	GetAll(ctx context.Context, userId int32) []responses.TransactionResponse
	FindById(ctx context.Context, transactionId int32, userId int32) responses.TransactionResponse
	Create(ctx context.Context, request requests.TransactionCreateRequest) responses.TransactionResponse
	Update(ctx context.Context, request requests.TransactionUpdateRequest) responses.TransactionResponse
	Delete(ctx context.Context, transactionId int32, userId int32)
}
