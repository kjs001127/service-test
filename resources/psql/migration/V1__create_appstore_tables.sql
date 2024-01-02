CREATE TABLE apps
(
    id               CHARACTER VARYING PRIMARY KEY NOT NULL,
    secret           CHARACTER VARYING             NOT NULL,

    title            CHARACTER VARYING             NOT NULL,
    description      CHARACTER VARYING,
    icon             CHARACTER VARYING,

    host             CHARACTER VARYING,
    wam_uri          CHARACTER VARYING,
    command_uri      CHARACTER VARYING,
    function_uri     CHARACTER VARYING,
    installation_uri CHARACTER VARYING             NOT NULL,

    active           BOOLEAN                       NOT NULL DEFAULT FALSE,
    configs          JSONB,

    created_at       TIMESTAMP WITHOUT TIME ZONE   NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP WITHOUT TIME ZONE   NOT NULL DEFAULT NOW()
);


CREATE TABLE commands
(
    app_id      CHARACTER VARYING REFERENCES apps (id) NOT NULL,
    scope       CHARACTER VARYING                      NOT NULL,
    name        CHARACTER VARYING                      NOT NULL,

    title       CHARACTER VARYING                      NOT NULL,
    description CHARACTER VARYING,
    parameters  JSONB,

    created_at  TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),

    PRIMARY KEY (app_id, scope, name)
);

CREATE TABLE wams
(
    app_id     CHARACTER VARYING REFERENCES apps (id) NOT NULL,
    name       CHARACTER VARYING                      NOT NULL,

    created_at TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),

    PRIMARY KEY (app_id, name)
);

CREATE TABLE functions
(
    app_id     CHARACTER VARYING REFERENCES apps (id) NOT NULL,
    name       CHARACTER VARYING                      NOT NULL,

    title      CHARACTER VARYING                      NOT NULL,
    parameters JSONB,

    created_at TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE            NOT NULL DEFAULT NOW(),

    PRIMARY KEY (app_id, name)
);

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


