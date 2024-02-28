ALTER TABLE commands ADD COLUMN name_description_i18n_map JSONB;
ALTER TABLE commands DROP COLUMN description_i18n_map;
ALTER TABLE commands DROP COLUMN name_i18n_map;

