CREATE TABLE IF NOT EXISTS "user"
(
    id         uuid                        DEFAULT gen_random_uuid() UNIQUE PRIMARY KEY,
    first_name VARCHAR(255)                NOT NULL,
    last_name  VARCHAR(255)                NOT NULL,
    email      VARCHAR(255)                NOT NULL UNIQUE,
    password   VARCHAR(255)                NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

-- email idx
-- CREATE UNIQUE INDEX CONCURRENTLY user_email
-- ON "user" (email);

-- needful extension  password encryption
CREATE EXTENSION IF NOT EXISTS pgcrypto;