CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "refresh_tokens"(
    token varchar
);

CREATE TABLE IF NOT EXISTS "roles"(
    uuid uuid not null DEFAULT uuid_generate_v4() PRIMARY KEY,
    name varchar
);

CREATE UNIQUE INDEX unique_name on roles(name);

INSERT INTO roles(uuid, name) VALUES ('bd9ea385-6e2c-4baa-8ac7-d9d832bc09da','admin');
INSERT INTO roles(uuid, name) VALUES ('9c8e2b1e-b839-4e49-937e-d460065eccb6','user');