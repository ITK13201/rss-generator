-- +goose Up
-- +goose StatementBegin
CREATE TABLE test_feed_items
(
    id           INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    test_feed_id      CHAR(36)               NOT NULL,
    title        VARCHAR(1023)          NOT NULL,
    description  VARCHAR(2047)          NOT NULL,
    link         VARCHAR(2047),
    published_at DATETIME               NOT NULL,
    created_at   DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_test_feed_items_test_feed_id
        FOREIGN KEY (test_feed_id)
            REFERENCES test_feeds (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS test_feed_items;
-- +goose StatementEnd
