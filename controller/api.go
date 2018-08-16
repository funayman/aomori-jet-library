package controller

import (
	"encoding/json"
	"net/http"

	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model/book"
	"github.com/funayman/aomori-library/router"
)

func init() {
	router.Route("/api/v1/books", Books).Methods("GET")
	router.Route("/api/v1/book/{isbn}", BookIsbn).Methods("GET")
}

func Books(w http.ResponseWriter, r *http.Request) {
	var books []book.Book
	db.SQL.All(&books)
	json.NewEncoder(w).Encode(&books)
}

func BookIsbn(w http.ResponseWriter, r *http.Request) {
	vars := router.GetParams(r)
	var book book.Book
	db.SQL.One("Isbn", vars["isbn"], &book)

	encoder := json.NewEncoder(w)

	if book.Isbn == "" && book.Isbn10 == "" {
		w.WriteHeader(http.StatusNoContent)
		// encoder.Encode(&struct {
		// 	Error string `json:"error"`
		// }{Error: "book not found"})
	} else {
		w.WriteHeader(http.StatusOK)
		encoder.Encode(&book)
	}

}
