-- +goose Up
-- +goose StatementBegin
CREATE TABLE feed_items
(
    id           INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    feed_id      CHAR(36)               NOT NULL,
    title        VARCHAR(1023)          NOT NULL,
    description  VARCHAR(2047)          NOT NULL,
    link         VARCHAR(2047),
    published_at DATETIME               NOT NULL,
    created_at   DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_feed_items_feed_id
        FOREIGN KEY (feed_id)
            REFERENCES feeds (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feed_items;
-- +goose StatementEnd
