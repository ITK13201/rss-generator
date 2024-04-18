-- +goose Up
-- +goose StatementBegin
CREATE TABLE sites
(
    id         BINARY(16)             NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
    created_at DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sites;
-- +goose StatementEnd