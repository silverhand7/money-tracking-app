-- +goose Up
CREATE TABLE wallets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255),
    currency VARCHAR(10) NOT NULL DEFAULT 'USD',
    balance BIGINT DEFAULT 0,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE wallets;