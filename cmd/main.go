package main

import (
	"fmt"

	"github.com/funayman/aomori-library/client"
	"github.com/funayman/aomori-library/controller"
	"github.com/funayman/aomori-library/db"
	"github.com/funayman/aomori-library/router"
	"github.com/funayman/aomori-library/server"
)

var clients []client.Client

func init() {
}

func main() {
	fmt.Println("this is the cmd for the library app")

	db.Connect()
	client.Init()
	controller.Load()
	/*
		for _, isbn := range []string{"9780545010221", "9784840240536", "3319300024", "0316769487"} {
			//for _, isbn := range []string{"3319300024"} {
			book := client.FindBookByIsbn(isbn)
			db.SQL.Save(book)
		}

		var books []model.Book
		fmt.Println("books in db:")
		err := db.SQL.All(&books)
		if err != nil {
			log.Println("error retreiving books: ", err)
		}
		fmt.Printf("%+v\n", books)
	*/
	server.Start(router.Instance(), server.Server{Port: 8081})
}
