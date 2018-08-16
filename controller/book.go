package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model/book"
	"github.com/funayman/aomori-library/router"
)

func init() {
	router.Route("/book/{isbn}", ViewBook)
}

func ViewBook(w http.ResponseWriter, r *http.Request) {
	vars := router.GetParams(r)
	isbn := vars["isbn"]

	t, err := template.ParseFiles("www/tmpl/view-book.html", "www/tmpl/_base.html")
	if err != nil {
		log.Printf("[ViewBook:%s] - %s", r.URL.EscapedPath(), err.Error())
	}

	var b book.Book
	err = db.SQL.One("Isbn", isbn, &b)
	if err != nil {
		log.Printf("[ViewBook:%s] - %s", r.URL.EscapedPath(), err.Error())
	}

	title := fmt.Sprintf("%s | %s", "Aomori JET Library", b.Title)
	t.Execute(w, &BookPage{Title: title, Book: &b})
}
