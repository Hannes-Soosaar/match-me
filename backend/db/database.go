package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connectionString := "host=localhost port=5432 user=MatchMeDev password=SecretDevPassword dbname=Match-Me-Data sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
	return nil
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUserByEmail(email string) (*User, error) {
	query := "SELECT id, email, password_hash FROM users WHERE email = $1"

	var user User
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

func SaveUser(email string, password_hash string) error {
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2)"
	_, err := DB.Exec(query, email, password_hash)
	return err
}
