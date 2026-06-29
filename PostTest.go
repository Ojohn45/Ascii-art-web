package main

import (
	"fmt"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	message := r.FormValue("message")
	if name == "" {
		http.Error(w, "Please send a name!", http.StatusBadRequest)
		return
	}

	if message == "" {
		http.Error(w, "Please send a Message!", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, name+" says: "+message)
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/hello", postHandler) // "When someone requests /, call handler"
	http.ListenAndServe(":8080", nil)
}
