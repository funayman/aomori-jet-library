package model

import (
	"regexp"
	"strings"

	"github.com/funayman/aomori-library/db"
)

type Author struct {
	ID   int    `storm:"id,increment"`
	Name string `storm:"index"`
	Stub string `storm:"unique"`

	GoodReadsId   string
	OpenLibraryId string
}

func (author *Author) GetOrCreate(name string) {
	stub := BuildStub(name)
	if err := db.SQL.One("Stub", stub, author); err != nil {
		author.Name = name
		author.Stub = stub
		db.SQL.Save(author)
	}
}

func BuildStub(name string) (out string) {
	r := regexp.MustCompile("[^a-zA-Z\\d-]+")
	name = strings.Replace(name, " ", "-", -1)
	name = r.ReplaceAllString(name, "")
	out = strings.ToLower(name)
	return
}
