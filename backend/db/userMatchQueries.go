package db

import (
	"fmt"
	"log"
	"match_me_backend/models"
)

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

// Need to get all connected matches
// Need to get all new matches
// Need to get all online matches
// Need to get all blocked matches

func GetAllUserMatchesByUserId(userID string) ([]models.UsersMatches, error) {
	query := "SELECT id, user_id_1, user_id_2, match_score, created_at FROM user_matches WHERE user_id_1 = $1 OR user_id_2 = $1"
	rows, err := DB.Query(query, userID)
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

func GetSecondUserIdFromMatch(userID1 string, matchID int) (string, error) {
	query := "SELECT user_id_2 FROM user_matches WHERE user_id_1 = $1 AND id = $2"
	row := DB.QueryRow(query, userID1, matchID)
	var userID2 string
	err := row.Scan(&userID2)
	if err != nil {
		return "", fmt.Errorf("error scanning row: %w", err)
	}
	return userID2, nil
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

func AddUserMatchForAllExistingUsers(newUserId string) error {

	existingUserIDs, err := GetAllUsersUuid() // returns a []string
	log.Println("New user id: ", newUserId)
	log.Println("Adding existingUserId: ", existingUserIDs)

	if err != nil {
		return fmt.Errorf("error getting all existing users: %w", err)
	}

	query := `
	INSERT INTO user_matches (user_id_1, user_id_2,match_score, status, modified_at, created_at)
	VALUES ($1, $2, $3,'new',now(), now())
	ON CONFLICT DO NOTHING;
	`

	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	for _, existingUserId := range existingUserIDs {
		userMatch, err := CalculateMatchScore(newUserId, existingUserId)
		log.Println("User match score calculated: ", userMatch);
		if err != nil {
			return fmt.Errorf("error calculating match score: %w", err)
		}
		log.Println("Adding existingUserId: ", existingUserId)
		_, err = tx.Exec(query, newUserId, existingUserId, userMatch)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error adding user match for user %s: %w", existingUserId, err)
		}

	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
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
	log.Println(currentUserID, userID2, "User match score updated", userScore)
	return nil
}

func UpdateUserMatchStatus(matchId int, status string) (string, error) {
	query := "UPDATE user_matches SET status = $1 WHERE id = $2"
	_, err := DB.Exec(query, status, matchId)
	if err != nil {
		return "", fmt.Errorf("error updating user match status: %w", err)
	}
	return status + "  was updated", nil
}
