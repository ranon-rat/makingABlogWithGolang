-- this is for postgresql
CREATE DATABASE publications
    WITH 
    OWNER = ranon;

CREATE TABLE publ
(
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(40),
    mineatura VARCHAR(100),
    body VARCHAR(8000)
);