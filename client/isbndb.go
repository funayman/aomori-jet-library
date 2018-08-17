package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/funayman/aomori-library/model"
	"github.com/funayman/aomori-library/model/book"
)

const (
	BaseUrl = "https://api.isbndb.com"
)

type IsbnDbClient struct {
	key    string
	secret string
	client *http.Client
}

type IsbnDbResponse struct {
	Book struct {
		Title         string   `json:"title"`
		TitleLong     string   `json:"title_long"`
		Isbn          string   `json:"isbn"`
		Isbn13        string   `json:"isbn13"`
		DeweyDecimal  string   `json:"dewey_decimal"`
		Format        string   `json:"format"`
		Publisher     string   `json:"publisher"`
		Language      string   `json:"language"`
		DatePublished string   `json:"date_published"`
		Edition       string   `json:"edition"`
		Pages         int      `json:"pages"`
		Dimensions    string   `json:"dimensions"`
		Overview      string   `json:"overview"`
		Image         string   `json:"image"`
		Excerpt       string   `json:"excerpt"`
		Synopsys      string   `json:"synopsys"`
		Authors       []string `json:"authors"`
		Subjects      []string `json:"subjects"`
		Reviews       []string `json:"reviews"`
	} `json:"book"`
}

func (c *IsbnDbClient) Query(q string, p map[string]string) (books []*book.Book) {
	return
}

func (c *IsbnDbClient) QueryIsbn(isbn string) (b *book.Book) {
	b = book.New()
	url := fmt.Sprintf("%s/book/%s", BaseUrl, isbn)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-API-Key", c.key)

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	byteData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data IsbnDbResponse
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		log.Println("[isbndb] error parsing json", err)
		return
	}

	b = c.convertToBooks(data)
	return
}

func (c *IsbnDbClient) convertToBooks(b IsbnDbResponse) *book.Book {
	bb := book.New()

	bb.Title = b.Book.Title
	bb.Isbn = b.Book.Isbn13
	bb.Isbn10 = b.Book.Isbn
	bb.Lang = b.Book.Language
	bb.Pages = b.Book.Pages
	bb.Desc = b.Book.Synopsys
	bb.DatePublished = b.Book.DatePublished
	bb.ImgSrc = b.Book.Image
	bb.Category = b.Book.Subjects

	if len(b.Book.Authors) != 0 {
		for _, name := range b.Book.Authors {
			if name == "" {
				continue
			}
			a := model.Author{}
			a.GetOrCreate(name)
			bb.Authors = append(bb.Authors, a)
		}
	}

	return bb
}
