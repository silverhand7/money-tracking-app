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

func (repository *WalletRepository) GetAll(ctx context.Context, tx *sql.Tx, userId int32) []domain.Wallet {
	SQL := "SELECT * FROM wallets WHERE user_id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
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

func (repository *WalletRepository) FindById(ctx context.Context, tx *sql.Tx, walletId int32, userId int32) (domain.Wallet, error) {
	SQL := "SELECT * FROM wallets WHERE id = $1 and user_id = $2"
	rows, err := tx.QueryContext(ctx, SQL, walletId, userId)
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

func (repository *WalletRepository) GetWalletTransactions(ctx context.Context, tx *sql.Tx, walletId int32, userId int32) ([]domain.WalletTransaction, error) {
	SQL := `SELECT t.*,
		c.type, c.icon, c.name
		FROM transactions t
		join wallets w on w.id = t.wallet_id
		join categories c on c.id = t.category_id
		WHERE wallet_id = $1 and user_id = $2
		ORDER BY date_time DESC`
	rows, err := tx.QueryContext(ctx, SQL, walletId, userId)
	helpers.PanicIfError(err)
	defer rows.Close()

	var walletTransactions []domain.WalletTransaction
	for rows.Next() {
		transaction := domain.WalletTransaction{}
		err := rows.Scan(
			&transaction.ID,
			&transaction.WalletID,
			&transaction.CategoryID,
			&transaction.Nominal,
			&transaction.DateTime,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.Note,
			&transaction.Type,
			&transaction.Icon,
			&transaction.Name,
		)
		helpers.PanicIfError(err)
		walletTransactions = append(walletTransactions, transaction)
	}

	if len(walletTransactions) != 0 {
		return walletTransactions, nil
	}
	return walletTransactions, errors.New("wallet or transaction not found")
}

func (repository *WalletRepository) UpdateBalance(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet {
	SQL := `UPDATE wallets SET
	balance = $2,
	updated_at = $3
	WHERE id = $1
	RETURNING *`

	row := tx.QueryRowContext(
		ctx,
		SQL,
		wallet.ID,
		wallet.Balance,
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
