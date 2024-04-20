-- +goose Up
-- +goose StatementBegin
CREATE TABLE scraping_selectors
(
    id                   INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    site_id              INTEGER                NOT NULL UNIQUE,
    selector             VARCHAR(511)           NOT NULL,
    inner_selector       VARCHAR(511)           NOT NULL,
    title_selector       VARCHAR(511)           NOT NULL,
    description_selector VARCHAR(511),
    link_selector        VARCHAR(511),
    created_at           DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_scraping_selectors_site_id
        FOREIGN KEY (site_id)
        REFERENCES sites (id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS scraping_selectors;
-- +goose StatementEnd
