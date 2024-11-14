package main

import (
	"fmt"
	"log"
	"match_me_backend/db"
	"match_me_backend/routes"
	"net/http"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Backend server started with database connection verified.")

	router := routes.InitRoutes()

	log.Println("Server is running on port 4000")
	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
