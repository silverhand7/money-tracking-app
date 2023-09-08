package repositories

import (
	"context"
	"database/sql"

	"github.com/silverhand7/money-tracking-app/models/domain"
)

type CategoryRepositoryContract interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int32)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int32) (domain.Category, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
