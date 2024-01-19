-- Database System: PostgreSQL
-- Database Name: Novels-API
create table genres(id serial primary key, name varchar);

create table authors(id serial primary key, name varchar);

create table novels(
    id serial primary key,
    title varchar,
    description varchar,
    language varchar,
    type varchar,
    genre_id int,
    author_id int,
    year int,
    CONSTRAINT fk_genre FOREIGN KEY(genre_id) REFERENCES genres(id),
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES authors(id)
);