CREATE TABLE IF NOT EXISTS "user" --
(
    id         uuid                        NOT NULL DEFAULT gen_random_uuid() UNIQUE PRIMARY KEY,
    full_name  VARCHAR(256)                NOT NULL,
    email      VARCHAR(256)                NOT NULL UNIQUE,
    password   VARCHAR(256)                NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

-- needful extension  password encryption
CREATE EXTENSION IF NOT EXISTS pgcrypto;