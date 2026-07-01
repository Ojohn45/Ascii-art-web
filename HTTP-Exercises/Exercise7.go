package main

import (
	"fmt"
	"net/http"
)

func oldHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2!")
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/old", oldHandler)
	http.ListenAndServe(":8080", nil)
}
