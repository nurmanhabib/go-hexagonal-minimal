-- +goose Up
CREATE TABLE users (
                       id BIGINT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE users;
