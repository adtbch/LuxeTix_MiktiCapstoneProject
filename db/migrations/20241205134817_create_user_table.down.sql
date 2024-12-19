BEGIN;
DROP TYPE IF EXISTS role_type;
DROP TYPE IF EXISTS gender_type;

-- Menghapus trigger jika ada
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Menghapus fungsi jika ada
DROP FUNCTION IF EXISTS update_updated_at;

-- Menghapus tabel users jika ada
DROP TABLE IF EXISTS users;

-- Menghapus tipe ENUM jika ada

COMMIT;
