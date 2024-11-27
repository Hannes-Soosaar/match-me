package models

type Chats struct {
	ID        int     `json:"id"`
	UserID1   int     `json:"userId1"`
	UserID2   int     `json:"userId2"`
	CreatedAt float64 `json:"createdAt"`
}