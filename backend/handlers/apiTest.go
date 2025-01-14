package handlers

import (
	"encoding/json"
	"log"

	"net/http"
)

//to test any GET function use postman and run localhost:4000/test

func GetTestResultHandler(w http.ResponseWriter, r *http.Request) {


	

	successMessage := "Nothing to see here, move along!"
	log.Println("Somebody want to test something")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successMessage)
}
