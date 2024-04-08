CREATE SCHEMA IF NOT EXISTS files;

CREATE TABLE IF NOT EXISTS files.metadata
(
    "id" SERIAL PRIMARY KEY,
    "size" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "ext" TEXT NOT NULL,
    "content_type" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "is_deleted" BOOLEAN DEFAULT FALSE,
    "creating_date" TIMESTAMPTZ DEFAULT timezone('UTC', CURRENT_TIMESTAMP),
    "last_modified_time" TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS files.idx_file_metadata_name ON files.metadata("name");