-- +goose Up

-- noop 2
CREATE SCHEMA extensions;
CREATE EXTENSION IF NOT EXISTS "pg_uuidv7" SCHEMA extensions;

-- +goose Down
DROP SCHEMA extensions;
