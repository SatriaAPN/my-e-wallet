-- init.sql
CREATE TABLE mytable (
    id SERIAL PRIMARY KEY,
    name TEXT
);

INSERT INTO mytable (name) VALUES ('Alice'), ('Bob'), ('Charlie');
