-- +goose Up
CREATE TABLE posts (
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(50) UNIQUE NOT NULL,
    url VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    feed_id INT NOT NULL,
    CONSTRAINT fk_feed FOREIGN KEY (feed_id)
    REFERENCES feeds(id)
    ON DELETE CASCADE
);