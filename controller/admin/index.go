package admin

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := mT["index.html"].Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
