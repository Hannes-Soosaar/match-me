
SELECT * FROM users WHERE id = 1;

CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    category VARCHAR(255) NOT NULL
);

INSERT INTO categories (category)
VALUES 
('Genre'), -- 1
('Play Style'), --2
('Platform'), --3
('Communication'), --4
('Goals'), --5
('Session'), --6
('Vibe'), --7
('Language') --8
;

CREATE TABLE IF NOT EXISTS user_matches(
    id SERIAL PRIMARY KEY,
    user_id_1 INTEGER,
    user_id_2 INTEGER,
    match_score INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE user_interests (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    interest_id INTEGER,
    status  VARCHAR(20)
);
