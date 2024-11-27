package models

// The friend status is gotten from this structure, also blocks or deletes.
// We get the friends list by

type UserConnections struct {
	ID        int     `json:"id"`
	UserID1   int     `json:"userId1"`
	UserID2   int     `json:"userId2"`
	Status    string  `json:"status"`
	CreatedAt float64 `json:"createdAt"`
}
