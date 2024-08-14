CREATE TABLE app_widgets
(
    id                         CHARACTER VARYING PRIMARY KEY,
    app_id                     CHARACTER VARYING NOT NULL,

    default_name               CHARACTER VARYING,
    default_description        CHARACTER VARYING,
    default_name_desc_i18n_map JSONB,

    name                       CHARACTER VARYING NOT NULL,
    description                CHARACTER VARYING,
    name_desc_i18n_map         JSONB,

    action_function_name       CHARACTER VARYING NOT NULL
);

CREATE UNIQUE INDEX unique_index_app_widgets_on_app_id_and_name ON app_widgets USING btree (app_id, name)
