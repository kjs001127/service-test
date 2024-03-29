UPDATE apps SET state = 'enabled' WHERE state = 'stable';
UPDATE apps SET state = 'disabled' WHERE state = 'unstable';
