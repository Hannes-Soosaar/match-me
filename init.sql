CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    username VARCHAR(20) NOT NULL UNIQUE,
    about_me TEXT,
    profile_picture TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE bio_points (
    id SERIAL PRIMARY KEY,
    profile_id INT NOT NULL REFERENCES profiles(id) ON DELETE CASCADE,
    question VARCHAR(50) NOT NULL,
    answer TEXT NOT NULL
);
CREATE TABLE connections (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    connected_user_id INT NOT NULL REFERENCES USERS(ID) ON DELETE CASCADE
);
/*INSERT INTO users (email, password_hash)
 VALUES ('test1@example.com', '$2b$10$hashfortest1'),
 ('test2@example.com', '$2b$10$hashfortest2');*/
/*INSERT INTO profiles (
 user_id,
 username,
 about_me
 )
 VALUES (
 1,
 'Peeter123',
 'I am a gamer'
 );/*