-- this is for postgresql
CREATE DATABASE publications
    WITH 
    OWNER = ranon;

CREATE TABLE publ
(
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(40) NOT NULL,
    mineatura VARCHAR(100) NOT NULL,
    body VARCHAR(8000) NOT NULL
);