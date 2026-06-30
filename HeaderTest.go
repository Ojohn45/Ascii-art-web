package main

import (
	"fmt"
	"net/http"
)

func headerHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")

	w.WriteHeader(200)
	fmt.Fprint(w, "Your User-Agent is: "+userAgent)
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/whoami", headerHandler) // "When someone requests /, call handler"
	http.ListenAndServe(":8080", nil)
}
