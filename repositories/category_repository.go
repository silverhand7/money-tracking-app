package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/domain"
)

type CategoryRepository struct {
}

func (repository *CategoryRepository) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories (type, name, icon, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	var id int32
	err := tx.QueryRowContext(
		ctx,
		SQL,
		category.Type,
		category.Name,
		category.Icon,
		category.CreatedAt,
		category.UpdatedAt,
	).Scan(&id)

	helpers.PanicIfError(err)

	category.ID = int32(id)
	return category
}

func (repository *CategoryRepository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := `UPDATE categories SET
	type = $2,
	name = $3,
	icon = $4,
	updated_at = $5
	WHERE id = $1
	RETURNING id, type, name, icon, created_at, updated_at`

	row := tx.QueryRowContext(
		ctx,
		SQL,
		category.ID,
		category.Type,
		category.Name,
		category.Icon,
		category.UpdatedAt,
	)

	err := row.Scan(
		&category.ID,
		&category.Type,
		&category.Name,
		&category.Icon,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	helpers.PanicIfError(err)

	return category
}

func (repository *CategoryRepository) Delete(ctx context.Context, tx *sql.Tx, categoryId int32) {
	SQL := `DELETE FROM categories WHERE id=$1`
	_, err := tx.ExecContext(ctx, SQL, categoryId)
	helpers.PanicIfError(err)
}

func (repository *CategoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int32) (domain.Category, error) {
	SQL := "SELECT * FROM categories WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helpers.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(
			&category.ID,
			&category.Type,
			&category.Name,
			&category.Icon,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		helpers.PanicIfError(err)
		return category, nil
	}
	return category, errors.New("category is not found")
}

func (repository *CategoryRepository) GetAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, type, name, icon, created_at, updated_at FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Type,
			&category.Name,
			&category.Icon,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		helpers.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
