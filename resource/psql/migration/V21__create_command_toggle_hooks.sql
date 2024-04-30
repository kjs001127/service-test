CREATE TABLE command_toggle_hooks (
    app_id CHARACTER VARYING PRIMARY KEY,
    toggle_function_name CHARACTER VARYING NOT NULL
);

INSERT INTO command_toggle_hooks (app_id, toggle_function_name) SELECT app_id, toggle_function_name FROM command_activation_settings;

ALTER TABLE command_activation_settings DROP COLUMN toggle_function_name;