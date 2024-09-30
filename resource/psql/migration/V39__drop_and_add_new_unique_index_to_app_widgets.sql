DROP INDEX IF EXISTS unique_index_app_widgets_on_app_id_and_name;
CREATE UNIQUE INDEX unique_index_app_widgets_on_app_id_and_scope_and_name ON app_widgets USING btree (app_id, scope, name);
