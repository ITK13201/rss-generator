-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds 
    MODIFY COLUMN title TEXT NOT NULL,
    MODIFY COLUMN description TEXT NOT NULL,
    MODIFY COLUMN link TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feeds 
    MODIFY COLUMN title VARCHAR(1023) NOT NULL,
    MODIFY COLUMN description VARCHAR(2047) NOT NULL,
    MODIFY COLUMN link VARCHAR(2047) NOT NULL;
-- +goose StatementEnd
