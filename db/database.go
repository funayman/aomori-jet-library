package db

import (
	"fmt"
	"log"
	"time"

	"github.com/asdine/storm"
	bolt "github.com/coreos/bbolt"
)

var (
	// SQL is a wrapper for database/sql
	SQL *storm.DB

	// DB storage path
	DefaultConnection = "ringo-toshoshutsu.db"
)

func Connect() {
	var err error
	SQL, err = storm.Open(DefaultConnection, storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
	if err != nil {
		log.Fatal("sql open error: ", err)
	}

	fmt.Printf("DB Connected: %+v\n", SQL.Bolt.Info())
}
