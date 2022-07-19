package main

import (
	"database/sql"
	"log"
	"main/controllers"
	"main/driver"
	"main/models"
	"net/http"

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
	controllers := controllers.Controller{}
	/*
		- 'dbname',
		- 'password',
		- 'port',
		- 'user'

	*/

	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controllers.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controllers.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
