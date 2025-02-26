-- +goose Up
CREATE TABLE users(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;