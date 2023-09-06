package repositories

import (
	"context"
	"database/sql"

	"github.com/silverhand7/money-tracking-app/models/domain"
)

type UserRepositoryContract interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, category domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, category domain.User)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.User, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.User
}
