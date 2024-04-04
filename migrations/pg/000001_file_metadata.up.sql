CREATE SCHEMA IF NOT EXISTS files;

-- todo: consider creating a many-to-may table to store file meta tags
-- CREATE TABLE IF NOT EXISTS files.file_metadata_tags
-- (
--     "id" SERIAL PRIMARY KEY,
--     "name" TEXT NOT NULL,
--     "crate_at" TIMESTAMPTZ DEFAULT timezone('UTC', CURRENT_TIMESTAMP),
--     "update_at" TIMESTAMPTZ
-- );

CREATE TABLE IF NOT EXISTS files.files_metadata
(
    "id" SERIAL PRIMARY KEY,
    "size" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "ext" TEXT NOT NULL,
    "content_type" TEXT NOT NULL,
    "hash_sum" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "created_by" TEXT NOT NULL,
    "last_modified_by" TEXT NOT NULL,
    "creating_date" TIMESTAMPTZ DEFAULT timezone('UTC', CURRENT_TIMESTAMP),
    "last_modified_time" TIMESTAMPTZ
);