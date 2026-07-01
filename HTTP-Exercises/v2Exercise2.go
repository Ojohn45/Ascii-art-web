package main

import (
	"fmt"
	"io"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed, only POST!", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read body!", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		http.Error(w, "body cannot be empty", http.StatusBadRequest)
	}
	fmt.Fprint(w, string(body))

}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(":8080", nil)
}
