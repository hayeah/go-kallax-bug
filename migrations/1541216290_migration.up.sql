BEGIN;

ALTER TABLE test_models ADD COLUMN data2 bytea NOT NULL;

COMMIT;