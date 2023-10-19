CREATE TABLE IF NOT EXISTS subcategories (
	uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	/*----------------------------*/
    -- extras
    title varchar,
    category_uuid uuid REFERENCES categories(uuid)
);

CREATE UNIQUE INDEX unique_subcategories_uuid on categories(uuid);
CREATE UNIQUE INDEX unique_subcategories_title on categories(title);
