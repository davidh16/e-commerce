CREATE TABLE IF NOT EXISTS users (
     uuid uuid not null DEFAULT uuid_generate_v4() PRIMARY KEY ,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    /*----------------------------*/
    -- extras
     email_address varchar unique,
     password varchar,
     account_status int not null ,
     shipping_address uuid,
     payment_info uuid
);

CREATE UNIQUE INDEX unique_uuid on users(uuid);

CREATE UNIQUE INDEX unique_email on users(email_address);