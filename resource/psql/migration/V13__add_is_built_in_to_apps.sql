ALTER TABLE apps ADD COLUMN is_built_in BOOL DEFAULT false;

CREATE INDEX index_apps_on_is_built_in ON apps (is_built_in) WHERE apps.is_built_in IS TRUE;
