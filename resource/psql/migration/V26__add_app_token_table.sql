CREATE TABLE app_tokens (
    app_id CHARACTER VARYING PRIMARY KEY,
    token CHARACTER VARYING NOT NULL
);

CREATE UNIQUE INDEX unique_index_app_tokens_on_token ON app_tokens USING btree(token)