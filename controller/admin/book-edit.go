package admin

import (
	"html/template"
	"log"
	"net/http"

	"github.com/funayman/aomori-library/client"
	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model/book"
	"github.com/funayman/aomori-library/router"
)

func BookIsbnGet(w http.ResponseWriter, r *http.Request) {
	vars := router.GetParams(r)
	isbn := vars["isbn"]

	t, err := template.ParseFiles("www/tmpl/admin/add.html", "www/tmpl/_base.html")
	if err != nil {
		log.Println(err)
	}

	var b book.Book
	err = db.SQL.One("Isbn", isbn, &b)
	if err != nil {
		log.Println(err)
	}

	t.Execute(w, &BookPageWithError{
		Title:  "Edit Book",
		Header: "Editing " + b.Title,
		Book:   &b,
	})
}

func BookIsbnPost(w http.ResponseWriter, r *http.Request) {
	b, err := buildBookFromForm(r)
	if err != "" {
		tmpError = err
		tmpBook = b
		BookAddGet(w, r)
		return
	}

	imgsrc, e := client.SaveCover(b.Isbn, b.ImgSrc)
	if e != nil {
		log.Fatal(e)
	}
	b.ImgSrc = imgsrc
	dberr := db.SQL.Save(b)
	if dberr != nil {
		if dberr.Error() == "database is in read-only mode" {
			tmpError = dberr.Error() + ", please restart the server"
		} else {
			tmpError = dberr.Error() + ", call Dave to fix"
		}

		tmpBook = b
		BookAddGet(w, r)
		return
	}

	http.Redirect(w, r, "/admin/books", 302)
}
