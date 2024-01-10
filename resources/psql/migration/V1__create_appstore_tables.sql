CREATE TABLE apps
(
    id            CHARACTER VARYING PRIMARY KEY NOT NULL,

    title         CHARACTER VARYING             NOT NULL,
    description   CHARACTER VARYING,
    avatar_url    CHARACTER VARYING,

    wam_uri       CHARACTER VARYING,
    rpc_uri       CHARACTER VARYING,
    install_uri   CHARACTER VARYING             NOT NULL,
    check_uri     CHARACTER VARYING,
    state         CHARACTER VARYING             NOT NULL,
    config_scheme JSONB,

    created_at    TIMESTAMP WITHOUT TIME ZONE   NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITHOUT TIME ZONE   NOT NULL DEFAULT NOW()
);


CREATE TABLE commands
(
    id                CHARACTER VARYING PRIMARY KEY          NOT NULL,
    app_id            CHARACTER VARYING REFERENCES apps (id) NOT NULL,
    function_name     CHARACTER VARYING                      NOT NULL,

    name              CHARACTER VARYING                      NOT NULL,
    scope             CHARACTER VARYING                      NOT NULL,
    description       CHARACTER VARYING,
    param_definitions JSONB,

    created_at        TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_index_commands_on_app_id_and_scope_and_name ON commands USING btree (app_id, scope, name);


CREATE TABLE wams
(
    id         CHARACTER VARYING PRIMARY KEY          NOT NULL,

    app_id     CHARACTER VARYING REFERENCES apps (id) NOT NULL,
    name       CHARACTER VARYING                      NOT NULL,

    created_at TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX unique_index_wams_on_app_id_and_name ON wams USING btree (app_id, name);


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


