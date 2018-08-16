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

type OpenLibraryClient struct {
	key    string
	secret string
	client http.Client
}

type openLibraryApiResponse struct {
	Authors []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"authors"`
	ByStatement     string `json:"by_statement"`
	Classifications struct {
		DeweyDecimalClass []string `json:"dewey_decimal_class"`
		LcClassifications []string `json:"lc_classifications"`
	} `json:"classifications"`
	Cover struct {
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"cover"`
	Ebooks []struct {
		Availability string `json:"availability"`
		Formats      struct {
		} `json:"formats"`
		PreviewURL string `json:"preview_url"`
	} `json:"ebooks"`
	Identifiers struct {
		Google       []string `json:"google"`
		Goodreads    []string `json:"goodreads"`
		Isbn10       []string `json:"isbn_10"`
		Isbn13       []string `json:"isbn_13"`
		Lccn         []string `json:"lccn"`
		Librarything []string `json:"librarything"`
		Openlibrary  []string `json:"openlibrary"`
		Wikidata     []string `json:"wikidata"`
	} `json:"identifiers"`
	Key           string `json:"key"`
	Notes         string `json:"notes"`
	NumberOfPages int    `json:"number_of_pages"`
	Pagination    string `json:"pagination"`
	PublishDate   string `json:"publish_date"`
	PublishPlaces []struct {
		Name string `json:"name"`
	} `json:"publish_places"`
	Publishers []struct {
		Name string `json:"name"`
	} `json:"publishers"`
	Subjects []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"subjects"`
	Subtitle string `json:"subtitle"`
	Title    string `json:"title"`
	URL      string `json:"url"`
}

func (c *OpenLibraryClient) Query(q string, p map[string]string) (books []*book.Book) {
	return
}

func (c *OpenLibraryClient) QueryIsbn(isbn string) (b *book.Book) {
	b = book.New()
	url := fmt.Sprintf("http://openlibrary.org/api/books?bibkeys=ISBN:%s&format=json&jscmd=data", isbn)
	resp, err := c.client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	byteData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var m map[string]openLibraryApiResponse
	err = json.Unmarshal(byteData, &m)
	if err != nil {
		log.Println("[openlibrary] error parsing json", err)
		return
	}

	var key string
	for k := range m {
		key = k
	}
	b = c.buildBook(m[key])
	return
}

func (c OpenLibraryClient) buildBook(data openLibraryApiResponse) *book.Book {
	b := book.New()

	if len(data.Identifiers.Isbn13) != 0 {
		b.Isbn = data.Identifiers.Isbn13[0]
	}

	if len(data.Identifiers.Isbn10) != 0 {
		b.Isbn10 = data.Identifiers.Isbn10[0]
	}

	if data.Title != "" {
		b.Title = data.Title
	}

	if len(data.Authors) != 0 {
		for _, w := range data.Authors {
			if w.Name == "" {
				continue
			}
			a := model.Author{}
			a.GetOrCreate(w.Name)

			// TO-DO: add Author.OpenLibraryID by stripping w.URL

			b.Authors = append(b.Authors, a)
		}
	}

	if data.Cover.Large != "" {
		b.ImgSrc = data.Cover.Large
	} else if data.Cover.Medium != "" {
		b.ImgSrc = data.Cover.Medium
	} else if data.Cover.Small != "" {
		b.ImgSrc = data.Cover.Small
	}

	if data.NumberOfPages != 0 {
		b.Pages = data.NumberOfPages
	}

	if len(data.Identifiers.Goodreads) != 0 {
		b.GoodReadsId = data.Identifiers.Goodreads[0]
	}

	if len(data.Identifiers.Google) != 0 {
		b.GoogleBooksId = data.Identifiers.Google[0]
	}

	if len(data.Identifiers.Openlibrary) != 0 {
		b.OpenLibraryId = data.Identifiers.Openlibrary[0]
	}

	if len(data.Subjects) != 0 {
		for _, sub := range data.Subjects {
			b.Category = append(b.Category, sub.Name)
		}
	}
	return b
}
