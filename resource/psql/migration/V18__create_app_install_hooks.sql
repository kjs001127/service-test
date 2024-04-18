CREATE TABLE app_install_hooks
(
    app_id                  CHARACTER VARYING PRIMARY KEY NOT NULL,
    install_function_name   CHARACTER VARYING,
    uninstall_function_name CHARACTER VARYING
);
