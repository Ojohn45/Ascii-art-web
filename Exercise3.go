package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Could not read body!", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		w.WriteHeader(200)
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == http.MethodPost {
		w.WriteHeader(200)
		count := len(body)
		fmt.Fprint(w, "Characters: "+strconv.Itoa(count))
		return
	}
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/count", countHandler) // "When someone requests /, call handler"
	http.ListenAndServe(":6060", nil)
}
