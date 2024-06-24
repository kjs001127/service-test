ALTER TABLE commands
    ADD COLUMN enabled_by_default BOOL NOT NULL default true;

CREATE TABLE command_channel_activations
(
    command_id CHARACTER VARYING NOT NULL,
    channel_id CHARACTER VARYING NOT NULL,
    enabled    BOOL              NOT NULL,
    PRIMARY KEY (channel_id, command_id)
);

INSERT INTO command_channel_activations (command_id, channel_id, enabled)
SELECT DISTINCT commands.id, command_activations.channel_Id, command_activations.enabled
FROM command_activations
         JOIN commands ON commands.app_id = command_activations.app_id;

DROP TABLE command_activations;
DROP TABLE command_activation_settings;
