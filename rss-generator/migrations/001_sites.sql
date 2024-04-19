-- +goose Up
-- +goose StatementBegin
CREATE TABLE sites
(
    id          INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    slug        VARCHAR(65)            NOT NULL UNIQUE,
    title       VARCHAR(127)           NOT NULL,
    description VARCHAR(511),
    url         VARCHAR(511)           NOT NULL,
    created_at  DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sites;
-- +goose StatementEnd
