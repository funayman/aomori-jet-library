package client

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/funayman/aomori-library/model"
	"github.com/funayman/aomori-library/model/book"
)

const (
	IsbnToIdUrl = "https://www.goodreads.com/book/isbn_to_id"
	BooksUrl    = "https://www.goodreads.com/book/isbn"
)

type GoodReadsClient struct {
	key    string
	secret string
	client http.Client
}

type goodReadsResponse struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Book    struct {
		Id            int    `xml:"id"`
		Title         string `xml:"title"`
		Isbn13        string `xml:"isbn13"`
		Isbn10        string `xml:"isbn"`
		ImageUrl      string `xml:"image_url"`
		SmallImageUrl string `xml:"small_image_url"`
		Year          string `xml:"publication_year"`
		Month         string `xml:"publication_month"`
		LanguageCode  string `xml:"language_code"`
		Description   string `xml:"description"`
		AverageRating string `xml:"average_rating"`
		NumPage       int    `xml:"num_pages"`
		Authors       []struct {
			Id   string `xml:"id"`
			Name string `xml:"name"`
		} `xml:"authors>author"`
	} `xml:"book"`
}

func (c *GoodReadsClient) Query(q string, p map[string]string) (books []*book.Book) {
	return
}

func (c *GoodReadsClient) QueryIsbn(isbn string) (b *book.Book) {
	b = book.New()
	url := fmt.Sprintf("%s/%s?key=%s", BooksUrl, isbn, c.key)
	resp, err := c.client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data goodReadsResponse
	err = xml.Unmarshal(body, &data)
	if err != nil {
		log.Println("[openlibrary] error parsing json", err)
		return
	}

	b = c.buildBooksFromResponse(data)
	return
}

func (c *GoodReadsClient) buildBooksFromResponse(data goodReadsResponse) *book.Book {
	b := data.Book
	bb := book.New()

	bb.Isbn = b.Isbn13
	bb.Isbn10 = b.Isbn10
	bb.Title = b.Title
	bb.Lang = b.LanguageCode
	bb.ImgSrc = b.ImageUrl
	bb.Pages = b.NumPage
	bb.Desc = b.Description

	if len(b.Authors) != 0 {
		for _, w := range b.Authors {
			if w.Name == "" {
				continue
			}

			a := model.Author{}
			a.GetOrCreate(w.Name)

			// TO-DO: add Author.GoodReadsID by stripping w.ID

			bb.Authors = append(bb.Authors, a)
		}
	}

	if b.Year != "" {
		bb.DatePublished = b.Year
		if b.Month != "" {
			bb.DatePublished = bb.DatePublished + " " + b.Month
		}
	}

	return bb
}
