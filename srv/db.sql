create extension if not exists "uuid-ossp";
create extension if not exists citext;

CREATE TABLE if not exists authors
(
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    bio  text
);