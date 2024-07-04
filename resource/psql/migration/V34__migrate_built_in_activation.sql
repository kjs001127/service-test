INSERT INTO command_channel_activations (command_id, channel_id, enabled)
SELECT DISTINCT commands.id, app_installations.channel_id, true
FROM apps
         JOIN commands ON apps.id = commands.app_id
         JOIN app_installations ON app_installations.app_id = apps.id
WHERE apps.is_built_in = true ON CONFLICT (command_id, channel_id) DO NOTHING;
