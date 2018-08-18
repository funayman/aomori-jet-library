package controller

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/funayman/aomori-library/model/book"
)

const (
	templateDir = "www/tmpl"
)

var mT map[string]*template.Template

func init() {
	mT = make(map[string]*template.Template)

	var templates []string
	var required []string
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// paths = append(paths, path)
			if filepath.Base(path)[0] == '_' {
				required = append(required, path)
			} else {
				templates = append(templates, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range templates {
		fn := filepath.Base(t)
		tm := template.Must(template.ParseFiles(append([]string{t}, required...)...))
		if err != nil {
			log.Fatal(err)
		}
		mT[fn] = tm
	}

}

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
