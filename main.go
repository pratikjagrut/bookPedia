package main

import (
	"encoding/json"
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

//Book ...
type Book struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

var books []Book

var db = pg.Connect(&pg.Options{
	User:     "postgres",
	Password: "docker",
	Database: "bookpedia",
	Addr:     "localhost:5432",
})

var f, err = os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

var logger = log.New(f, "prefix", log.LstdFlags)

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	_, err := db.Query(&books, `SELECT * FROM books`)
	logger.Println(err)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, _ := strconv.ParseInt(param["id"], 10, 64)
	book := &Book{
		ID: int(id),
	}
	err := db.Select(book)
	logger.Println(err)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(book)
}

func addNewBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	err := db.Insert(&book)
	logger.Println(err)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, _ := strconv.ParseInt(param["id"], 10, 64)
	book := &Book{
		ID: int(id),
	}
	err := db.Delete(book)
	logger.Println(err)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book/add", addNewBook).Methods("POST")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
