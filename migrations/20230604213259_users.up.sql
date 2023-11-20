CREATE TABLE IF NOT EXISTS users (
     uuid uuid not null DEFAULT uuid_generate_v4() PRIMARY KEY,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    /*----------------------------*/
    -- extras
     email_address varchar unique,
     password varchar,
     is_active boolean,
     role_uuid uuid REFERENCES roles(uuid)
);

CREATE UNIQUE INDEX unique_uuid on users(uuid);

CREATE UNIQUE INDEX unique_email on users(email_address);