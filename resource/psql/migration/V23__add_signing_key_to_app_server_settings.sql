ALTER TABLE app_urls ADD COLUMN signing_key CHARACTER VARYING;
ALTER TABLE app_urls RENAME TO app_server_settings;
