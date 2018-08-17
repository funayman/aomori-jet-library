package admin

import (
	"encoding/json"
	"net/http"

	"github.com/funayman/aomori-library/client"
	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model/book"
	"github.com/funayman/aomori-library/router"
)

func ClientIsbn(w http.ResponseWriter, r *http.Request) {
	vars := router.GetParams(r)
	isbn := vars["isbn"]

	if !book.IsValidIsbn(isbn) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{error:\"not valid isbn\"}"))
		return
	}

	book := client.FindBookByIsbn(isbn)
	if book.Isbn == "" || book.Isbn10 == "" {
		// not a valid book
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := router.GetParams(r)
	isbn := vars["isbn"]

	if !book.IsValidIsbn(isbn) {
		writeJsonRequest(w, http.StatusBadRequest, &DeleteResponse{false, "not a valid isbn"})
		return
	}

	var b book.Book
	dberr := db.SQL.One("Isbn", isbn, &b)
	if dberr != nil {
		writeJsonRequest(w, http.StatusNoContent, &DeleteResponse{false, dberr.Error()})
		return
	}

	dberr = db.SQL.DeleteStruct(&b)
	if dberr != nil {
		writeJsonRequest(w, http.StatusBadRequest, &DeleteResponse{false, dberr.Error()})
		return
	}

	writeJsonRequest(w, http.StatusOK, &DeleteResponse{Success: true})
	return
}
