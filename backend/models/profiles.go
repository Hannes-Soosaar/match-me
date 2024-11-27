package models

import "time"

type Profiles struct {
	ID        int       `json:"id"`
	UserID1   int       `json:"userId1"`
	UserID2   int       `json:"userId2"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
