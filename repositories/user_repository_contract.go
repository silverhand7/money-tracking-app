package repositories

import (
	"context"
	"database/sql"

	"github.com/silverhand7/money-tracking-app/models/domain"
)

type UserRepositoryContract interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, userId int32)
	FindById(ctx context.Context, tx *sql.Tx, userId int32) (domain.User, error)
	FindByApiKey(ctx context.Context, tx *sql.Tx, apiKey string) (domain.User, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
}
