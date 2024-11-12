package handlers

import (
	"fmt"
	"net/http"
)

func GetLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Login page!", "text/plain")

	fmt.Fprintln(w, "Welcome to the login page!")
}
