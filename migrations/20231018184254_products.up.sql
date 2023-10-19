CREATE TABLE IF NOT EXISTS products (
	uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	/*----------------------------*/
    -- extras
    name varchar,
    brand varchar,
    description text,
    price float,
    image_url varchar,
    color varchar,
    code varchar,
    subcategory_uuid uuid REFERENCES subcategories(uuid),
    category_uuid uuid REFERENCES categories(uuid)
);

CREATE UNIQUE INDEX unique_products_uuid on products(uuid);
CREATE UNIQUE INDEX unique_products_name on products(name);
CREATE UNIQUE INDEX unique_products_code on products(code);

