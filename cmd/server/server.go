package main

import (
	"encoding/json"
	"net/http"

	"github.com/RaphaelMarquesMartorella/go-project/domain/model"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/Books/Create", CreateBook)
	http.ListenAndServe(":3000", nil)
}
