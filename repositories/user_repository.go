package repositories

import (
	"context"
	"database/sql"

	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/domain"
	"golang.org/x/crypto/bcrypt"
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

func (userRepository *UserRepository) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	var id int32
	err = tx.QueryRowContext(
		ctx,
		SQL,
		user.Name,
		user.Email,
		hashedPassword,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&id)

	helpers.PanicIfError(err)

	user.ID = int32(id)
	return user
}

func (userRepository *UserRepository) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	panic("not implemented") // TODO: Implement
}

func (userRepository *UserRepository) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	panic("not implemented") // TODO: Implement
}

func (userRepository *UserRepository) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	panic("not implemented") // TODO: Implement
}
