package main

import "github.com/gorilla/mux"

type Book struct {
	ISBN        string `json:"isbn"`
	Name        string `json:"name"`
	Authors     string `json:"authors"`
	Press       string `json:"press"`
	PublishedAt string `json:"published_at"`
}

var (
	mapBooks map[string]*Book
	slcBooks []*Book
)

func InitBooksRouter(r *mux.Router) {
	bs := r.PathPrefix("/books").Subrouter()
	bs.HandleFunc("/", BooksHandler)
	bs.HandleFunc("/{isbn}", BookHandler).Name("book")
	bs.HandleFunc("/books/{isbn:\\d{3}-\\d-\\d{3}-\\d{5}-\\d}", BookHandler)
}