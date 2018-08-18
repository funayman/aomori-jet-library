package admin

import (
	"html/template"
	"log"
	"net/http"

	"github.com/funayman/aomori-library/client"
	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model/book"
)

func BookAddGet(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("www/tmpl/admin/add.html", "www/tmpl/_nav.html", "www/tmpl/_base.html")
	if err != nil {
		log.Printf("[BookAddGet] %s", err.Error())
	}

	b := book.New()
	pageErr := ""
	if tmpBook != nil {
		b = tmpBook
		tmpBook = nil

		pageErr = tmpError
		tmpError = ""
	}

	t.Execute(w, &BookPageWithError{
		Title:  "Add New Book",
		Header: "Add Book to Database",
		Error:  pageErr,
		Book:   b,
	})
}

func BookAddPost(w http.ResponseWriter, r *http.Request) {

	b, err := buildBookFromForm(r)
	if err != "" {
		tmpError = err
		tmpBook = b
		BookAddGet(w, r)
		return
	}

	imgsrc, e := client.SaveCover(b.Isbn, b.ImgSrc)
	if e != nil {
		log.Printf("[BookAddPost] %s", e.Error())
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
