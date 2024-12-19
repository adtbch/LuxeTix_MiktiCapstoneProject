BEGIN;

-- Membuat tipe ENUM untuk Role dan Gender
CREATE TYPE role_type AS ENUM ('Admin', 'User');
CREATE TYPE gender_type AS ENUM ('Male', 'Female');

-- Membuat tabel users
CREATE TABLE IF NOT EXISTS users (
    ID SERIAL PRIMARY KEY,
    Fullname VARCHAR(255) NOT NULL,
    Username VARCHAR(255) NOT NULL UNIQUE,
    Email VARCHAR(255) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    Role role_type NOT NULL,
    Gender gender_type NOT NULL,
    verify_token VARCHAR(255) NOT NULL,
    reset_token VARCHAR(255) NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Membuat trigger untuk memperbarui kolom Updated_at saat ada perubahan
CREATE OR REPLACE FUNCTION update_updated_at() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.Updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

COMMIT;
