BEGIN;

-- Membuat tipe ENUM untuk Status
CREATE TYPE status_type AS ENUM ('paid', 'unpaid');

-- Membuat tabel transactions
CREATE TABLE IF NOT EXISTS transactions (
    ID SERIAL PRIMARY KEY,
    UserID INT NOT NULL,
    EventID INT NOT NULL,
    Status status_type NOT NULL,
    Quantity INT NOT NULL,
    Total INT NOT NULL,
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

CREATE TRIGGER update_transactions_updated_at
    BEFORE UPDATE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

COMMIT;
