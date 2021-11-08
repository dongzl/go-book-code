package main

import (
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age int
}

var (
	hashKey = securecookie.GenerateRandomKey(16)
	blockKey = securecookie.GenerateRandomKey(16)
	s = securecookie.New(hashKey, blockKey)
)

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/set_cookie", SetCookieHandler)
	r.HandleFunc("/read_cookie", ReadCookieHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}