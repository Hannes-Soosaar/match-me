CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- there should be a way to delete a user with out deleting the user's history.
--OK
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    uuid UUID UNIQUE NOT NULL, 
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_city VARCHAR(50),
    user_nation VARCHAR(50),
    user_region VARCHAR(50),
    latitude FLOAT,
    longitude FLOAT,
	register_location GEOGRAPHY(POINT, 4326), 
	browser_location GEOGRAPHY(POINT, 4326)
);

-- OK
CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    uuid UUID UNIQUE NOT NULL,
    username VARCHAR(20) UNIQUE,
    about_me TEXT,
    profile_picture TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    birthdate DATE
);
-- OK not used!
CREATE TABLE IF NOT EXISTS bio_points (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    question VARCHAR(50) NOT NULL,
    answer TEXT NOT NULL
);
-- OK
CREATE TABLE IF NOT EXISTS sessions (
	id SERIAL PRIMARY KEY,
	session_guid UUID UNIQUE,
	user_id UUID
);


CREATE TABLE Messages (
    MessageID INT PRIMARY KEY,
    SenderID UUID NOT NULL,
    ReceiverID UUID NOT NULL,
    MessageText TEXT NOT NULL,
    SentAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    IsRead BOOLEAN DEFAULT FALSE
);

-- OK 
CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    category VARCHAR(255) NOT NULL
);
-- OK
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
-- OK
CREATE TABLE IF NOT EXISTS interests (
    id SERIAL PRIMARY KEY,
    categoryID VARCHAR(255) NOT NULL,
    interest VARCHAR(255) NOT NULL
);

-- OK
INSERT INTO interests(categoryId,interest)
VALUES 
(1,'RPG'),
(1,'FPS'),
(1,'MMO'),
(1,'MOBA'),
(1,'Turn Based'),
(1,'Simulation'),
(1,'A-RPG'),
(1,'Survival'),
(1,'PVP'),
(1,'PVE'),
(2,'Casual'),
(2,'Competitive'),
(2,'Relaxed'),
(2,'AFK'),
(2,'Enthusiast'),
(2,'Leeroy Jenkins'),
(3,'X-box'),
(3,'Switch'),
(3,'PC'),
(3,'Playstation'),
(3,'Mobile'),
(4,'voice chat '),
(4,'text chat '),
(4,'in-game chat'),
(4,'Discord'),
(4,'What ever, works'),
(5,'Socialize'),
(5,'Ranking'),
(5,'Challenge'),
(5,'Hang-out'),
(5,'For laughs'),
(6,'I got an hour to play'),
(6,'I need at least a few hours.'),
(6,'I can go all day, every day.'),
(7,'Competitive'),
(7,'Chill'),
(7,'High-energy'),
(7,'Laid-back'),
(7,'What-ever'),
(7,'Banter ahead'),
(7,'Black humor'),
(8,'Estonian'),
(8,'English'),
(8,'Texan'),
(8,'German'),
(8,'French'),
(8,'Russian'),
(8,'Chinese')
;

-- if there is a need to do time zone management we should use TIMESTAMPTZ
-- OK 
CREATE TABLE IF NOT EXISTS user_matches(
    id SERIAL PRIMARY KEY,
    user_id_1 UUID NOT NULL,
    user_id_2 UUID NOT NULL,
    match_score INTEGER,
    status VARCHAR(20),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    distance FLOAT
);


--OK
CREATE TABLE IF NOT EXISTS user_interests (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    interest_id INTEGER,
    status  VARCHAR(20)
);