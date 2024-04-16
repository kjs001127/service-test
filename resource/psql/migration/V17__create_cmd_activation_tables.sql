CREATE TABLE command_activation
(
    app_id     CHARACTER VARYING NOT NULL,
    channel_id CHARACTER VARYING NOT NULL,
    enabled    BOOLEAN           NOT NULL DEFAULT (TRUE),
    PRIMARY KEY (channel_id, app_id)
);


CREATE TABLE command_activation_setting
(
    app_id               CHARACTER VARYING PRIMARY KEY,
    enabled_by_default   BOOLEAN NOT NULL DEFAULT (TRUE),
    toggle_function_name CHARACTER VARYING
);