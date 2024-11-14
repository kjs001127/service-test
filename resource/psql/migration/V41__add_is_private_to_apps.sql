ALTER TABLE apps ADD COLUMN is_private BOOLEAN NOT NULL DEFAULT TRUE;

UPDATE apps
SET is_private = app_displays.is_private
FROM app_displays
WHERE apps.id = app_displays.app_id;
