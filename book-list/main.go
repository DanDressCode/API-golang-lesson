package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang pointers", Author: "Mr.Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutine", Author: "Mr.Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr.routerrs", Year: "2012"},
		Book{ID: 4, Title: "Golang Currency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr.Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
	//log.Println("Gets all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets one book")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Adds one book")
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Updates one book")
}

func removeBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove one book")
}