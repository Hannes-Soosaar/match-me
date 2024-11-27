package models

import "time"

// There will be a function to calculate a score for two users.
// how to page
// how many scores should we calculate
// when should the score be calculated
// Compare interest  each interest in-common is 1 point 
// Compare multiple choice questions each some questions will give a point if its a match others if the answer is not a match
// Compare locations if not in specific range the score modified if x0  // only filter if is shown. ?? do we need to even do this in the scores tab?
// Compare Slider scores each match is 1 point each difference is  1 -0.25 points and if one answer is opposite the score is  there are 4 steps
//


type UsersMatches struct {
	ID        int     `json:"id"`
	UserID1   int     `json:"userId1"`
	UserID2   int     `json:"userId2"`
	MatchScore     int  `json:"MatchScore"`
	CreatedAt time.Time `json:"createdAt"`
}
