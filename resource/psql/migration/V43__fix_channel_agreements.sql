ALTER TABLE channel_agreements DROP COLUMN app_role_id;
ALTER TABLE channel_agreements ADD COLUMN app_role_id INTEGER;
ALTER TABLE channel_agreements ADD CONSTRAINT channel_agreements_pkey PRIMARY KEY (channel_id, app_role_id);