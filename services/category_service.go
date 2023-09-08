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

type CategoryService struct {
	CategoryRepository repositories.CategoryRepositoryContract
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (service *CategoryService) Create(ctx context.Context, request requests.CategoryCreateRequest) responses.CategoryResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	isIconValid := false
	if len(request.Icon) > 0 {
		isIconValid = true
	}
	category := domain.Category{
		Type: request.Type,
		Name: request.Name,
		Icon: sql.NullString{
			String: request.Icon,
			Valid:  isIconValid,
		},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return responses.CategoryResponse{
		ID:        category.ID,
		Type:      category.Type,
		Name:      category.Name,
		Icon:      category.Icon,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func (service *CategoryService) Update(ctx context.Context, request requests.CategoryUpdateRequest) responses.CategoryResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	isIconValid := false
	if len(request.Icon) > 0 {
		isIconValid = true
	}
	category = domain.Category{
		ID:   category.ID,
		Type: request.Type,
		Name: request.Name,
		Icon: sql.NullString{
			String: request.Icon,
			Valid:  isIconValid,
		},
		UpdatedAt: time.Now().UTC(),
	}

	category = service.CategoryRepository.Update(ctx, tx, category)

	return responses.CategoryResponse{
		ID:        category.ID,
		Type:      category.Type,
		Name:      category.Name,
		Icon:      category.Icon,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func (service *CategoryService) Delete(ctx context.Context, categoryId int32) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category.ID)
}

func (service *CategoryService) FindById(ctx context.Context, categoryId int32) responses.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	categoryResponse := responses.CategoryResponse{
		ID:        category.ID,
		Type:      category.Type,
		Name:      category.Name,
		Icon:      category.Icon,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}

	return categoryResponse
}

func (service *CategoryService) GetAll(ctx context.Context) []responses.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	categories := service.CategoryRepository.GetAll(ctx, tx)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	var categoryResponses []responses.CategoryResponse
	for _, category := range categories {
		row := responses.CategoryResponse{
			ID:        category.ID,
			Type:      category.Type,
			Name:      category.Name,
			Icon:      category.Icon,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}

		categoryResponses = append(categoryResponses, row)
	}

	return categoryResponses
}
