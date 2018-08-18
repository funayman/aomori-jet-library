package controller

import (
	"github.com/funayman/aomori-library/model/book"
)

func Load() {
	// Do nothing!
	// Forces init() to be called on all files in controller package

	/*
	 * Do things to organize easier templates and rendering:
	 * https://stackoverflow.com/questions/11467731/is-it-possible-to-have-nested-templates-in-go-using-the-standard-library-googl/11468132#11468132
	 * https://stackoverflow.com/questions/24093923/optimising-html-template-composition
	 * https://www.reddit.com/r/golang/comments/27ls5a/including_htmltemplate_snippets_is_there_a_better/
	 */
}

type BookPage struct {
	Title string
	Book  *book.Book
}
