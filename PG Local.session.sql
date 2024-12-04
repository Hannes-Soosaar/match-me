-- INSERT INTO profiles (
--         user_id,
--         username,
--         about_me,
--         profile_picture
--     )
-- VALUES (
--         1,
--         'Peeter123',
--         'I am a gamer',
--         'picture path'
--     );

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