CREATE TABLE IF NOT EXISTS categories (
	uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	/*----------------------------*/
    -- extras
    title varchar
);

CREATE UNIQUE INDEX unique_uuid on categories(uuid);