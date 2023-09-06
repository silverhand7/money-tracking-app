-- +goose Up
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    "type" VARCHAR(1) NOT NULL CHECK ("type" IN ('I', 'E')), --I = income, E = expense
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE categories;