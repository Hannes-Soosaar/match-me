package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/db"
	"net/http"
)

//to test any GET function use postman and run localhost:4000/test

func GetTestResultHandler(w http.ResponseWriter, r *http.Request){
	allCategories,err :=  db.GetAllCategories();

	if err != nil {
		log.Printf(" the error is , %s",err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Printf("all categories are: %v", allCategories)
	fmt.Println("getting results")
	json.NewEncoder(w).Encode(allCategories)
}