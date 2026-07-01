package main

import (
	"fmt"
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")

	if userAgent == "" {
		userAgent = "Unknown client"
	}
	w.WriteHeader(200)
	fmt.Fprint(w, "You are visiting us using: "+userAgent)
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/agent", HeaderHandler)
	http.ListenAndServe(":8080", nil)
}
