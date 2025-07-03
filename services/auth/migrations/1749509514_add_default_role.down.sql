-- +migrate Dowr
DELETE FROM roles WHERE name = 'user';
