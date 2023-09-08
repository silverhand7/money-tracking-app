package services

import (
	"context"

	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
)

type CategoryServiceContract interface {
	Create(ctx context.Context, request requests.CategoryCreateRequest) responses.CategoryResponse
	Update(ctx context.Context, request requests.CategoryUpdateRequest) responses.CategoryResponse
	Delete(ctx context.Context, categoryId int32)
	FindById(ctx context.Context, categoryId int32) responses.CategoryResponse
	GetAll(ctx context.Context) []responses.CategoryResponse
}
