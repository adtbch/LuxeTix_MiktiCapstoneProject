BEGIN;

-- Membuat tabel events (menggunakan VARCHAR untuk StatusEvent dan StatusRequest)
CREATE TABLE IF NOT EXISTS "public"."events" (
    ID SERIAL PRIMARY KEY,
    UserID INT NOT NULL,
    Title VARCHAR(255) NOT NULL UNIQUE,
    Description VARCHAR(255) NOT NULL,
    Date DATE NOT NULL,
    Price INT NOT NULL,
    Quantity INT NOT NULL,
    Time TIME NOT NULL,
    Location VARCHAR(255) NOT NULL,
    StatusEvent VARCHAR(50) NOT NULL,    -- Mengganti status_event_type ENUM dengan VARCHAR
    StatusRequest VARCHAR(50) NOT NULL,  -- Mengganti status_request_type ENUM dengan VARCHAR
    Category VARCHAR(255) NOT NULL,
    Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

COMMIT;
