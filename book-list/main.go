package main

import (
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

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets all books")
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
