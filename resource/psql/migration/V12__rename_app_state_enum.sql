UPDATE apps SET state = 'enabled' where state = 'stable';
UPDATE apps SET state = 'disabled' where state = 'unstable';
