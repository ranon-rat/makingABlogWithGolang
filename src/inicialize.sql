-- this is for postgresql
CREATE DATABASE publications
    WITH 
    OWNER = ranon;

CREATE TABLE publ
(
    id SERIAL PRIMARY KEY,
    titulo TEXT NOT NULL,
    mineatura TEXT NOT NULL,
    body TEXT NOT NULL
);