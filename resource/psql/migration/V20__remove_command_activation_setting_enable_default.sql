ALTER TABLE command_activation RENAME TO command_activations;
ALTER TABLE command_activation_setting RENAME TO command_activation_settings;

ALTER TABLE command_activation_settings ALTER COLUMN enabled_by_default DROP DEFAULT;