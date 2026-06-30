package main

import (
	"fmt"
	"io"
	"net/http"
)

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Could not read body!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "You sent: "+string(body))
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/hello", bodyHandler) // "When someone requests /, call handler"
	http.ListenAndServe(":8080", nil)
}
