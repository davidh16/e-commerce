CREATE TABLE IF NOT EXISTS users (
     uuid uuid not null DEFAULT uuid_generate_v4() PRIMARY KEY,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    /*----------------------------*/
    -- extras
     email_address varchar unique,
     password varchar,
     is_active boolean,
     role_uuid uuid REFERENCES roles(uuid),
     first_name varchar,
     last_name varchar
);

CREATE UNIQUE INDEX unique_uuid on users(uuid);

CREATE UNIQUE INDEX unique_email on users(email_address);

INSERT INTO users(uuid, created_at, email_address, password, is_active, role_uuid) VALUES ('e92f5139-284a-46f4-b3b9-1002196f7792', current_timestamp, 'admin@heyclothing.com', '$2a$10$uYI1Gcq0HdNa7WPlBqivceawk82G6vg1em8cWjbSrSYqneNMb680e', true, 'bd9ea385-6e2c-4baa-8ac7-d9d832bc09da')