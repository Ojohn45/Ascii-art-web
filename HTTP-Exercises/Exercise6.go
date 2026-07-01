package main

import (
	"fmt"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	const apikey = "secret123"

	key := r.Header.Get("X-APT-Key")

	if key == "" {
		http.Error(w, "No API Key provided!", http.StatusUnauthorized)
		return
	}

	if key != apikey {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "Welcome to the dashboard Minato!")
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/dashboard", dashboardHandler)
	http.ListenAndServe(":8080", nil)
}
