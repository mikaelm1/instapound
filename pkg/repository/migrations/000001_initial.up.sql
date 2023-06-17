CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;

CREATE TABLE IF NOT EXISTS "base_table" (
id uuid NOT NULL DEFAULT uuid_generate_v4 (),
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);


create table IF NOT EXISTS "user" (
active boolean NOT NULL DEFAULT true,
email varchar(255) UNIQUE NOT NULL,
role VARCHAR(255) NOT NULL,
password varchar(255) NOT NULL,
user_name varchar(255) UNIQUE NOT NULL,
PRIMARY KEY (id)
) INHERITS ("base_table");

CREATE INDEX idx_user_email
ON "user" (email);
CREATE INDEX idx_user_id
ON "user"(id);

CREATE TABLE IF NOT EXISTS "post" (
content VARCHAR(255) NOT NULL,
user_id uuid NOT NULL,
CONSTRAINT user_id
    FOREIGN KEY(user_id)
    REFERENCES "user"(id)
    ON DELETE CASCADE,
PRIMARY KEY (id)
) INHERITS ("base_table");

CREATE TABLE IF NOT EXISTS "post_likes" (
    post_id uuid NOT NULL,
    CONSTRAINT post_id
        FOREIGN KEY(post_id)
        REFERENCES post(id),
    user_id uuid NOT NULL,
    CONSTRAINT user_id
        FOREIGN KEY(user_id)
        REFERENCES "user"(id)
        ON DELETE CASCADE,
PRIMARY KEY (post_id)
) INHERITS ("base_table");

-- --
-- `being_followed_id` is the id of the user who is being followed by someone
-- --
CREATE TABLE IF NOT EXISTS "followers" (
    being_followed_id uuid NOT NULL,
    CONSTRAINT being_followed_id
        FOREIGN KEY(being_followed_id)
        REFERENCES "user"(id)
        ON DELETE CASCADE,
    following_id uuid NOT NULL,
    CONSTRAINT following_id
        FOREIGN KEY(following_id)
        REFERENCES "user"(id)
        ON DELETE CASCADE,
PRIMARY KEY (being_followed_id)
) INHERITS ("base_table");

COMMIT;