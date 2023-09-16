package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/domain"
)

type TransactionRepository struct {
}

func (repository *TransactionRepository) GetAll(ctx context.Context, tx *sql.Tx, userId int32) []domain.Transaction {
	SQL := `SELECT t.id as id, wallet_id, category_id, nominal, date_time, t.created_at, t.updated_at
		FROM transactions t join wallets w on w.id = t.wallet_id
		WHERE user_id = $1`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helpers.PanicIfError(err)
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(
			&transaction.ID,
			&transaction.WalletID,
			&transaction.CategoryID,
			&transaction.Nominal,
			&transaction.DateTime,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		helpers.PanicIfError(err)
		transactions = append(transactions, transaction)
	}

	return transactions
}

func (repository *TransactionRepository) FindById(ctx context.Context, tx *sql.Tx, transactionId int32, userId int32) (domain.Transaction, error) {
	SQL := `SELECT t.id, wallet_id, category_id, nominal, date_time, t.created_at, t.updated_at
		FROM transactions t join wallets w on w.id = t.wallet_id
		WHERE t.id = $1 AND user_id = $2`
	rows, err := tx.QueryContext(ctx, SQL, transactionId, userId)
	helpers.PanicIfError(err)
	defer rows.Close()

	transaction := domain.Transaction{}
	for rows.Next() {
		err := rows.Scan(
			&transaction.ID,
			&transaction.WalletID,
			&transaction.CategoryID,
			&transaction.Nominal,
			&transaction.DateTime,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		helpers.PanicIfError(err)
		return transaction, nil
	}

	return transaction, errors.New("transaction is not found")
}

func (repository *TransactionRepository) Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := "INSERT INTO transactions (wallet_id, category_id, nominal, date_time, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	var id int32
	err := tx.QueryRowContext(
		ctx,
		SQL,
		transaction.WalletID,
		transaction.CategoryID,
		transaction.Nominal,
		transaction.DateTime,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	).Scan(&id)

	helpers.PanicIfError(err)

	transaction.ID = int32(id)
	return transaction
}

func (repository *TransactionRepository) Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `UPDATE transactions SET
	wallet_id = $2,
	category_id = $3,
	nominal = $4,
	date_time = $5,
	updated_at = $6
	WHERE id = $1
	RETURNING *`

	row := tx.QueryRowContext(
		ctx,
		SQL,
		transaction.ID,
		transaction.WalletID,
		transaction.CategoryID,
		transaction.Nominal,
		transaction.DateTime,
		transaction.UpdatedAt,
	)

	err := row.Scan(
		&transaction.ID,
		&transaction.WalletID,
		&transaction.CategoryID,
		&transaction.Nominal,
		&transaction.DateTime,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	helpers.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepository) Delete(ctx context.Context, tx *sql.Tx, transactionId int32) {
	SQL := `DELETE FROM transactions WHERE id=$1`
	_, err := tx.ExecContext(ctx, SQL, transactionId)
	helpers.PanicIfError(err)
}
