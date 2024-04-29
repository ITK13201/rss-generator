-- +goose Up
-- +goose StatementBegin
CREATE TABLE scraping_settings
(
    id                   INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    site_id              INTEGER                NOT NULL UNIQUE,
    selector             VARCHAR(1023)          NOT NULL,
    inner_selector       VARCHAR(1023)          NOT NULL,
    title_selector       VARCHAR(1023)          NOT NULL,
    description_selector VARCHAR(1023),
    link_selector        VARCHAR(1023),
    created_at           DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_scraping_settings_site_id
        FOREIGN KEY (site_id)
            REFERENCES sites (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS scraping_settings;
-- +goose StatementEnd
