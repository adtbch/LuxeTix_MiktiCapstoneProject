BEGIN;

DROP TYPE IF EXISTS status_event_type;
DROP TYPE IF EXISTS status_request_type;
-- Menghapus trigger jika ada
DROP TRIGGER IF EXISTS update_events_updated_at ON events;

-- Menghapus fungsi jika ada
DROP FUNCTION IF EXISTS update_updated_at;

-- Menghapus tabel events jika ada
DROP TABLE IF EXISTS events;

-- Menghapus tipe ENUM jika ada

COMMIT;
