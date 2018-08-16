package client

import (
	"fmt"
	"net/http"

	"github.com/funayman/aomori-library/model/book"
)

var (
	clients []Client
)

const (
	IsbnDbKey    = "IsbnDB API Key"
	GoodReadsKey = "GoodReads API Key"
)

type Client interface {
	Query(q string, p map[string]string) []*book.Book
	QueryIsbn(isbn string) *book.Book
}

func Init() {
	clients = []Client{
		&IsbnDbClient{key: IsbnDbKey, client: http.Client{}},
		&GoodReadsClient{key: GoodReadsKey, client: http.Client{}},
		&OpenLibraryClient{client: http.Client{}},
	}
}

func FindBookByIsbn(isbn string) *book.Book {
	book := book.New()
	for _, client := range clients {
		book.Merge(client.QueryIsbn(isbn))
	}
	return book
}

func buildGetParams(params map[string]string) (p string) {
	for k, v := range params {
		p = fmt.Sprintf("%s%s=%s&", p, k, v)
	}

	return
}
