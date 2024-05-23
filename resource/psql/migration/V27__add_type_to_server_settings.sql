ALTER TABLE app_server_settings ADD COLUMN access_type CHARACTER VARYING NOT NULL DEFAULT ('external');
UPDATE app_server_settings SET access_type = 'internal'
