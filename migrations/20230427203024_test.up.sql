CREATE TABLE IF NOT EXISTS test (
	uuid uuid DEFAULT uuid_generate_v4(),
	username VARCHAR,
	email VARCHAR,
	password VARCHAR
);
