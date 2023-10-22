CREATE TABLE IF NOT EXISTS media (
	uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	/*----------------------------*/
    -- extras
    filename varchar,
    path varchar,
    size int
);

CREATE UNIQUE INDEX unique_media_uuid on media(uuid);
