package admin

import (
	"encoding/json"
	"net/http"

	"github.com/funayman/aomori-library/client"
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
