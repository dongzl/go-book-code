package main

type Movie struct {
	IMDB        string `json:"imdb"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
	Duration    uint32 `json:"duration"`
	Lang        string `json:"lang"`
}

var (
	mapMovies map[string]*Movie
	slcMovies []*Movie
)
