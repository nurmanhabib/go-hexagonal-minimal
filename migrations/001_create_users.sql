-- +goose Up
CREATE TABLE users (
                       id VARCHAR(36) PRIMARY KEY,
                       name VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE users;
