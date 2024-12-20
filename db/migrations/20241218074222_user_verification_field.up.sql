BEGIN;

ALTER TABLE users
    ADD COLUMN reset_password_token TEXT,
    ADD COLUMN verify_email_token TEXT,
    ADD COLUMN is_verified BOOLEAN NOT NULL DEFAULT FALSE;

COMMIT;
