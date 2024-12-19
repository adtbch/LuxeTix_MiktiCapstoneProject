BEGIN;
-- Menghapus tipe ENUM jika ada
DROP TYPE IF EXISTS status_type;

-- Menghapus trigger jika ada
DROP TRIGGER IF EXISTS update_transactions_updated_at ON transactions;

-- Menghapus fungsi jika ada
DROP FUNCTION IF EXISTS update_updated_at;

-- Menghapus tabel transactions jika ada
DROP TABLE IF EXISTS transactions;


COMMIT;
