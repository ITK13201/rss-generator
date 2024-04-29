-- +goose Up
-- +goose StatementBegin
CREATE TABLE test_feeds
(
    id           CHAR(36)      NOT NULL DEFAULT (UUID()) PRIMARY KEY,
    site_id      INTEGER       NOT NULL,
    title        VARCHAR(1023) NOT NULL,
    description  VARCHAR(2047) NOT NULL,
    link         VARCHAR(2047) NOT NULL,
    published_at DATETIME      NOT NULL,
    created_at   DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_test_feeds_site_id
        FOREIGN KEY (site_id)
            REFERENCES sites (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS test_feeds;
-- +goose StatementEnd
