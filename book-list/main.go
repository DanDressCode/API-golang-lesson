package main

import (
	"encoding/json"
	"log"
	"main/driver"
	"main/models"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"

	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectionDB()

	router := mux.NewRouter()

	/*
		- 'dbname',
		- 'password',
		- 'port',
		- 'user'

	*/

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	books = []models.Book{}

	rows, err := db.Query("select * from books")
	driver.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		driver.LogFatal(err)

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", params["id"])
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	driver.LogFatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)

	driver.LogFatal(err)

	json.NewEncoder(w).Encode(bookID)

}

func updateBooks(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	result, err := db.Exec("update books set title=$1, author=$2,year=$3 where id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)
	driver.LogFatal(err)

	rowsUpdated, err := result.RowsAffected()
	driver.LogFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id = $1", params["id"])
	driver.LogFatal(err)

	rowsDeleted, err := result.RowsAffected()
	driver.LogFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}
