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

var ebook *book.Book

func BookEditGet(w http.ResponseWriter, r *http.Request) {
	vars := router.GetParams(r)
	isbn := vars["isbn"]

	t, err := template.ParseFiles("www/tmpl/admin/add.html", "www/tmpl/_nav.html", "www/tmpl/_base.html")
	if err != nil {
		log.Println(err)
	}

	var b book.Book
	err = db.SQL.One("Isbn", isbn, &b)
	if err != nil {
		log.Println(err)
	}
	ebook = &b

	t.Execute(w, &BookPageWithError{
		Title:  "Edit Book",
		Header: "Editing " + b.Title,
		Book:   &b,
	})
}

func BookEditPost(w http.ResponseWriter, r *http.Request) {
	b, err := buildBookFromForm(r)
	if err != "" {
		tmpError = err
		tmpBook = b
		BookAddGet(w, r)
		return
	}

	// if ImgSrc is different then save
	if b.ImgSrc != ebook.ImgSrc {
		imgsrc, e := client.SaveCover(b.Isbn, b.ImgSrc)
		if e != nil {
			log.Printf("[BookIsbnPost] %s", e.Error())
		} else {
			b.ImgSrc = imgsrc
		}
	}

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
