package main

import (
	"log"
	"match_me_backend/routes"
	"net/http"
)

func main() {
	router := routes.InitRoutes()

	log.Println("Server is running on port 4000")
	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
