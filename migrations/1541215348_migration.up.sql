BEGIN;

CREATE TABLE test_models (
	id serial NOT NULL PRIMARY KEY,
	data bytea NOT NULL,
	counter bigint NOT NULL
);


COMMIT;
