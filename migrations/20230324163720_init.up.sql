CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "refresh_tokens"(
    token varchar
);

CREATE TABLE IF NOT EXISTS "roles"(
    uuid uuid not null DEFAULT uuid_generate_v4() PRIMARY KEY,
    name varchar
);

CREATE UNIQUE INDEX unique_name on roles(name);