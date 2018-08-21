package client

import (
	"net/http"

	"github.com/funayman/aomori-library/model/book"
)

var (
	clients []Client
	client  *http.Client
)

type Cfg struct {
	IsbnDbKey    string
	GoodReadsKey string
}

type Client interface {
	Query(q string, p map[string]string) []*book.Book
	QueryIsbn(isbn string) *book.Book
}

func Init(cfg *Cfg) {
	client = &http.Client{}
	clients = []Client{
		&IsbnDbClient{key: cfg.IsbnDbKey, client: client},
		&GoodReadsClient{key: cfg.GoodReadsKey, client: client},
		&OpenLibraryClient{client: client},
	}
}

func FindBookByIsbn(isbn string) *book.Book {
	book := book.New()
	for _, client := range clients {
		book.Merge(client.QueryIsbn(isbn))
	}
	return book
}
