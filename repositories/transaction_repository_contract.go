package repositories

import (
	"context"
	"database/sql"

	"github.com/silverhand7/money-tracking-app/models/domain"
)

type TransactionRepositoryContract interface {
	GetAll(ctx context.Context, tx *sql.Tx, userId int32) []domain.Transaction
	FindById(ctx context.Context, tx *sql.Tx, transactionId int32, userId int32) (domain.Transaction, error)
	Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Delete(ctx context.Context, tx *sql.Tx, transactionId int32)
}
