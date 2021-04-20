-- its for sqlite3
CREATE TABLE publ
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    mineatura TEXT NOT NULL,
    body TEXT NOT NULL
);
