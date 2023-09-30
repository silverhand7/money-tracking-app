-- +goose Up
ALTER TABLE transactions ADD COLUMN note VARCHAR(255);

-- +goose Down
ALTER TABLE transactions DROP COLUMN note