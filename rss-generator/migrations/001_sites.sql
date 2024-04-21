-- +goose Up
-- +goose StatementBegin
CREATE TABLE sites
(
    id                  INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    slug                VARCHAR(255)           NOT NULL UNIQUE,
    title               VARCHAR(1023)          NOT NULL,
    description         VARCHAR(2047),
    url                 VARCHAR(2047)          NOT NULL,
    enable_js_rendering BOOLEAN                NOT NULL DEFAULT FALSE,
    created_at          DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sites;
-- +goose StatementEnd
