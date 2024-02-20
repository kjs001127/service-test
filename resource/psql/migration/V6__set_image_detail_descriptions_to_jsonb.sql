ALTER TABLE apps ALTER COLUMN detail_descriptions TYPE JSONB USING '[]'::JSONB;
