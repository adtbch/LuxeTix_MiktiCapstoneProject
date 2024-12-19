BEGIN;

-- Membuat tipe ENUM untuk StatusEvent dan StatusRequest
CREATE TYPE status_event_type AS ENUM ('available', 'unavailable');
CREATE TYPE status_request_type AS ENUM ('pending', 'unpaid', 'accepted', 'rejected');

-- Membuat tabel events
CREATE TABLE IF NOT EXISTS events (
    ID SERIAL PRIMARY KEY,
    UserID INT NOT NULL,
    Title VARCHAR(255) NOT NULL UNIQUE,
    Description VARCHAR(255) NOT NULL,
    Date DATE NOT NULL,
    Price INT NOT NULL,
    Quantity INT NOT NULL,
    Time TIME NOT NULL,
    Location VARCHAR(255) NOT NULL,
    StatusEvent status_event_type NOT NULL,
    StatusRequest status_request_type NOT NULL,
    Category VARCHAR(255) NOT NULL,
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

CREATE TRIGGER update_events_updated_at
    BEFORE UPDATE ON events
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

COMMIT;
