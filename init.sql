CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO users (email, password_hash)
VALUES (
        'test1@example.com',
        '$2b$10$hashfortest1'
    ),
    (
        'test2@example.com',
        '$2b$10$hashfortest2'
    );