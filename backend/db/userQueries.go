package db

import (
	"database/sql"
	"fmt"
	"match_me_backend/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, email, password_hash FROM users WHERE email = $1"
	var user models.User
	err := DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with that email")
		}
		fmt.Printf("error querying the database: %v", err)
		return nil, fmt.Errorf("error querying the database: %v", err)
	}
	return &user, nil
}

//Rename to AddUser
func SaveUser(email string, password_hash string) error {
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2)"
	_, err := DB.Exec(query, email, password_hash)
	return err
}

func GetUserByID(userID int) (*models.User, error) {
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

func GetUserConnectionsByUserID(userID int) (*[]models.UserConnections, error) {
	query := "SELECT * FROM users WHERE id = $1" // TODO: need to add check for status 
	rows, err := DB.Query(query, userID)
	var connections []models.UserConnections
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var connection models.UserConnections
		if err := rows.Scan(&connection.ID, &connection.UserID1, &connection.UserID2, &connection.Status, &connection.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		connections = append(connections, connection)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}
	return &connections, nil
}

func AddUserConnection(userID1 int, userID2 int)(error){
	return nil
}

//Change the "satus" of a connection
func ModifyUserConnection(userID int )(error){
	return nil
}

func RemoveUserConnection(currentUserID, userID2 int) (error){
	// GET the logged in userID from session to avoid potential 
	return nil 
}

