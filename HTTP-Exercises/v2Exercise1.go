package main

import (
	"fmt"
	"net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "You made a %s request!", r.Method)
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/method-inspector", methodHandler)
	http.ListenAndServe(":8080", nil)
}
