package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator"
	"github.com/silverhand7/money-tracking-app/exceptions"
	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/domain"
	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
	"github.com/silverhand7/money-tracking-app/repositories"
)

type WalletService struct {
	WalletRepository      repositories.WalletRepositoryContract
	TransactionRepository repositories.TransactionRepositoryContract
	DB                    *sql.DB
	Validate              *validator.Validate
}

func (service *WalletService) GetAll(ctx context.Context, userId int32) []responses.WalletResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	wallets := service.WalletRepository.GetAll(ctx, tx, userId)

	var walletResponses []responses.WalletResponse
	for _, wallet := range wallets {
		walletResponses = append(walletResponses, responses.WalletResponse{
			ID:        wallet.ID,
			Name:      wallet.Name,
			Icon:      wallet.Icon.String,
			Currency:  wallet.Currency,
			Balance:   wallet.Balance,
			UserID:    wallet.UserID,
			CreatedAt: wallet.CreatedAt,
			UpdatedAt: wallet.UpdatedAt,
		})
	}

	return walletResponses
}

func (service *WalletService) FindById(ctx context.Context, walletId int32, userId int32) responses.WalletResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	wallet, err := service.WalletRepository.FindById(ctx, tx, walletId, userId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	walletResponse := responses.WalletResponse{
		ID:        wallet.ID,
		Name:      wallet.Name,
		Icon:      wallet.Icon.String,
		Currency:  wallet.Currency,
		Balance:   wallet.Balance,
		UserID:    wallet.UserID,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}

	return walletResponse
}

func (service *WalletService) Create(ctx context.Context, request requests.WalletCreateRequest) responses.WalletResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	isIconValid := false
	if len(request.Icon) > 0 {
		isIconValid = true
	}
	wallet := domain.Wallet{
		Name: request.Name,
		Icon: sql.NullString{
			String: request.Icon,
			Valid:  isIconValid,
		},
		Currency:  request.Currency,
		Balance:   request.Balance,
		UserID:    request.UserID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	wallet = service.WalletRepository.Save(ctx, tx, wallet)

	service.TransactionRepository.Save(ctx, tx, domain.Transaction{
		WalletID:   wallet.ID,
		CategoryID: 1,
		Nominal:    wallet.Balance,
		DateTime:   time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Note: sql.NullString{
			String: "Initial Balance",
			Valid:  true,
		},
	})

	return responses.WalletResponse{
		ID:        wallet.ID,
		Name:      wallet.Name,
		Icon:      wallet.Icon.String,
		Currency:  wallet.Currency,
		Balance:   wallet.Balance,
		UserID:    wallet.UserID,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}

}

func (service *WalletService) Update(ctx context.Context, request requests.WalletUpdateRequest) responses.WalletResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	wallet, err := service.WalletRepository.FindById(ctx, tx, request.ID, request.UserID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	isIconValid := false
	if len(request.Icon) > 0 {
		isIconValid = true
	}
	wallet = domain.Wallet{
		ID:   request.ID,
		Name: request.Name,
		Icon: sql.NullString{
			String: request.Icon,
			Valid:  isIconValid,
		},
		Currency:  request.Currency,
		Balance:   request.Balance,
		UserID:    request.UserID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	wallet = service.WalletRepository.Update(ctx, tx, wallet)

	return responses.WalletResponse{
		ID:        wallet.ID,
		Name:      wallet.Name,
		Icon:      wallet.Icon.String,
		Currency:  wallet.Currency,
		Balance:   wallet.Balance,
		UserID:    wallet.UserID,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}
}

func (service *WalletService) Delete(ctx context.Context, walletId int32, userId int32) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	wallet, err := service.WalletRepository.FindById(ctx, tx, walletId, userId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	service.WalletRepository.Delete(ctx, tx, wallet.ID)
}

func (service *WalletService) GetWalletTransactions(ctx context.Context, walletId int32, userId int32) []responses.WalletTransactionResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	walletTransactions, err := service.WalletRepository.GetWalletTransactions(ctx, tx, walletId, userId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	var walletTransactionResponses []responses.WalletTransactionResponse
	for _, transaction := range walletTransactions {
		walletTransactionResponses = append(walletTransactionResponses, responses.WalletTransactionResponse{
			ID:         transaction.ID,
			WalletID:   transaction.WalletID,
			CategoryID: transaction.CategoryID,
			Nominal:    transaction.Nominal,
			Type:       transaction.Type,
			Note:       transaction.Note.String,
			Icon:       transaction.Icon.String,
			Name:       transaction.Name,
			DateTime:   transaction.DateTime,
			CreatedAt:  transaction.CreatedAt,
			UpdatedAt:  transaction.UpdatedAt,
		})
	}

	return walletTransactionResponses

}
