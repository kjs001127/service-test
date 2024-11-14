ALTER TABLE app_roles
    ADD COLUMN version INTEGER NOT NULL DEFAULT (1);
DROP INDEX unique_index_app_roles_on_app_id_and_type;
CREATE UNIQUE INDEX unique_index_app_roles_on_app_id_and_type_and_version ON app_roles USING btree (app_id, type, version);

ALTER TABLE app_roles
    ADD COLUMN id SERIAL;
ALTER TABLE app_roles
    DROP CONSTRAINT app_roles_pkey;
ALTER TABLE app_roles
    ADD CONSTRAINT app_roles_pkey PRIMARY KEY (id);


CREATE TABLE channel_agreements
(
    channel_id  CHARACTER VARYING NOT NULL,
    app_role_id CHARACTER VARYING NOT NULL,
    PRIMARY KEY (channel_id, app_role_id)
);