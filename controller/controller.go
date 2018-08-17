package controller

import (
	"github.com/funayman/aomori-library/model/book"
)

func Load() {
	// Do nothing!
	// Forces init() to be called on all files in controller package
}

type BookPage struct {
	Title string
	Book  *book.Book
}
