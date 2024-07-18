DELETE
FROM app_accounts
WHERE app_id in (SELECT acc.app_id
                 FROM app_accounts acc
                          LEFT JOIN apps a ON acc.app_id = a.id
                 WHERE a.id IS NULL);
