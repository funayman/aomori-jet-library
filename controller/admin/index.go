package admin

import (
	"log"
	"net/http"

	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model/book"
)

type adminPage struct {
	Title string
	Total int
	Books []*book.Book
}

func Index(w http.ResponseWriter, r *http.Request) {
	var books []*book.Book

	err := db.SQL.Select().OrderBy("DateAdded").Reverse().Limit(5).Find(&books)
	if err != nil {
		log.Println(err)
	}

	n, err := db.SQL.Count(book.New())
	if err != nil {
		log.Println(err)
	}

	err = mT["index.html"].Execute(w, &adminPage{Title: "ADMIN | Home", Total: n, Books: books})
	if err != nil {
		log.Println(err)
	}
}
