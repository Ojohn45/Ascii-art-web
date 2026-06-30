package main

import (
	"fmt"
	"net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(200)
		fmt.Fprint(w, "You used GET - here is some data!")
		return
	}

	if r.Method == http.MethodPost {
		w.WriteHeader(200)
		fmt.Fprint(w, "You used POST - data received")
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/hello", methodHandler) // "When someone requests /, call handler"
	http.ListenAndServe(":8080", nil)
}
