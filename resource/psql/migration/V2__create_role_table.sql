CREATE TABLE app_roles
(
    app_id    CHARACTER VARYING NOT NULL REFERENCES apps(id),
    role_id   CHARACTER VARYING NOT NULL,
    client_Id CHARACTER VARYING NOT NULL,
    secret    CHARACTER VARYING NOT NULL,
    PRIMARY KEY (app_id, role_id)
);

