package handlers

import (
	"fmt"
	"net/http"
	"log"
	"match_me_backend/db"
)

//to test any GET function use postman and run localhost:4000/test 

func GetTestResultHandler(w http.ResponseWriter, r *http.Request){
	allCategories,err :=  db.GetAllCategories();

	if err != nil {
		log.Printf(" the error is , %s",err)
	}

	fmt.Printf("all categories are: %v", allCategories)

	fmt.Println("getting results")
}