CREATE TABLE app_secrets
(
    app_id CHARACTER VARYING PRIMARY KEY,
    secret CHARACTER VARYING NOT NULL
);

CREATE UNIQUE INDEX unique_index_app_tokens_on_token ON app_secrets USING btree (secret)