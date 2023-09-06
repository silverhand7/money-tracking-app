package repositories

import (
	"context"
	"database/sql"

	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/domain"
)

type UserRepository struct {
}

func NewUserRepository() UserRepositoryContract {
	return &UserRepository{}
}

func (userRepository *UserRepository) GetAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, name, email, password, created_at, updated_at FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		helpers.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (userRepository *UserRepository) Save(ctx context.Context, tx *sql.Tx, category domain.User) domain.User {
	panic("not implemented") // TODO: Implement
}

func (userRepository *UserRepository) Update(ctx context.Context, tx *sql.Tx, category domain.User) domain.User {
	panic("not implemented") // TODO: Implement
}

func (userRepository *UserRepository) Delete(ctx context.Context, tx *sql.Tx, category domain.User) {
	panic("not implemented") // TODO: Implement
}

func (userRepository *UserRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.User, error) {
	panic("not implemented") // TODO: Implement
}
