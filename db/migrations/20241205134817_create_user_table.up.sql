BEGIN;

-- Membuat tabel users (menggunakan VARCHAR untuk Role dan Gender)
CREATE TABLE IF NOT EXISTS "public"."users" (
    ID SERIAL PRIMARY KEY,
    Fullname VARCHAR(255) NOT NULL,
    Username VARCHAR(255) NOT NULL UNIQUE,
    Email VARCHAR(255) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    Role VARCHAR(50) NOT NULL,    -- Mengganti role_type ENUM dengan VARCHAR
    Gender VARCHAR(50) NOT NULL,  -- Mengganti gender_type ENUM dengan VARCHAR
    verify_token VARCHAR(255) NOT NULL,
    reset_token VARCHAR(255) NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Menambahkan default value
    Updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Menambahkan default value (tidak ada koma di sini)
);

COMMIT;
