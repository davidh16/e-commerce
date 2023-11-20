CREATE TABLE IF NOT EXISTS tokens (
	uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	/*----------------------------*/
    -- extras
    token varchar,
    user_uuid varchar,
    is_used boolean,
    expires_at TIMESTAMP
);

CREATE UNIQUE INDEX unique_tokens_uuid on tokens(uuid);
