-- +goose Up
CREATE TABLE "user" (
    id integer PRIMARY KEY NOT NULL,
    first_name varchar(128) NOT NULL,
    last_name varchar(128) NOT NULL
);

CREATE TABLE user_mfa_shared_secret (
    user_id integer NOT NULL PRIMARY KEY REFERENCES "user" (id) ON DELETE CASCADE,
    shared_secret varchar(128) NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS user_mfa_shared_secret;
DROP TABLE IF EXISTS "user";