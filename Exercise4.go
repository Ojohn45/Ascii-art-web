package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GETis allowed!", http.StatusMethodNotAllowed)
		return
	}

	op := r.URL.Query().Get("op")
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	if op == "" || aStr == "" || bStr == "" {
		http.Error(w, "Please provide op, a and b!", http.StatusBadRequest)
		return
	}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "a must be a number!", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "b must be a number!", http.StatusBadRequest)
		return
	}

	result := 0

	if op == "add" {
		result = a + b
	} else if op == "sub" {
		result = a - b
	} else if op == "mult" {
		result = a * b
	} else if op == "div" {
		if b == 0 {
			http.Error(w, "Cannot divide by zero!", http.StatusBadRequest)
			return
		}
		result = a / b
	} else {
		http.Error(w, "Unknown operation! use add, sub, mult, or div", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "Result: "+strconv.Itoa(result))
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/calculate", calHandler) // "When someone requests /, call handler"
	http.ListenAndServe(":8080", nil)
}
