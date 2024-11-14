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
		return fmt.Errorf("Error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("Error connecting to database: %w", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
	return nil
}
