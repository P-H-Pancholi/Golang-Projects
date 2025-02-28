-- +goose Up
CREATE TABLE users(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE feeds(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,
    url VARCHAR(50) UNIQUE NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

CREATE TABLE feed_follows(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id INT NOT NULL,
    feed_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_feed FOREIGN KEY (feed_id)
    REFERENCES feeds(id)
    ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
DROP TABLE feeds;
DROP TABLE users;