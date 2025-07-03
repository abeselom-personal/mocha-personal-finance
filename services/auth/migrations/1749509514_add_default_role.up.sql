-- +migrate Up
INSERT INTO roles (name) 
SELECT 'user'
WHERE NOT EXISTS (
  SELECT 1 FROM roles WHERE name = 'user'
);
