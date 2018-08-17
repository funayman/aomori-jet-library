package admin

import (
	"html/template"
	"log"
	"net/http"

	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model/book"
)

func Books(w http.ResponseWriter, r *http.Request) {
	var books []*book.Book
	err := db.SQL.All(&books)
	if err != nil {
		log.Println(err)
	}

	t, err := template.ParseFiles("www/tmpl/admin/books.html", "www/tmpl/_base.html")
	if err != nil {
		log.Println(err)
	}

	t.Execute(w, &BasicPage{Title: "ADMIN | Aomori JET Library", Books: books})
}
