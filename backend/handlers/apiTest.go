package handlers

import (
	"fmt"
	"net/http"
)

//to test any function use postman and run localhost:4000/test 

func GetTestResultHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("getting results")
}