CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	uuid uuid DEFAULT uuid_generate_v4(),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	----------------------------
    -- extras

);
