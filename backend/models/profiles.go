package models

import "time"

//TODO: add correct table structure.

type Profiles struct {
	ID        int       `json:"id"`
	// UserID1   int       `json:"userId1"`
	// UserID2   int       `json:"userId2"`
	// Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}


// CREATE TABLE profiles (
//     id SERIAL PRIMARY KEY,
//     user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,  //If there are no foreign key constraints it will not cascade 
//     username VARCHAR(20) NOT NULL UNIQUE,
//     about_me TEXT,
//     profile_picture TEXT,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );