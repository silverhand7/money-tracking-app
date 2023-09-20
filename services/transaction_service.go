package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/silverhand7/money-tracking-app/exceptions"
	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/domain"
	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
	"github.com/silverhand7/money-tracking-app/repositories"
)

type TransactionService struct {
	TransactionRepository repositories.TransactionRepositoryContract
	WalletRepository      repositories.WalletRepositoryContract
	CategoryRepository    repositories.CategoryRepositoryContract
	DB                    *sql.DB
	Validate              *validator.Validate
}

func (service *TransactionService) GetAll(ctx context.Context, userId int32) []responses.TransactionResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	transactions := service.TransactionRepository.GetAll(ctx, tx, userId)

	var transactionResponses []responses.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, responses.TransactionResponse{
			ID:         transaction.ID,
			WalletID:   transaction.WalletID,
			CategoryID: transaction.CategoryID,
			Nominal:    transaction.Nominal,
			DateTime:   transaction.DateTime,
			CreatedAt:  transaction.CreatedAt,
			UpdatedAt:  transaction.UpdatedAt,
		})
	}

	return transactionResponses
}

func (service *TransactionService) FindById(ctx context.Context, transactionId int32, userId int32) responses.TransactionResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId, userId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	transactionResponse := responses.TransactionResponse{
		ID:         transaction.ID,
		WalletID:   transaction.WalletID,
		CategoryID: transaction.CategoryID,
		Nominal:    transaction.Nominal,
		DateTime:   transaction.DateTime,
		CreatedAt:  transaction.CreatedAt,
		UpdatedAt:  transaction.UpdatedAt,
	}

	return transactionResponse
}

func (service *TransactionService) Create(ctx context.Context, request requests.TransactionCreateRequest) responses.TransactionResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.CategoryID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	wallet, err := service.WalletRepository.FindById(ctx, tx, request.WalletID, request.UserID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	dateLayout := "2006-01-02 15:04:05"
	date, err := time.Parse(dateLayout, request.DateTime)
	helpers.PanicIfError(err)

	transaction := domain.Transaction{
		WalletID:   request.WalletID,
		CategoryID: request.CategoryID,
		Nominal:    request.Nominal,
		DateTime:   date,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	transaction = service.TransactionRepository.Save(ctx, tx, transaction)

	if category.Type == "E" {
		wallet.Balance -= transaction.Nominal
	} else {
		wallet.Balance += transaction.Nominal
	}

	service.WalletRepository.UpdateBalance(ctx, tx, domain.Wallet{
		ID:        wallet.ID,
		Balance:   wallet.Balance,
		UpdatedAt: time.Now().UTC(),
	})

	return responses.TransactionResponse{
		ID:         transaction.ID,
		WalletID:   transaction.WalletID,
		CategoryID: transaction.CategoryID,
		Nominal:    transaction.Nominal,
		DateTime:   transaction.DateTime,
		CreatedAt:  transaction.CreatedAt,
		UpdatedAt:  transaction.UpdatedAt,
	}
}

func (service *TransactionService) Update(ctx context.Context, request requests.TransactionUpdateRequest) responses.TransactionResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.CategoryID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	wallet, err := service.WalletRepository.FindById(ctx, tx, request.WalletID, request.UserID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	t, err := service.TransactionRepository.FindById(ctx, tx, request.ID, request.UserID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	wallet.Balance += t.Nominal

	fmt.Println(wallet.Balance)

	dateLayout := "2006-01-02 15:04:05"
	date, err := time.Parse(dateLayout, request.DateTime)
	helpers.PanicIfError(err)

	transaction := domain.Transaction{
		ID:         request.ID,
		WalletID:   request.WalletID,
		CategoryID: request.CategoryID,
		Nominal:    request.Nominal,
		DateTime:   date,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	transaction = service.TransactionRepository.Update(ctx, tx, transaction)

	if category.Type == "E" {
		wallet.Balance -= transaction.Nominal
	} else {
		wallet.Balance += transaction.Nominal
	}

	service.WalletRepository.UpdateBalance(ctx, tx, domain.Wallet{
		ID:        wallet.ID,
		Balance:   wallet.Balance,
		UpdatedAt: time.Now().UTC(),
	})

	return responses.TransactionResponse{
		ID:         transaction.ID,
		WalletID:   transaction.WalletID,
		CategoryID: transaction.CategoryID,
		Nominal:    transaction.Nominal,
		DateTime:   transaction.DateTime,
		CreatedAt:  transaction.CreatedAt,
		UpdatedAt:  transaction.UpdatedAt,
	}
}

func (service *TransactionService) Delete(ctx context.Context, transactionId int32, userId int32) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId, userId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	service.TransactionRepository.Delete(ctx, tx, transaction.ID)
}
