package db

import (
	"database/sql"
	"fmt"
	"match_me_backend/models"
)

func GetUserByID(userID int)(*models.User, error){
	query := "SELECT id, email FROM users WHERE id = $1"
	row := DB.QueryRow(query, userID)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	return &user, nil
}

func GetUserConnectionsByUserID(userID int)(*models.UserConnections, error){
	var connections models.UserConnections;
	return &connections, nil
}