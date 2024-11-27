package models

import "time"

// keeps track of all chats

type Chats struct {
	ID        int     `json:"id"`
	UserID1   int     `json:"userId1"`
	UserID2   int     `json:"userId2"`
	CreatedAt time.Time `json:"createdAt"`
}