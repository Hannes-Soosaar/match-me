package models

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

//! Below is some information for quick access.

// CREATE TABLE IF NOT EXISTS categories(
//     id SERIAL PRIMARY KEY,
//     category VARCHAR(255) NOT NULL
// );

// const (
// 	GENRE         = 1
// 	PLAY_STYLE    = 2
// 	PLATFORM      = 3
// 	COMMUNICATION = 4
// 	GOALS         = 5
// 	SESSION       = 6
// 	VIBE          = 7
// 	LANGUAGE      = 8
// )
