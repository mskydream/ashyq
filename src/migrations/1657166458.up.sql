BEGIN;
CREATE TABLE posts(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description TEXT NOT NULL,
    image_src varchar(255)
);
ALTER TABLE posts OWNER TO study;
COMMIT;
