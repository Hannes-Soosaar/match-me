CREATE TABLE IF NOT EXISTS unread_messages (
    id SERIAL PRIMARY KEY,
    match_id INTEGER UNIQUE NOT NULL,
    latest_message TIMESTAMP DEFAULT NULL,
    is_unread BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (match_id) REFERENCES user_matches(id)
);