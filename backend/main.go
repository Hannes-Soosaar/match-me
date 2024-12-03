package main

import (
	"fmt"
	"log"
	"match_me_backend/db"
	"match_me_backend/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Backend server started with database connection verified.")

	defer db.CloseDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := routes.InitRoutes()

	log.Println("Server is running on port 4000")
	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		log.Println("Shutting down gracefully...")
		db.CloseDB()
		os.Exit(0)
	}()


}
