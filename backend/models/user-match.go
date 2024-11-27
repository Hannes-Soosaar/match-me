package models

import "time"

// There will be a function to calculate a score for two users.
// how to page
// how many scores should we calculate
// when should the score be calculated

type UsersMatches struct {
	ID        int     `json:"id"`
	UserID1   int     `json:"userId1"`
	UserID2   int     `json:"userId2"`
	MatchScore     int  `json:"MatchScore"`
	CreatedAt time.Time `json:"createdAt"`
}
