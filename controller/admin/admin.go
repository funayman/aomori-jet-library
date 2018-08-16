package admin

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/model"
	"github.com/funayman/aomori-library/model/book"
	"github.com/funayman/aomori-library/router"
)

type page struct {
	Title string
	Books []*book.Book
}

type addEditPage struct {
	Title  string
	Header string
	Book   *book.Book
	Error  string
}

var (
	tmpError string
	tmpBook  *book.Book
)

func init() {
	router.Route("/admin/books", Books).Methods("GET")
	router.Route("/admin/book/add", BookAddGet).Methods("GET")
	router.Route("/admin/book/add", BookAddPost).Methods("POST")
	router.Route("/admin/book/{isbn}", BookIsbnGet).Methods("GET")
	router.Route("/admin/book/{isbn}", BookIsbnPost).Methods("POST")
}

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

	t.Execute(w, &page{Title: "ADMIN | Aomori JET Library", Books: books})
}

func BookAddGet(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("www/tmpl/admin/add.html", "www/tmpl/_base.html")
	if err != nil {
		log.Fatal(err)
	}

	b := book.New()
	pageErr := ""
	if tmpBook != nil {
		b = tmpBook
		tmpBook = nil

		pageErr = tmpError
		tmpError = ""
	}

	page := &addEditPage{
		Title:  "Add New Book",
		Header: "Add Book to Database",
		Error:  pageErr,
		Book:   b,
	}

	t.Execute(w, page)
}

func buildBookFromForm(r *http.Request) (b *book.Book, err string) {
	r.ParseForm()

	// build book
	b = book.New()

	b.Isbn = r.FormValue("isbn")
	b.Isbn10 = r.FormValue("isbn10")
	b.Title = r.FormValue("title")
	for _, writer := range strings.Split(r.FormValue("authors"), ";") {
		var a model.Author
		a.GetOrCreate(strings.TrimSpace(writer))
		b.Authors = append(b.Authors, a)
	}
	b.Desc = r.FormValue("desc")
	b.Lang = r.FormValue("lang")
	b.ImgSrc = r.FormValue("imgsrc")

	pages, errs := strconv.Atoi(r.FormValue("pages"))
	if errs != nil {
		// handle err
	}
	b.Pages = pages

	copies, errs := strconv.Atoi(r.FormValue("copies"))
	if errs != nil {
		// handle err
	}
	b.Copies = copies

	b.GoodReadsId = r.FormValue("goodreadsid")
	b.GoogleBooksId = r.FormValue("googlebooksid")
	b.OpenLibraryId = r.FormValue("openlibraryid")
	b.DateAdded = time.Now()

	// check for errors
	if (b.Isbn == "" && b.Isbn10 == "") || b.Title == "" || len(b.Authors) == 0 {
		err = "Required: Title, Authors, (ISBN or ISBN10)"
		return
	}

	if !(b.Isbn != "" && book.IsValidIsbn(b.Isbn)) || !(b.Isbn10 != "" && book.IsValidIsbn(b.Isbn10)) {
		err = "Invaid ISBN / ISBN10"
		return
	}

	if b.Desc == "" {
		b.Desc = book.DefaultDesc
	}

	if b.Copies < 1 {
		err = "No. of Copies must be at least 1"
		return
	}

	if b.Pages < 0 {
		err = "No. of Pages must be greater than or equal 1 (zero for N/A)"
		return
	}

	return
}

func BookAddPost(w http.ResponseWriter, r *http.Request) {

	b, err := buildBookFromForm(r)
	if err != "" {
		tmpError = err
		tmpBook = b
		BookAddGet(w, r)
		return
	}

	db.SQL.Save(b)

	http.Redirect(w, r, "/admin/books", 302)
}

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

	page := &addEditPage{
		Title:  "Edit Book",
		Header: "Editing " + b.Title,
		Book:   &b,
	}

	t.Execute(w, page)
}

func BookIsbnPost(w http.ResponseWriter, r *http.Request) {
	b, err := buildBookFromForm(r)
	if err != "" {
		tmpError = err
		tmpBook = b
		BookAddGet(w, r)
		return
	}

	db.SQL.Save(b)

	http.Redirect(w, r, "/admin/books", 302)
}
