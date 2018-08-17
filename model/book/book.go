package book

import (
	"regexp"
	"strings"
	"time"

	"github.com/funayman/aomori-library/model"
)

const (
	Delim       = ";"
	DefaultDesc = "No description available (╯°□°）╯︵ ┻━┻"
)

type Book struct {
	Isbn    string `storm:"id"`
	Isbn10  string `storm:"unique"`
	Title   string
	Authors []model.Author `storm:"inline"`
	Lang    string
	ImgSrc  string
	Pages   int
	Desc    string
	Copies  int
	Genre   string

	GoodReadsId   string
	OpenLibraryId string
	GoogleBooksId string

	DatePublished string
	DateAdded     time.Time

	Category []string
	//Reviews     []Review
}

func New() *Book {
	return &Book{}
}

func (b *Book) PrintAuthors() (auth string) {
	if len(b.Authors) == 0 {
		auth = ""
		return
	}

	auth = b.Authors[0].Name
	for i := 1; i < len(b.Authors); i++ {
		auth += "; "
		auth += b.Authors[i].Name
	}

	return
}

func (b *Book) Merge(book *Book) error {
	if b.Isbn == "" && book.Isbn != "" {
		b.Isbn = book.Isbn
	}

	if b.Isbn10 == "" && book.Isbn10 != "" {
		b.Isbn10 = book.Isbn10
	}

	if b.Title == "" && book.Title != "" {
		b.Title = book.Title
	}

	if len(b.Authors) == 0 {
		b.Authors = book.Authors
	} else if len(book.Authors) > 0 {
		// compare authors and add them
	}

	if b.Lang == "" && book.Lang != "" {
		b.Lang = book.Lang
	}

	if (b.ImgSrc == "" || strings.Contains(b.ImgSrc, "isbndb")) && book.ImgSrc != "" && !strings.Contains(b.ImgSrc, "nophoto") {
		b.ImgSrc = book.ImgSrc
	}

	if b.Pages == 0 && book.Pages != 0 {
		b.Pages = book.Pages
	}

	if b.Desc == "" && book.Desc != "" {
		b.Desc = cleanDescription(book.Desc)
	}

	if b.GoodReadsId == "" && book.GoodReadsId != "" {
		b.GoodReadsId = book.GoodReadsId
	}

	if b.DatePublished == "" && book.DatePublished != "" {
		b.DatePublished = book.DatePublished
	}

	if len(book.Category) != 0 {
		c := make(map[string]int)
		for i, k := range b.Category {
			c[k] = i
		}
		for _, k := range book.Category {
			if _, ok := c[k]; !ok {
				b.Category = append(b.Category, k)
			}
		}
	}
	return nil
}

func (b *Book) ShortDesc() string {
	limit := 300
	rdesc := []rune(b.Desc)
	var str string

	if len(rdesc) > limit {
		str = string(rdesc[:limit])
		s := strings.Split(str, " ")
		str = strings.Join(s[:len(s)-1], " ") + "..."
	} else {
		str = b.Desc
	}

	return str
}

func IsValidIsbn(isbn string) bool {
	if !(len(isbn) == 10 || len(isbn) == 13) {
		return false
	}

	r := regexp.MustCompile("[0-9]+")
	if !(r.MatchString(isbn)) {
		return false
	}

	return true
}

func cleanDescription(s string) string {
	// https://haacked.com/archive/2004/10/25/usingregularexpressionstomatchhtml.aspx/
	r := regexp.MustCompile("</?\\w+((\\s+\\w+(\\s*=\\s*(?:\".*?\"|'.*?'|[\\^'\">\\s]+))?)+\\s*|\\s*)/?>")
	return r.ReplaceAllString(s, "")
}
