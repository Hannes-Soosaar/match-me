package db

import (
	"fmt"
	"match_me_backend/models"
)

//TODO:

func GetAllUserMatches() ([]models.UsersMatches, error) {
	query := "SELECT id, user_id_1, user_id_2, match_score, created_at FROM user_matches"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()
	var userMatches []models.UsersMatches
	for rows.Next() {
		var userMatch models.UsersMatches
		err = rows.Scan(&userMatch.ID, &userMatch.UserID1, &userMatch.UserID2, &userMatch.MatchScore, &userMatch.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		userMatches = append(userMatches, userMatch)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iterations: %w", err)
	}
	return userMatches, nil
}


func AddUserMatch(userID1, userID2 string) error {
query := `
	INSERT INTO user_matches (user_id_1, user_id_2, match_score, created_at)
	SELECT $1, $2, $3, now()
	WHERE NOT EXISTS (
		SELECT 1 
		FROM user_matches 
		WHERE user_id_1 = $1 AND user_id_2 = $2
	);`
	_, err := DB.Exec(query, userID1, userID2, 0)
	if err != nil {
		return fmt.Errorf("error adding user match: %w", err)
	}
	return nil
}


// User IDs are passed in as uuid strings
func UpdateUserMatchScore(currentUserID, userID2 string, userScore int) error {
	query := "UPDATE user_matches SET match_score = $1 WHERE user_id_1 = $2 AND user_id_2 = $3"
	_, err := DB.Exec(query, userScore, currentUserID, userID2)
	if err != nil {
		return fmt.Errorf("error updating user match score: %w", err)	
	}
	return nil
}

