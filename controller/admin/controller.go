package admin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/funayman/aomori-library/model"
	"github.com/funayman/aomori-library/model/book"
	"github.com/funayman/aomori-library/router"
)

var (
	tmpError string
	tmpBook  *book.Book
)

type BasicPage struct {
	Title string
	Books []*book.Book
}

type BookPageWithError struct {
	Title  string
	Header string
	Book   *book.Book
	Error  string
}

func Load() {
	// Unlike the normal controller running init commands,
	// All routes must be defined in the Load() function
	// According to documentation:
	// 		init() is always called, regardless if there's main or not,
	// 		if you import a package that has an init function, it will be executed
	router.Route("/api/v1/book/isbn/{isbn}", DeleteBook).Methods("DELETE")

	router.Route("/api/v1/admin/client/isbn/{isbn}", ClientIsbn)

	router.Route("/admin/book/add", BookAddGet).Methods("GET")
	router.Route("/admin/book/add", BookAddPost).Methods("POST")

	router.Route("/admin/book/{isbn}", BookEditGet).Methods("GET")
	router.Route("/admin/book/{isbn}", BookEditPost).Methods("POST")

	router.Route("/admin/books", Books).Methods("GET")
}

type DeleteResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

/* HELPER FUNCTIONS */
func writeJsonRequest(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
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
	b.Genre = r.FormValue("genre")
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
