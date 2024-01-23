CREATE TABLE apps
(
    id                 CHARACTER VARYING PRIMARY KEY NOT NULL,

    client_id          CHARACTER VARYING             NOT NULL,
    secret             CHARACTER VARYING             NOT NULL,

    role_id             CHARACTER VARYING            NOT NULL,

    title              CHARACTER VARYING             NOT NULL,
    description        CHARACTER VARYING,
    detail_description JSONB,
    detail_image_urls  CHARACTER VARYING,
    avatar_url         CHARACTER VARYING,

    wam_url            CHARACTER VARYING,
    function_url       CHARACTER VARYING,
    hook_url           CHARACTER VARYING,
    check_url          CHARACTER VARYING,
    manual_url         CHARACTER VARYING,

    state              CHARACTER VARYING             NOT NULL,
    config_schema      JSONB,

    created_at         TIMESTAMP WITHOUT TIME ZONE   NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP WITHOUT TIME ZONE   NOT NULL DEFAULT NOW()
);


CREATE TABLE commands
(
    id                         CHARACTER VARYING PRIMARY KEY          NOT NULL,

    app_id                     CHARACTER VARYING REFERENCES apps (id) NOT NULL,
    name                       CHARACTER VARYING                      NOT NULL,
    scope                      CHARACTER VARYING                      NOT NULL,

    function_name              CHARACTER VARYING                      NOT NULL,
    autocomplete_function_name CHARACTER VARYING,
    param_definitions          JSONB                                  NOT NULL,

    description                CHARACTER VARYING,

    created_at                 TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),
    updated_at                 TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_index_commands_on_app_id_and_scope_and_name ON commands USING btree (app_id, scope, name);


CREATE TABLE app_channels
(
    app_id     CHARACTER VARYING REFERENCES apps (id) NOT NULL,
    channel_id CHARACTER VARYING                      NOT NULL,

    active     BOOLEAN                                NOT NULL DEFAULT FALSE,
    configs    JSONB,

    created_at TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),

    PRIMARY KEY (app_id, channel_id)
);


