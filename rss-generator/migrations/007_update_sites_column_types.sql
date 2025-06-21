-- +goose Up
-- +goose StatementBegin
ALTER TABLE sites
    MODIFY COLUMN url TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sites
    MODIFY COLUMN url VARCHAR(2047) NOT NULL;
-- +goose StatementEnd
