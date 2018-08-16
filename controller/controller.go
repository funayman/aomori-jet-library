package controller

import (
	"github.com/funayman/aomori-library/controller/admin"
	"github.com/funayman/aomori-library/model/book"
)

func Load() {
	// Do nothing!
	// Forces init() to be called
	admin.Load()
}

type BookPage struct {
	Title string
	Book  *book.Book
}
