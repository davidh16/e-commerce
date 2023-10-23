CREATE TABLE IF NOT EXISTS addresses (
	uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	/*----------------------------*/
    -- extras
    street_address varchar,
    city varchar,
    country varchar,
    postal_code varchar,
    state varchar,
    user_uuid uuid REFERENCES users(uuid)
);

CREATE UNIQUE INDEX unique_addresses_uuid on addresses(uuid);