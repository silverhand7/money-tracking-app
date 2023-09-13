package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/domain"
)

type WalletRepository struct {
}

func (repository *WalletRepository) GetAll(ctx context.Context, tx *sql.Tx, apiKey string) []domain.Wallet {
	SQL := "SELECT * FROM wallets WHERE user_id = $1"
	rows, err := tx.QueryContext(ctx, SQL, apiKey)
	helpers.PanicIfError(err)
	defer rows.Close()

	var wallets []domain.Wallet
	for rows.Next() {
		wallet := domain.Wallet{}
		err := rows.Scan(
			&wallet.ID,
			&wallet.Name,
			&wallet.Icon,
			&wallet.Currency,
			&wallet.Balance,
			&wallet.UserID,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
		)
		helpers.PanicIfError(err)
		wallets = append(wallets, wallet)
	}

	return wallets
}

func (repository *WalletRepository) FindById(ctx context.Context, tx *sql.Tx, walletId int32) (domain.Wallet, error) {
	SQL := "SELECT * FROM wallets WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, walletId)
	helpers.PanicIfError(err)
	defer rows.Close()

	wallet := domain.Wallet{}
	if rows.Next() {
		err := rows.Scan(
			&wallet.ID,
			&wallet.Name,
			&wallet.Icon,
			&wallet.Currency,
			&wallet.Balance,
			&wallet.UserID,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
		)
		helpers.PanicIfError(err)
		return wallet, nil
	}
	return wallet, errors.New("wallet is not found")
}

func (repository *WalletRepository) Save(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet {
	SQL := "INSERT INTO wallets (name, icon, currency, balance, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	var id int32
	err := tx.QueryRowContext(
		ctx,
		SQL,
		wallet.Name,
		wallet.Icon,
		wallet.Currency,
		wallet.Balance,
		wallet.UserID,
		wallet.CreatedAt,
		wallet.UpdatedAt,
	).Scan(&id)

	helpers.PanicIfError(err)

	wallet.ID = int32(id)
	return wallet
}

func (repository *WalletRepository) Update(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet {
	SQL := `UPDATE wallets SET
	name = $2,
	icon = $3,
	currency = $4,
	balance = $5,
	user_id = $6,
	updated_at = $7
	WHERE id = $1
	RETURNING *`

	row := tx.QueryRowContext(
		ctx,
		SQL,
		wallet.ID,
		wallet.Name,
		wallet.Icon,
		wallet.Currency,
		wallet.Balance,
		wallet.UserID,
		wallet.UpdatedAt,
	)

	err := row.Scan(
		&wallet.ID,
		&wallet.Name,
		&wallet.Icon,
		&wallet.Currency,
		&wallet.Balance,
		&wallet.UserID,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	helpers.PanicIfError(err)

	return wallet
}

func (repository *WalletRepository) Delete(ctx context.Context, tx *sql.Tx, walletId int32) {
	SQL := `DELETE FROM wallets WHERE id=$1`
	_, err := tx.ExecContext(ctx, SQL, walletId)
	helpers.PanicIfError(err)
}
