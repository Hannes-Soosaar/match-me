package handlers

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Login page!", "text/plain")

	fmt.Fprintln(w, "Welcome to the login page!")
}
