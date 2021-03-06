package controller

import (
	"log"
	"net/http"

	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model/book"
	"github.com/funayman/aomori-library/router"
)

type page struct {
	Title string
	Books []*book.Book
}

func init() {
	router.Route("/", Index).Methods("GET")
}

func Index(w http.ResponseWriter, r *http.Request) {
	var books []*book.Book
	err := db.SQL.All(&books)
	if err != nil {
		log.Println(err)
	}

	// t, err := template.ParseFiles("www/tmpl/index.html", "www/tmpl/_nav.html", "www/tmpl/_base.html")
	// if err != nil {
	// 	log.Println(err)
	// }

	if err := mT["index.html"].Execute(w, &page{Title: "Aomori JET Library", Books: books}); err != nil {
		log.Println(err)
	}
}
