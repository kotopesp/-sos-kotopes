ALTER TABLE IF EXISTS seekers
    DROP COLUMN IF EXISTS location,
    DROP COLUMN IF EXISTS equipment
    DROP COLUMN IF EXISTS price,
    DROP COLUMN IF EXISTS have_car,
    DROP COLUMN IF EXISTS is_deleted,
    DROP COLUMN IF EXISTS deleted_at;
