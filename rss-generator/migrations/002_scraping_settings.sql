-- +goose Up
-- +goose StatementBegin
CREATE TABLE scraping_settings
(
    id                   INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    site_id              INTEGER                NOT NULL UNIQUE,
    selector             TEXT                   NOT NULL,
    inner_selector       TEXT                   NOT NULL,
    title_selector       TEXT                   NOT NULL,
    description_selector TEXT,
    link_selector        TEXT,
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
