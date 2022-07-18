package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectionDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	LogFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	LogFatal(err)

	err = db.Ping()
	LogFatal(err)

	return db
}
