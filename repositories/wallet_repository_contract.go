package repositories

import (
	"context"
	"database/sql"

	"github.com/silverhand7/money-tracking-app/models/domain"
)

type WalletRepositoryContract interface {
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Wallet
	FindById(ctx context.Context, tx *sql.Tx, walletId int32) (domain.Wallet, error)
	Save(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet
	Update(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet
	Delete(ctx context.Context, tx *sql.Tx, walletId int32)
}
