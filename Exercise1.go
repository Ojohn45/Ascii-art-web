package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong") 
}

func main() {
	fmt.Println("Server has started")
    http.HandleFunc("/ping", HomeHandler)   // "When someone requests /, call handler"
    http.ListenAndServe(":8080", nil)
}
