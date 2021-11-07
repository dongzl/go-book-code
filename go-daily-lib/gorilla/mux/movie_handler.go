package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func init() {
	mapMovies = make(map[string]*Movie)
	slcMovies = make([]*Movie, 0, 1)

	data, _ := ioutil.ReadFile("../../data/movies.json")
	json.Unmarshal(data, &slcMovies)
	for _, movie := range slcMovies {
		mapMovies[movie.IMDB] = movie
	}
}

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(slcMovies)
}

func MovieHandler(w http.ResponseWriter, r *http.Request) {
	book, ok := mapMovies[mux.Vars(r)["isbn"]]
	if !ok {
		http.NotFound(w, r)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(book)
}

func InitMoviesRouter(r *mux.Router) {
	ms := r.PathPrefix("/movies").Subrouter()
	ms.HandleFunc("/", MoviesHandler)
	ms.HandleFunc("/{imdb}", MovieHandler)
}
