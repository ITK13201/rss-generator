-- +goose Up
-- +goose StatementBegin
ALTER TABLE feed_items 
    MODIFY COLUMN title TEXT NOT NULL,
    MODIFY COLUMN description TEXT NOT NULL,
    MODIFY COLUMN link TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feed_items 
    MODIFY COLUMN title VARCHAR(1023) NOT NULL,
    MODIFY COLUMN description VARCHAR(2047) NOT NULL,
    MODIFY COLUMN link VARCHAR(2047);
-- +goose StatementEnd
