BEGIN;

-- Membuat tabel transactions (menggunakan VARCHAR untuk Status)
CREATE TABLE IF NOT EXISTS "public"."transactions" (
    ID SERIAL PRIMARY KEY,
    UserID INT NOT NULL,
    EventID INT NOT NULL,
    Status VARCHAR(50) NOT NULL,    -- Mengganti status_type ENUM dengan VARCHAR
    Quantity INT NOT NULL,
    Amount INT NOT NULL,
    Type VARCHAR(255) NOT NULL,
    Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


COMMIT;
