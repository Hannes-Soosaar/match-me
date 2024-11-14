CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL
);
INSERT INTO users (username, email)
VALUES ('test1', 'test1@example.com'),
    ('test2', 'test2@example.com');