package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"match_me_backend/models"
	"time"

	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT uuid, email, password_hash FROM users WHERE email = $1"
	var user models.User
	err := DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with that email")
		}
		fmt.Printf("error querying the database: %v", err)
		log.Printf("Error querying the database: %v", err)
		return nil, fmt.Errorf("error querying the database: %v", err)
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT u.uuid, u.email, u.password_hash FROM users u JOIN profiles p ON u.uuid = p.uuid WHERE p.username = $1"
	var user models.User
	err := DB.QueryRow(query, username).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with that user")
		}
		fmt.Printf("error querying the database: %v", err)
		log.Printf("Error querying the database: %v", err)
		return nil, fmt.Errorf("error querying the database: %v", err)
	}
	return &user, nil
}

func GetUserByID(userID string) (*models.User, error) {
	query := "SELECT uuid, email, password_hash FROM users WHERE uuid = $1"
	var user models.User

	row := DB.QueryRow(query, userID)
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user not found: %v", err)
			return nil, fmt.Errorf("user not found: %w", err)
		}
		log.Printf("error querying user by ID: %v", err)
		return nil, fmt.Errorf("error querying user by ID: %w", err)
	}

	return &user, nil
}

func SaveUser(email string, password_hash string) error {
	userUUID := uuid.New()

	tx, err := DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return err
	}

	userQuery := "INSERT INTO users (uuid, email, password_hash) VALUES ($1, $2, $3)"
	_, err = tx.Exec(userQuery, userUUID, email, password_hash)
	if err != nil {
		tx.Rollback()
		log.Printf("Error inserting into users table: %v", err)
		return err
	}

	profileQuery := "INSERT INTO profiles (uuid) VALUES ($1)"
	_, err = tx.Exec(profileQuery, userUUID)
	if err != nil {
		tx.Rollback()
		log.Printf("Error inserting into profiles table: %v", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}

	return nil
}

func GetUserConnectionsByUserID(userID int) (*[]models.UserConnections, error) {
	query := "SELECT * FROM users WHERE uuid = $1" // TODO: need to add check for status
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

func GetUserInformation(userID string) (*models.ProfileInformation, error) {
	query := `
        SELECT 
            p.username, u.email, u.created_at, u.user_city, 
            p.about_me, p.birthdate
        FROM 
            users u 
        JOIN 
            profiles p 
        ON 
            u.uuid = p.uuid 
        WHERE 
            u.uuid = $1`

	var userInfo models.ProfileInformation

	var username sql.NullString
	var email sql.NullString
	var created sql.NullTime
	var city sql.NullString
	var about sql.NullString
	var birthdate sql.NullTime

	err := DB.QueryRow(query, userID).Scan(
		&username,
		&email,
		&created,
		&city,
		&about,
		&birthdate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not found for uuid=%v: %v", userID, err)
			return nil, fmt.Errorf("user not found: %w", err)
		}
		log.Printf("Error querying user by ID: %v", err)
		return nil, fmt.Errorf("error querying user by ID: %w", err)
	}

	// Check if any fields are NULL and assign them to appropriate defaults
	userInfo.Username = username.String
	if !username.Valid {
		userInfo.Username = "" // If the value is NULL, set it to an empty string or whatever default you want
	}

	userInfo.Email = email.String
	if !email.Valid {
		userInfo.Email = "" // Default to empty string if NULL
	}

	userInfo.Created = created.Time
	if !created.Valid {
		userInfo.Created = time.Time{} // Default to zero time if NULL
	}

	userInfo.City = city.String
	if !city.Valid {
		userInfo.City = "" // Default to empty string if NULL
	}

	userInfo.About = about.String
	if !about.Valid {
		userInfo.About = "" // Default to empty string if NULL
	}

	userInfo.Birthdate = birthdate.Time
	if !birthdate.Valid {
		userInfo.Birthdate = time.Time{} // Default to zero time if NULL
	}

	return &userInfo, nil
}

func AddUserConnection(userID1 int, userID2 int) error {
	return nil
}

// Change the "satus" of a connection
func ModifyUserConnection(userID int) error {
	return nil
}

func RemoveUserConnection(currentUserID, userID2 int) error {
	// GET the logged in userID from session to avoid potential
	return nil
}
