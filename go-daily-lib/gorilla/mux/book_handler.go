package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	mapBooks = make(map[string]*Book)
	slcBooks = make([]*Book, 0, 1)

	data, err := ioutil.ReadFile("/Users/dongzonglei/source_code/Github/go-book-code/go-daily-lib/gorilla/data/books.json")

	if err != nil {
		log.Fatalf("failed to read book.json:%v", err)
	}

	err = json.Unmarshal(data, &slcBooks)

	if err != nil {
		log.Fatalf("failed to unmarshal books:%v", err)
	}

	for _, book := range slcBooks {
		mapBooks[book.ISBN] = book
	}
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(slcBooks)
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	book, ok := mapBooks[mux.Vars(r)["isbn"]]
	if !ok {
		http.NotFound(w, r)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(book)
}
