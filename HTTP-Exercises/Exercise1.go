package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/ping", pingHandler) // "When someone requests /, call handler"
	http.ListenAndServe(":8080", nil)
}
