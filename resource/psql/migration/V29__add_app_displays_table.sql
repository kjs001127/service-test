CREATE TABLE app_displays
(
    app_id              CHARACTER VARYING PRIMARY KEY,
    is_private          BOOLEAN NOT NULL,
    manual_url          CHARACTER VARYING,
    detail_descriptions JSONB,
    detail_image_urls   CHARACTER VARYING[],
    i18n_map            JSONB
);

INSERT INTO app_displays (app_id, is_private, manual_url, detail_descriptions, detail_image_urls, i18n_map)
SELECT id, is_private, manual_url, detail_descriptions, detail_image_urls, i18n_map FROM apps;

ALTER TABLE apps
    DROP COLUMN is_private,
    DROP COLUMN manual_url,
    DROP COLUMN detail_descriptions,
    DROP COLUMN detail_image_urls;
