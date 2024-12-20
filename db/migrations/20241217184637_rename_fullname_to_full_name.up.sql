BEGIN;

ALTER TABLE users RENAME COLUMN fullname TO full_name;

COMMIT;