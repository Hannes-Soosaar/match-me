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

func GetUserUUIDFromUserEmail(email string) (string, error) {
	query := "SELECT uuid FROM users WHERE email = $1"
	var userUUID string
	err := DB.QueryRow(query, email).Scan(&userUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found with that email")
		}
		return "", fmt.Errorf("error querying the database: %v", err)
	}
	return userUUID, nil
}

func GetAllUsersUuid() ([]string, error) {
	query := "SELECT uuid FROM users"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()
	var uuids []string
	for rows.Next() {
		var uuid string
		err = rows.Scan(&uuid)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		uuids = append(uuids, uuid)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iterations: %w", err)
	}
	return uuids, nil
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
	if err != nil {
		log.Printf("Error adding user match for all existing users: %v", err)
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
            p.username, u.email, u.created_at, u.user_city, u.user_nation, u.user_region, 
            p.about_me, p.birthdate, p.profile_picture
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
	var nation sql.NullString
	var region sql.NullString
	var about sql.NullString
	var birthdate sql.NullTime
	var picture sql.NullString

	err := DB.QueryRow(query, userID).Scan(
		&username,
		&email,
		&created,
		&city,
		&nation,
		&region,
		&about,
		&birthdate,
		&picture,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not found for uuid=%v: %v", userID, err)
			return nil, fmt.Errorf("user not found: %w", err)
		}
		log.Printf("Error querying user by ID: %v", err)
		return nil, fmt.Errorf("error querying user by ID: %w", err)
	}

	// Calculate age if birthdate is valid
	var age int
	if birthdate.Valid {
		birthdayTime := birthdate.Time
		currentTime := time.Now()
		age = currentTime.Year() - birthdayTime.Year()
		// Adjust age if the birthday hasn't occurred yet this year
		if birthdayTime.After(currentTime.AddDate(-age, 0, 0)) {
			age--
		}
	}

	// Check if any fields are NULL and assign them to appropriate defaults
	userInfo.Username = ""
	if username.Valid {
		userInfo.Username = username.String
	}

	userInfo.Email = ""
	if email.Valid {
		userInfo.Email = email.String
	}

	userInfo.Created = time.Time{}
	if created.Valid {
		userInfo.Created = created.Time
	}

	userInfo.City = ""
	if city.Valid {
		userInfo.City = city.String
	}

	userInfo.Nation = ""
	if nation.Valid {
		userInfo.Nation = nation.String
	}

	userInfo.Region = ""
	if region.Valid {
		userInfo.Region = region.String
	}

	userInfo.About = ""
	if about.Valid {
		userInfo.About = about.String
	}

	userInfo.Birthdate = time.Time{}
	if birthdate.Valid {
		userInfo.Birthdate = birthdate.Time
	}

	userInfo.Picture = ""
	if picture.Valid {
		userInfo.Picture = picture.String
	}

	// Set calculated age
	userInfo.Age = fmt.Sprintf("%d", age) // Convert age to string for userInfo

	return &userInfo, nil
}

// Change the "satus" of a connection
func ModifyUserConnection(userID int) error {
	return nil
}

func RemoveUserConnection(currentUserID, userID2 int) error {
	// GET the logged in userID from session to avoid potential
	return nil
}

func DeleteUser(email string) error {
	query := "DELETE FROM users WHERE email = $1"
	_, err := DB.Exec(query, email)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
